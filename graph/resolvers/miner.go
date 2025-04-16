package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.68

import (
	"context"
	"fmt"
	"time"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/api"
	types2 "github.com/filecoin-project/lotus/chain/types"
	"github.com/multiformats/go-multiaddr"
	"github.com/samber/lo"
	"github.com/web3tea/curio-dashboard/graph"
	"github.com/web3tea/curio-dashboard/graph/cachecontrol"
	"github.com/web3tea/curio-dashboard/graph/model"
	"github.com/web3tea/curio-dashboard/types"
)

// Info is the resolver for the info field.
func (r *minerResolver) Info(ctx context.Context, obj *model.Miner) (*model.MinerInfo, error) {
	mi, err := r.fullNode.StateMinerInfo(ctx, obj.ID.Address, types2.EmptyTSK)
	if err != nil {
		return nil, err
	}
	res := &model.MinerInfo{
		Owner:     types.Address{Address: mi.Owner},
		Worker:    types.Address{Address: mi.Worker},
		NewWorker: types.Address{Address: mi.NewWorker},
		ControlAddresses: lo.Map(mi.ControlAddresses, func(addr address.Address, _ int) *types.Address {
			return &types.Address{Address: addr}
		}),
		WorkerChangeEpoch: int(mi.WorkerChangeEpoch),
		PeerID: lo.IfF(mi.PeerId != nil, func() string {
			return mi.PeerId.String()
		}).Else(""),
		MultiAddrs: lo.Map(mi.Multiaddrs, func(addr abi.Multiaddrs, idx int) string {
			ma, err := multiaddr.NewMultiaddrBytes(addr)
			if err != nil {
				log.Errorf("failed to parse multiaddr: %s", err)
				return ""
			}
			return ma.String()
		}),
		WindowPoStProofType:        int(mi.WindowPoStProofType),
		SectorSize:                 int(mi.SectorSize),
		WindowPoStPartitionSectors: int(mi.WindowPoStPartitionSectors),
		ConsensusFaultElapsed:      int(mi.ConsensusFaultElapsed),
		Beneficiary:                types.Address{Address: mi.Beneficiary},
	}

	if mi.PendingOwnerAddress != nil {
		res.PendingOwnerAddress = types.Address{Address: *mi.PendingOwnerAddress}
	}

	if mi.PendingBeneficiaryTerm != nil {
		res.PendingBeneficiaryChange = &model.MinerPendingBeneficiaryChange{
			NewBeneficiary:        types.Address{Address: mi.PendingBeneficiaryTerm.NewBeneficiary},
			NewQuota:              types.BigInt{Int: mi.PendingBeneficiaryTerm.NewQuota.Int},
			NewExpiration:         int(mi.PendingBeneficiaryTerm.NewExpiration),
			ApprovedByBeneficiary: mi.PendingBeneficiaryTerm.ApprovedByBeneficiary,
			ApprovedByNominee:     mi.PendingBeneficiaryTerm.ApprovedByNominee,
		}
	}

	if mi.BeneficiaryTerm != nil {
		res.BeneficiaryTerm = &model.MinerBeneficiaryTerm{
			Quota:      types.BigInt{Int: mi.BeneficiaryTerm.Quota.Int},
			UsedQuota:  types.BigInt{Int: mi.BeneficiaryTerm.UsedQuota.Int},
			Expiration: int(mi.BeneficiaryTerm.Expiration),
		}
	}

	cachecontrol.SetHint(ctx, cachecontrol.ScopePrivate, time.Second*30)
	return res, nil
}

// Power is the resolver for the power field.
func (r *minerResolver) Power(ctx context.Context, obj *model.Miner) (*model.MinerPower, error) {
	mp, err := r.fullNode.StateMinerPower(ctx, obj.ID.Address, types2.EmptyTSK)
	if err != nil {
		return nil, err
	}
	res := &model.MinerPower{
		MinerPower: &model.PowerClaim{
			RawBytePower:    &types.BigInt{Int: mp.MinerPower.RawBytePower.Int},
			QualityAdjPower: &types.BigInt{Int: mp.MinerPower.QualityAdjPower.Int},
		},
		TotalPower: &model.PowerClaim{
			RawBytePower:    &types.BigInt{Int: mp.TotalPower.RawBytePower.Int},
			QualityAdjPower: &types.BigInt{Int: mp.TotalPower.QualityAdjPower.Int},
		},
		ID:          obj.ID.String(),
		HasMinPower: mp.HasMinPower,
	}
	cachecontrol.SetHint(ctx, cachecontrol.ScopePrivate, time.Minute*5)
	return res, nil
}

// Balance is the resolver for the balance field.
func (r *minerResolver) Balance(ctx context.Context, obj *model.Miner) (*model.MinerBalance, error) {
	return &model.MinerBalance{
		ID:    obj.ID,
		Actor: obj.Actor,
	}, nil
}

