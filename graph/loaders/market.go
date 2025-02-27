package loaders

import (
	"context"

	"github.com/filecoin-project/curio/web/api/webrpc"
	"github.com/strahe/curio-dashboard/graph/model"
	"github.com/strahe/curio-dashboard/types"
)

type MarketLoader interface {
	MarketMk12StorageAsks(ctx context.Context) ([]*model.MarketMk12StorageAsk, error)
	MarketMk12StorageAsk(ctx context.Context, spID uint64) (*model.MarketMk12StorageAsk, error)
	MarketMk12StorageAsksCount(ctx context.Context) (int, error)
	MarketMk12PriceFilter(ctx context.Context, name string) (*webrpc.PriceFilter, error)
	MarketMk12ClientFilter(ctx context.Context, name string) (*webrpc.ClientFilter, error)
	MarketAllowFilter(ctx context.Context, wallet types.Address) (*model.MarketAllowFilter, error)
	MarketMk12Deals(ctx context.Context, filter model.MarketMk12DealFilterInput, limit int, offset int) ([]*model.MarketMk12Deal, error)
	MarketMk12DealsCount(ctx context.Context, filter model.MarketMk12DealFilterInput) (int, error)
	MarketPieceMetadata(ctx context.Context, id types.Cid) (*model.MarketPieceMetadata, error)
	MakretPieceDeals(ctx context.Context, id types.Cid) ([]*model.MarketPieceDeal, error)
}

type MarketLoaderImpl struct {
	loader *Loader
}

func NewMarketLoader(loader *Loader) MarketLoader {
	return &MarketLoaderImpl{loader}
}

func (l *MarketLoaderImpl) MarketMk12StorageAsks(ctx context.Context) ([]*model.MarketMk12StorageAsk, error) {
	var result []*model.MarketMk12StorageAsk

	err := l.loader.db.Select(ctx, &result, `SELECT
		sp_id,
		price,
		verified_price,
		min_size,
		max_size,
		created_at,
		expiry,
		sequence
		FROM market_mk12_storage_ask`)
	return result, err
}

func (l *MarketLoaderImpl) MarketMk12StorageAsk(ctx context.Context, spID uint64) (*model.MarketMk12StorageAsk, error) {
	var result model.MarketMk12StorageAsk

	err := l.loader.db.QueryRow(ctx, `SELECT
		sp_id,
		price,
		verified_price,
		min_size,
		max_size,
		created_at,
		expiry,
		sequence
		FROM market_mk12_storage_ask
		WHERE sp_id = $1`, spID).Scan(
		&result.SpID,
		&result.Price,
		&result.VerifiedPrice,
		&result.MinSize,
		&result.MaxSize,
		&result.CreatedAt,
		&result.Expiry,
		&result.Sequence,
	)
	return &result, err
}

func (l *MarketLoaderImpl) MarketMk12StorageAsksCount(ctx context.Context) (int, error) {
	var result int
	err := l.loader.db.QueryRow(ctx, `SELECT COUNT(*) FROM market_mk12_storage_ask`).Scan(&result)
	return result, err
}

func (l *MarketLoaderImpl) MarketMk12PriceFilter(ctx context.Context, name string) (*webrpc.PriceFilter, error) {
	var res []*webrpc.PriceFilter
	err := l.loader.db.Select(ctx, &res, `
	SELECT
		name,
		min_duration_days,
		max_duration_days,
		min_size,
		max_size,
		price,
		verified
	FROM market_mk12_pricing_filters
	WHERE name = $1`, name)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, ErrorNotFound
	}
	return res[0], nil
}

func (l *MarketLoaderImpl) MarketMk12ClientFilter(ctx context.Context, name string) (*webrpc.ClientFilter, error) {
	var res []*webrpc.ClientFilter
	err := l.loader.db.Select(ctx, &res, `
	SELECT
		name,
		active,
		wallets,
		peer_ids,
		pricing_filters,
		max_deals_per_hour,
		max_deal_size_per_hour,
		additional_info
	FROM market_mk12_client_filters
	WHERE name = $1`, name)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, ErrorNotFound
	}
	return res[0], nil
}

