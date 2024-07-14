package loaders

import (
	"context"

	"github.com/strahe/curio-dashboard/graph/model"
)

// TaskHistories is the resolver for the taskHistories field.
func (l *Loader) TaskHistories(ctx context.Context, offset int, limit int) ([]*model.TaskHistory, error) {
	var out []*model.TaskHistory
	if err := l.db.Select(ctx, &out, "SELECT * FROM harmony_task_history LIMIT $1 OFFSET $2", limit, offset); err != nil {
		return nil, err
	}
	return out, nil
}
