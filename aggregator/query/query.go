package query

import (
	"context"
	"time"

	"github.com/strahe/curio-dashboard/db"
)

type Query struct {
	db *db.HarmonyDB
}

func NewQuery(db *db.HarmonyDB) *Query {
	return &Query{db: db}
}

type TaskHistory struct {
	Hour    time.Time
	Name    string
	Total   int
	Success int
	Failure int
}

func (q *Query) AggregateTaskHistory(ctx context.Context, startTime, endTime time.Time) ([]*TaskHistory, error) {
	var agg []*TaskHistory
	err := q.db.Select(ctx, &agg, `SELECT
    name,
    DATE_TRUNC('hour', work_end) AS hour,
    COUNT(*) AS total,
    COUNT(CASE WHEN result = 'true' THEN 1 END) AS success,
    COUNT(CASE WHEN result = 'false' THEN 1 END) AS failure
FROM
    harmony_task_history
WHERE
    work_end BETWEEN $1 AND $2
GROUP BY
    hour, name;`, startTime, endTime)

	if err != nil {
		return nil, err
	}
	return agg, nil
}
