package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"
	"time"

	"github.com/strahe/curio-dashboard/graph/cachecontrol"
	"github.com/strahe/curio-dashboard/graph/model"
)

// MarketMk12Deals is the resolver for the marketMk12Deals field.
func (r *queryResolver) MarketMk12Deals(ctx context.Context, filter model.MarketMk12DealFilterInput, limit int, offset int) ([]*model.MarketMk12Deal, error) {
	cachecontrol.SetHint(ctx, cachecontrol.ScopePrivate, time.Minute)
	return r.loader.MarketMk12Deals(ctx, filter, limit, offset)
}

// MarketMk12DealsCount is the resolver for the MarketMk12DealsCount field.
func (r *queryResolver) MarketMk12DealsCount(ctx context.Context, filter model.MarketMk12DealFilterInput) (int, error) {
	cachecontrol.SetHint(ctx, cachecontrol.ScopePrivate, time.Minute)
	return r.loader.MarketMk12DealsCount(ctx, filter)
}
