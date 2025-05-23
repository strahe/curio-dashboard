package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.68

import (
	"context"
	"time"

	"github.com/web3tea/curio-dashboard/graph/cachecontrol"
	"github.com/web3tea/curio-dashboard/graph/model"
)

// NodesInfo is the resolver for the nodesInfo field.
func (r *queryResolver) NodesInfo(ctx context.Context) ([]*model.NodeInfo, error) {
	infos, err := r.curioAPI.SyncerState(ctx)
	if err != nil {
		return nil, err
	}

	var out []*model.NodeInfo
	for _, info := range infos {
		out = append(out, &model.NodeInfo{
			ID:        info.Address,
			Address:   info.Address,
			Layers:    info.CLayers,
			Reachable: info.Reachable,
			SyncState: info.SyncState,
			Version:   info.Version,
		})
	}

	cachecontrol.SetHint(ctx, cachecontrol.ScopePrivate, time.Second*30)
	return out, nil
}

// NodeHealthSummary is the resolver for the nodeHealthSummary field.
func (r *queryResolver) NodeHealthSummary(ctx context.Context) (*model.NodeHealthSummary, error) {
	cachecontrol.SetHint(ctx, cachecontrol.ScopePrivate, time.Minute)
	return r.loader.NodeHealthSummary(ctx)
}
