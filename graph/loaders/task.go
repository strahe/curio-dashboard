package loaders

import (
	"context"
	"time"

	"github.com/samber/lo"
	"github.com/strahe/curio-dashboard/graph/model"
)

func (l *Loader) Task(ctx context.Context, id int) (*model.Task, error) {
	var out model.Task
	err := l.db.QueryRow(ctx, `SELECT id,initiated_by,update_time,posted_time,owner_id,added_by,previous_task,name FROM harmony_task WHERE id = $1`, id).
		Scan(&out.ID, &out.InitiatedByID, &out.UpdateTime, &out.PostedTime, &out.OwnerID, &out.AddedByID, &out.PreviousTaskID, &out.Name)
	return &out, err
}

func (l *Loader) Tasks(ctx context.Context) ([]*model.Task, error) {
	var out []*model.Task
	if err := l.db.Select(ctx, &out, "SELECT * FROM harmony_task"); err != nil {
		return nil, err
	}
	return out, nil
}

func (l *Loader) TaskSummary(ctx context.Context, lastDays int) ([]*model.TaskSummary, error) {
	if lastDays <= 0 {
		lastDays = 1
	} // todo: check max days
	var m []*model.TaskSummary
	err := l.db.Select(ctx, &m, `SELECT name, count(case when result = 'true' then 1 end) as true_count,
		count(case when result = 'false' then 1 end) as false_count, count(*) as total_count
		from harmony_task_history where work_end > current_timestamp - make_interval(days => $1)
		group by name order by total_count desc`, lastDays)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (l *Loader) TaskSummaryByDay(ctx context.Context, lastDays int) ([]*model.TaskSummaryDay, error) {
	var tasks []struct {
		Day    time.Time `db:"day"`
		Result bool      `db:"result"`
		Count  int       `db:"task_count"`
	}

	// todo: cache
	err := l.db.Select(ctx, &tasks,
		`SELECT DATE(work_end) AS day, result, COUNT(*) AS task_count 
FROM harmony_task_history 
WHERE work_end > current_timestamp - make_interval(days => $1) 
GROUP BY day, result 
ORDER BY day DESC, result;`, lastDays)
	if err != nil {
		return nil, err
	}

	var taskMap = make(map[time.Time]*model.TaskSummaryDay)
	for _, task := range tasks {
		day := task.Day
		if _, ok := taskMap[day]; !ok {
			taskMap[day] = &model.TaskSummaryDay{Day: day}
		}
		if task.Result {
			taskMap[day].TrueCount = task.Count
		} else {
			taskMap[day].FalseCount = task.Count
		}
	}
	return lo.MapToSlice(taskMap, func(key time.Time, value *model.TaskSummaryDay) *model.TaskSummaryDay {
		value.TotalCount = value.TrueCount + value.FalseCount
		return value
	}), nil
}
