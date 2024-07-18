package loaders

import (
	"context"
	"github.com/strahe/curio-dashboard/graph/model"
	"time"
)

func (l *Loader) TaskAggregatesByDay(_ context.Context, start, end time.Time) ([]*model.TaskAggregate, error) {
	var out []*model.TaskAggregate

	err := l.appDB.Raw(`SELECT
    DATE(time) AS time,
    SUM(total) AS total,
    SUM(failure) AS failure,
    SUM(success) AS success
FROM
    task_history_aggregates
WHERE
    time BETWEEN ? AND ?
GROUP BY
    time
ORDER BY
    time DESC;`, start, end).Scan(&out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (l *Loader) TaskAggregatesByHour(_ context.Context, start, end time.Time) ([]*model.TaskAggregate, error) {
	var out []*model.TaskAggregate

	err := l.appDB.Raw(`SELECT
    time,
    SUM(total) AS total,
    SUM(failure) AS failure,
    SUM(success) AS success
FROM
    task_history_aggregates
WHERE
    time BETWEEN ? AND ?
GROUP BY
    time
ORDER BY
    time;`, start, end).Scan(&out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (l *Loader) TaskAggregatesByTask(_ context.Context, start, end time.Time) ([]*model.TaskNameAggregate, error) {
	var out []*model.TaskNameAggregate

	err := l.appDB.Raw(`SELECT
    task,
    SUM(total) AS total,
    SUM(failure) AS failure,
    SUM(success) AS success
FROM
    task_history_aggregates
WHERE
    time BETWEEN ? AND ?
GROUP BY
    task`, start, end).Scan(&out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}
