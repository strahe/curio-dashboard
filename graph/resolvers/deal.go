package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.68

import (
	"context"
	"fmt"
	"time"

	"github.com/web3tea/curio-dashboard/graph/cachecontrol"
	"github.com/web3tea/curio-dashboard/graph/model"
	"github.com/web3tea/curio-dashboard/types"
)

// DealSealNow is the resolver for the dealSealNow field.
func (r *mutationResolver) DealSealNow(ctx context.Context, miner types.Address, sectorNumber uint64) (bool, error) {
	err := r.curioAPI.DealsSealNow(ctx, uint64(miner.ID), sectorNumber)
	if err != nil {
		return false, err
	}
	return true, nil
}

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

// MarketDealInfo is the resolver for the marketDealInfo field.
func (r *queryResolver) MarketDealInfo(ctx context.Context, id string) (*model.DealInfo, error) {
	cachecontrol.SetHint(ctx, cachecontrol.ScopePrivate, time.Minute)
	deal, err := r.curioAPI.StorageDealInfo(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get deal info: %w", err)
	}
	return &model.DealInfo{
		ID:                deal.ID,
		SpID:              types.ActorID{ID: uint64(deal.MinerID)},
		Sector:            types.MustNullInt64(deal.Sector),
		CreatedAt:         deal.CreatedAt,
		SignedProposalCid: deal.SignedProposalCid,
		Offline:           deal.Offline,
		Verified:          deal.Verified,
		StartEpoch:        deal.StartEpoch,
		EndEpoch:          deal.EndEpoch,
		ClientPeerID:      types.MustParsePeerID(deal.ClientPeerId),
		ChainDealID:       types.MustNullInt64(deal.ChainDealId),
		PublishCid:        types.MustNullString(deal.PublishCid),
		PieceCid:          deal.PieceCid,
		PieceSize:         deal.PieceSize,
		FastRetrieval:     deal.FastRetrieval,
		AnnounceToIpni:    deal.AnnounceToIpni,
		Urls:              deal.URLS,
		URLHeaders:        types.MustJSON(deal.UrlHeaders),
		Error:             deal.Error,
		Miner:             deal.Miner,
		IsLegacy:          deal.IsLegacy,
		Indexed:           types.MustNullBool(deal.Indexed),
		IsDdo:             deal.IsDDO,
	}, nil
}

// MarketDealCountSummary is the resolver for the marketDealCountSummary field.
func (r *queryResolver) MarketDealCountSummary(ctx context.Context) (*model.DealCountSummary, error) {
	cachecontrol.SetHint(ctx, cachecontrol.ScopePrivate, time.Minute*10)
	return r.loader.MarketDealCountSummary(ctx)
}

// DealsPending is the resolver for the dealsPending field.
func (r *queryResolver) DealsPending(ctx context.Context) ([]*model.OpenSectorPiece, error) {
	cachecontrol.SetHint(ctx, cachecontrol.ScopePrivate, time.Minute)
	return r.loader.OpenSectorPieces(ctx)
}