// Balance is the resolver for the balance field.
func (r *minerBalanceResolver) Balance(ctx context.Context, obj *model.MinerBalance) (*types.BigInt, error) {
	return &types.BigInt{
		Int: obj.Actor.Balance.Int,
	}, nil
}

// Available is the resolver for the available field.
func (r *minerBalanceResolver) Available(ctx context.Context, obj *model.MinerBalance) (*types.BigInt, error) {
	mas, err := obj.State(ctx, r.fullNode)
	if err != nil {
		return nil, err
	}
	available, err := mas.AvailableBalance(obj.Actor.Balance)
	if err != nil {
		return nil, err
	}

	return &types.BigInt{
		Int: available.Int,
	}, nil
}

// InitialPledge is the resolver for the initialPledge field.
func (r *minerBalanceResolver) InitialPledge(ctx context.Context, obj *model.MinerBalance) (*types.BigInt, error) {
	mas, err := obj.State(ctx, r.fullNode)
	if err != nil {
		return nil, err
	}
	lockedFunds, err := mas.LockedFunds()
	if err != nil {
		return nil, err
	}
	return &types.BigInt{
		Int: lockedFunds.InitialPledgeRequirement.Int,
	}, nil
}

// Vesting is the resolver for the vesting field.
func (r *minerBalanceResolver) Vesting(ctx context.Context, obj *model.MinerBalance) (*types.BigInt, error) {
	mas, err := obj.State(ctx, r.fullNode)
	if err != nil {
		return nil, err
	}
	lockedFunds, err := mas.LockedFunds()
	if err != nil {
		return nil, err
	}
	return &types.BigInt{
		Int: lockedFunds.VestingFunds.Int,
	}, nil
}

// PreCommitDeposits is the resolver for the preCommitDeposits field.
func (r *minerBalanceResolver) PreCommitDeposits(ctx context.Context, obj *model.MinerBalance) (*types.BigInt, error) {
	mas, err := obj.State(ctx, r.fullNode)
	if err != nil {
		return nil, err
	}
	lockedFunds, err := mas.LockedFunds()
	if err != nil {
		return nil, err
	}
	return &types.BigInt{
		Int: lockedFunds.PreCommitDeposits.Int,
	}, nil
}

// Miner is the resolver for the miner field.
func (r *queryResolver) Miner(ctx context.Context, address types.Address) (*model.Miner, error) {
	act, err := r.fullNode.StateGetActor(ctx, address.Address, types2.EmptyTSK)
	if err != nil {
		return nil, err
	}
	return &model.Miner{
		ID:    address,
		Actor: *act,
	}, nil
}

// MinerPower is the resolver for the minerPower field.
func (r *queryResolver) MinerPower(ctx context.Context, address *types.Address) (*model.MinerPower, error) {
	var (
		rawPower    = types2.NewInt(0)
		qaPower     = types2.NewInt(0)
		mp          *api.MinerPower // the last miner power
		err         error
		handlePower = func(ctx context.Context, addr types.Address) error {
			mp, err = r.fullNode.StateMinerPower(ctx, addr.Address, types2.EmptyTSK)
			if err != nil {
				return fmt.Errorf("failed to get miner power: %w", err)
			}
			rawPower.Add(rawPower.Int, mp.MinerPower.RawBytePower.Int)
			qaPower.Add(qaPower.Int, mp.MinerPower.QualityAdjPower.Int)
			return nil
		}
	)

	if address != nil {
		if err := handlePower(ctx, *address); err != nil {
			return nil, err
		}
	} else {
		actors, err := r.loader.Actors(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get actors: %w", err)
		}
		for _, act := range actors {
			if err := handlePower(ctx, act.Address); err != nil {
				return nil, err
			}
		}
	}
	out := &model.MinerPower{
		MinerPower: &model.PowerClaim{
			RawBytePower:    &types.BigInt{Int: rawPower.Int},
			QualityAdjPower: &types.BigInt{Int: qaPower.Int},
		},
	}
	if mp != nil {
		out.TotalPower = &model.PowerClaim{
			RawBytePower:    &types.BigInt{Int: mp.TotalPower.RawBytePower.Int},
			QualityAdjPower: &types.BigInt{Int: mp.TotalPower.QualityAdjPower.Int},
		}
	}
	if address != nil {
		out.ID = address.String()
	}
	cachecontrol.SetHint(ctx, cachecontrol.ScopePrivate, time.Hour)
	return out, nil
}

// Miner returns graph.MinerResolver implementation.
func (r *Resolver) Miner() graph.MinerResolver { return &minerResolver{r} }

// MinerBalance returns graph.MinerBalanceResolver implementation.
func (r *Resolver) MinerBalance() graph.MinerBalanceResolver { return &minerBalanceResolver{r} }

type minerResolver struct{ *Resolver }
type minerBalanceResolver struct{ *Resolver }
