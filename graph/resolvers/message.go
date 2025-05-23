package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.68

import (
	"context"
	"time"

	"github.com/web3tea/curio-dashboard/graph/cachecontrol"
	"github.com/web3tea/curio-dashboard/graph/model"
	"github.com/web3tea/curio-dashboard/types"
)

// MessageSends is the resolver for the messageSends field.
func (r *queryResolver) MessageSends(ctx context.Context, account *types.Address, offset int, limit int) ([]*model.MessageSend, error) {
	cachecontrol.SetHint(ctx, cachecontrol.ScopePrivate, time.Minute)
	return r.loader.MessageSends(ctx, account, offset, limit)
}

// MessageSendsCount is the resolver for the messageSendsCount field.
func (r *queryResolver) MessageSendsCount(ctx context.Context, account *types.Address) (int, error) {
	cachecontrol.SetHint(ctx, cachecontrol.ScopePrivate, time.Minute)
	return r.loader.MessageSendCount(ctx, account)
}

// MessageSend is the resolver for the messageSend field.
func (r *queryResolver) MessageSend(ctx context.Context, sendTaskID *int, fromKey *string, nonce *int, signedCid *string) (*model.MessageSend, error) {
	cachecontrol.SetHint(ctx, cachecontrol.ScopePrivate, time.Minute)
	return r.loader.MessageSend(ctx, sendTaskID, fromKey, nonce, signedCid)
}
