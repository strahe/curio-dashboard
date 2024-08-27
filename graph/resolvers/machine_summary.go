package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"

	"github.com/strahe/curio-dashboard/graph"
	"github.com/strahe/curio-dashboard/graph/model"
)

// Total is the resolver for the total field.
func (r *machineSummaryResolver) Total(ctx context.Context, obj *model.MachineSummary) (int, error) {
	var total int
	err := r.db.QueryRow(ctx, "SELECT COUNT(*) FROM harmony_machines").Scan(&total)
	return total, err
}

// TotalUp is the resolver for the totalUp field.
func (r *machineSummaryResolver) TotalUp(ctx context.Context, obj *model.MachineSummary) (int, error) {
	var totalUp int
	err := r.db.QueryRow(ctx, "SELECT COUNT(*) FROM harmony_machines WHERE last_contact > NOW() - INTERVAL '5 MINUTES' ").Scan(&totalUp)
	return totalUp, err
}

// TotalDown is the resolver for the totalDown field.
func (r *machineSummaryResolver) TotalDown(ctx context.Context, obj *model.MachineSummary) (int, error) {
	var totalDown int
	err := r.db.QueryRow(ctx, "SELECT COUNT(*) FROM harmony_machines WHERE last_contact < NOW() - INTERVAL '5 MINUTES' ").Scan(&totalDown)
	return totalDown, err
}

// UniqueHostsTotal is the resolver for the uniqueHostsTotal field.
func (r *machineSummaryResolver) UniqueHostsTotal(ctx context.Context, obj *model.MachineSummary) (int, error) {
	var total int
	err := r.db.QueryRow(ctx, `SELECT
    COUNT(DISTINCT SPLIT_PART(host_and_port, ':', 1)) AS unique_host_count
FROM
    harmony_machines;`).Scan(&total)
	return total, err
}

// UniqueHostsUp is the resolver for the uniqueHostsUp field.
func (r *machineSummaryResolver) UniqueHostsUp(ctx context.Context, obj *model.MachineSummary) (int, error) {
	var totalUp int
	err := r.db.QueryRow(ctx, `SELECT
    COUNT(DISTINCT SPLIT_PART(host_and_port, ':', 1)) AS unique_host_count
FROM
    harmony_machines
WHERE
    last_contact > NOW() - INTERVAL '5 MINUTES'`).Scan(&totalUp)
	return totalUp, err
}

// UniqueHostsDown is the resolver for the uniqueHostsDown field.
func (r *machineSummaryResolver) UniqueHostsDown(ctx context.Context, obj *model.MachineSummary) (int, error) {
	var totalDown int
	err := r.db.QueryRow(ctx, `SELECT
    COUNT(DISTINCT SPLIT_PART(host_and_port, ':', 1)) AS unique_host_count
FROM
    harmony_machines
WHERE
    last_contact <= NOW() - INTERVAL '5 MINUTES'`).Scan(&totalDown)
	return totalDown, err
}

// TotalRAM is the resolver for the totalRam field.
func (r *machineSummaryResolver) TotalRAM(ctx context.Context, obj *model.MachineSummary) (int, error) {
	var total int
	err := r.db.QueryRow(ctx, "SELECT SUM(ram) FROM harmony_machines").Scan(&total)
	return total, err
}

// TotalCPU is the resolver for the totalCpu field.
func (r *machineSummaryResolver) TotalCPU(ctx context.Context, obj *model.MachineSummary) (int, error) {
	var total int
	err := r.db.QueryRow(ctx, "SELECT SUM(cpu) FROM harmony_machines").Scan(&total)
	return total, err
}

// TotalGpu is the resolver for the totalGpu field.
func (r *machineSummaryResolver) TotalGpu(ctx context.Context, obj *model.MachineSummary) (float64, error) {
	var total float64
	err := r.db.QueryRow(ctx, "SELECT SUM(gpu) FROM harmony_machines").Scan(&total)
	return total, err
}

// MachineSummary returns graph.MachineSummaryResolver implementation.
func (r *Resolver) MachineSummary() graph.MachineSummaryResolver { return &machineSummaryResolver{r} }

type machineSummaryResolver struct{ *Resolver }
