package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"

	"github.com/strahe/curio-dashboard/graph"
	"github.com/strahe/curio-dashboard/graph/model"
)

// Detail is the resolver for the detail field.
func (r *machineResolver) Detail(ctx context.Context, obj *model.Machine) (*model.MachineDetail, error) {
	var out model.MachineDetail
	if err := r.db.QueryRow(ctx, "SELECT id,machine_name,tasks,layers,startup_time,miners,machine_id FROM harmony_machine_details WHERE machine_id = $1", obj.ID).
		Scan(&out.ID, &out.MachineName, &out.Tasks, &out.Layers, &out.StartupTime, &out.Miners, &out.MachineID); err != nil {
		return nil, err
	}
	return &out, nil
}

// Tasks is the resolver for the tasks field.
func (r *machineResolver) Tasks(ctx context.Context, obj *model.Machine) ([]*model.Task, error) {
	var out []*model.Task
	if err := r.db.Select(ctx, &out, "SELECT * FROM harmony_task WHERE owner_id = $1", obj.ID); err != nil {
		return nil, err
	}
	return out, nil
}

// TaskHistories is the resolver for the taskHistories field.
func (r *machineResolver) TaskHistories(ctx context.Context, obj *model.Machine, last int) ([]*model.TaskHistory, error) {
	var out []*model.TaskHistory
	if err := r.db.Select(ctx, &out, "SELECT * FROM harmony_task_history WHERE work_end > CURRENT_TIMESTAMP - INTERVAL '24 hours' AND harmony_task_history.completed_by_host_and_port = $1 ORDER BY work_end DESC LIMIT $2", obj.HostAndPort, last); err != nil {
		return nil, err
	}
	return out, nil
}

// Machine returns graph.MachineResolver implementation.
func (r *Resolver) Machine() graph.MachineResolver { return &machineResolver{r} }

type machineResolver struct{ *Resolver }
