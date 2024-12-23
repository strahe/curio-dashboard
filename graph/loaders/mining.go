package loaders

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/strahe/curio-dashboard/types"

	"github.com/strahe/curio-dashboard/graph/model"
)

type MiningLoader interface {
	MiningSummaryByDay(ctx context.Context, start, end time.Time) ([]*model.MiningSummaryDay, error)
	MiningCount(ctx context.Context, start, end time.Time, actor *types.ActorID) (*model.MiningCount, error)
	MiningTasks(ctx context.Context, start time.Time, end time.Time, actor *types.ActorID, won *bool) ([]*model.MiningTask, error)
	MiningTasksCount(ctx context.Context, start, end *time.Time, actor *types.ActorID, won *bool) (int, error)
	MiningCountSummary(ctx context.Context, start, end time.Time, actor *types.ActorID) (*model.MiningCountSummary, error)
	MiningCountAggregate(ctx context.Context, start, end time.Time, actor *types.ActorID, unit model.TaskHistoriesAggregateInterval) ([]*model.MiningCountAggregated, error)
}

func (l *Loader) MiningSummaryByDay(ctx context.Context, start, end time.Time) ([]*model.MiningSummaryDay, error) {
	if end.IsZero() {
		end = time.Now()
	}
	if end.Before(start) {
		return nil, fmt.Errorf("end time is before start time")
	}

	var result []*model.MiningSummaryDay
	err := l.db.Select(ctx, &result,
		`SELECT
    DATE_TRUNC('day', base_compute_time) AS day,
    sp_id as miner,
    COUNT(*) AS won_block
FROM
    mining_tasks
WHERE
    won = true
  AND base_compute_time BETWEEN $1 AND $2
GROUP BY
    day, sp_id
ORDER BY
    day, sp_id;`, start, end)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (l *Loader) MiningCount(ctx context.Context, start, end time.Time, actor *types.ActorID) (*model.MiningCount, error) {
	if end.IsZero() {
		end = time.Now()
	}
	if end.Before(start) {
		return nil, fmt.Errorf("end time is before start time")
	}
	type mm struct {
		Included sql.NullBool `db:"included"`
		Count    int          `db:"count"`
	}
	var res []mm

	err := l.db.Select(ctx, &res, `
SELECT
    included,
    COUNT(*) AS count
FROM
    mining_tasks
WHERE
    won = true AND
    ($1::int IS NULL OR sp_id = $1) AND
    base_compute_time BETWEEN $2 AND $3
GROUP BY
    included;`, actor, start, end)
	if err != nil {
		return nil, err
	}
	result := &model.MiningCount{}

	for _, r := range res {
		if r.Included.Bool {
			result.Include = r.Count
		} else {
			result.Exclude = r.Count
		}
	}
	return result, nil
}

func (l *Loader) MiningTasksCount(ctx context.Context, start, end *time.Time, actor *types.ActorID, won *bool, include *bool) (int, error) {
	var count int
	err := l.db.QueryRow(ctx, `
SELECT COUNT(*)
FROM
    mining_tasks
WHERE
    ($1::bool IS NULL OR won = $1) AND
    ($2::int IS NULL OR sp_id = $2) AND
    ($3::bool IS NULL OR included = $3) AND
    ($4::timestamp IS NULL OR base_compute_time >= $4) AND
    ($5::timestamp IS NULL OR base_compute_time <= $5);`, won, actor, include, start, end).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (l *Loader) MiningTasks(ctx context.Context, start, end *time.Time, actor *types.ActorID, won *bool, include *bool, offset int, limit int) ([]*model.MiningTask, error) {
	var result []*model.MiningTask

	err := l.db.Select(ctx, &result, `
SELECT
    task_id,
    sp_id,
    epoch,
    base_compute_time,
    won,
    mined_cid,
    mined_header,
    mined_at,
    submitted_at,
    included
FROM
    mining_tasks
WHERE
    ($1::bool IS NULL OR won = $1) AND
    ($2::int IS NULL OR sp_id = $2) AND
    ($3::bool IS NULL OR included = $3) AND
    ($4::timestamp IS NULL OR base_compute_time >= $4) AND
    ($5::timestamp IS NULL OR base_compute_time <= $5)
ORDER BY
    base_compute_time DESC 
LIMIT $6 OFFSET $7;`, won, actor, include, start, end, limit, offset)

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (l *Loader) MiningCountSummary(ctx context.Context, start, end time.Time, actor *types.ActorID) (*model.MiningCountSummary, error) {
	if end.IsZero() {
		end = time.Now()
	}
	if end.Before(start) {
		return nil, fmt.Errorf("end time (%v) is before start time (%v)", end, start)
	}
	stats := model.MiningCountSummary{
		Start: start,
		End:   end,
		Actor: actor,
	}

	query := `
        WITH params AS (
            SELECT
                $1::timestamptz as start_time,
                $2::timestamptz as end_time,
                $3::int8 as actor_id
        )
        SELECT 
            COUNT(*) as total,
            COUNT(CASE WHEN won = true THEN 1 END) as won,
            COUNT(CASE WHEN included = true THEN 1 END) as included
        FROM mining_tasks mt
        WHERE base_compute_time BETWEEN (SELECT start_time FROM params) AND (SELECT end_time FROM params)
          AND (
              (SELECT actor_id FROM params) IS NULL 
              OR mt.sp_id = (SELECT actor_id FROM params)
          )
    `
	err := l.db.QueryRow(ctx, query, start, end, actor).Scan(&stats.Total, &stats.Won, &stats.Included)
	if err != nil {
		return nil, err
	}
	return &stats, nil
}

func (l *Loader) MiningCountAggregate(ctx context.Context, start, end time.Time, actor *types.ActorID, interval model.MiningTaskAggregateInterval) ([]*model.MiningCountAggregated, error) {
	if end.IsZero() {
		end = time.Now()
	}
	if end.Before(start) {
		return nil, fmt.Errorf("end time is before start time")
	}
	var result []*model.MiningCountAggregated

	query := `
    WITH params AS (
        SELECT
            $1::timestamptz as start_time,
            $2::timestamptz as end_time,
            $3::int8 as actor_id
    )
    SELECT
        DATE_TRUNC($4, base_compute_time) AS time,
        COUNT(*) AS total,
        COUNT(CASE WHEN won = true THEN 1 END) AS won,
        COUNT(CASE WHEN included = true THEN 1 END) AS included
    FROM
        mining_tasks
    WHERE
        base_compute_time BETWEEN (SELECT start_time FROM params) AND (SELECT end_time FROM params)
        AND (
            (SELECT actor_id FROM params) IS NULL
            OR sp_id = (SELECT actor_id FROM params)
        )
    GROUP BY
        time
    ORDER BY
        time;
`
	err := l.db.Select(ctx, &result, query, start, end, actor, interval)
	if err != nil {
		return nil, err
	}
	return result, nil
}