func (l *MarketLoaderImpl) MarketAllowFilter(ctx context.Context, wallet types.Address) (*model.MarketAllowFilter, error) {
	var res []*model.MarketAllowFilter
	err := l.loader.db.Select(ctx, &res, `
	SELECT
		wallet,
		status
	FROM market_allow_list
	WHERE wallet = $1`, wallet)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, ErrorNotFound
	}
	return res[0], nil
}

func (l *MarketLoaderImpl) MarketMk12Deals(ctx context.Context, filter model.MarketMk12DealFilterInput, limit int, offset int) ([]*model.MarketMk12Deal, error) {
	var result []*model.MarketMk12Deal
	err := l.loader.db.Select(ctx, &result, `SELECT
            uuid,
            sp_id,
            created_at,
            signed_proposal_cid,
            proposal_signature,
            proposal,
            offline,
            verified,
            start_epoch,
            end_epoch,
            client_peer_id,
            chain_deal_id,
            publish_cid,
            piece_cid,
            piece_size,
            fast_retrieval,
            announce_to_ipni,
            url,
            url_headers,
            error,
            label,
            proposal_cid
        FROM market_mk12_deals
        WHERE 1=1
        AND ($1::int8 IS NULL OR sp_id = $1)
        AND ($2::text IS NULL OR uuid = $2)
        AND ($3::text IS NULL OR signed_proposal_cid = $3)
        AND ($4::text IS NULL OR proposal_cid = $4)
        AND ($5::text IS NULL OR piece_cid = $5)
        LIMIT $6 OFFSET $7
    `,
		filter.SpID,
		filter.UUID,
		filter.SignedProposalCid,
		filter.ProposalCid,
		filter.PieceCid,
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (l *MarketLoaderImpl) MarketMk12DealsCount(ctx context.Context, filter model.MarketMk12DealFilterInput) (int, error) {
	var result int
	err := l.loader.db.QueryRow(ctx, `SELECT COUNT(*) FROM market_mk12_deals
		WHERE 1=1
		AND ($1::int8 IS NULL OR sp_id = $1)
		AND ($2::text IS NULL OR uuid = $2)
		AND ($3::text IS NULL OR signed_proposal_cid = $3)
		AND ($4::text IS NULL OR proposal_cid = $4)
		AND ($5::text IS NULL OR piece_cid = $5)
	`,
		filter.SpID,
		filter.UUID,
		filter.SignedProposalCid,
		filter.ProposalCid,
		filter.PieceCid,
	).Scan(&result)
	return result, err
}

func (l *MarketLoaderImpl) MarketPieceMetadata(ctx context.Context, id types.Cid) (*model.MarketPieceMetadata, error) {
	var result model.MarketPieceMetadata
	err := l.loader.db.QueryRow(ctx, `SELECT
		piece_cid,
		piece_size,
		version,
		created_at,
		indexed,
		indexed_at FROM market_piece_metadata
		WHERE piece_cid = $1`, id.String()).Scan(
		&result.PieceCid,
		&result.PieceSize,
		&result.Version,
		&result.CreatedAt,
		&result.Indexed,
		&result.IndexedAt)
	return &result, err
}

func (l *MarketLoaderImpl) MakretPieceDeals(ctx context.Context, id types.Cid) ([]*model.MarketPieceDeal, error) {
	var result []*model.MarketPieceDeal
	err := l.loader.db.Select(ctx, &result, `SELECT
		id,
		piece_cid,
		boost_deal,
		legacy_deal,
		chain_deal_id,
		sp_id,
		sector_num,
		piece_offset,
		piece_length,
		raw_size
		FROM market_piece_deal
		WHERE piece_cid = $1`, id.String())
	return result, err
}
