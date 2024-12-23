package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"
	"time"

	"github.com/strahe/curio-dashboard/graph"
	"github.com/strahe/curio-dashboard/graph/model"
)

// Tasks is the resolver for the tasks field.
func (r *taskAggregateResolver) Tasks(ctx context.Context, obj *model.TaskAggregate) ([]*model.TaskNameAggregate, error) {
	start := obj.Time
	end := obj.Time.Add(time.Hour * 24)

	return r.loader.TaskAggregatesByTask(ctx, start, end)
}

// TaskAggregate returns graph.TaskAggregateResolver implementation.
func (r *Resolver) TaskAggregate() graph.TaskAggregateResolver { return &taskAggregateResolver{r} }

type taskAggregateResolver struct{ *Resolver }
