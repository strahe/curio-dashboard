package jobs

import (
	"context"
	"github.com/filecoin-project/lotus/api/v1api"
	"github.com/strahe/curio-dashboard/graph/loaders"
	graphModel "github.com/strahe/curio-dashboard/graph/model"
	"github.com/strahe/curio-dashboard/model"
	"github.com/strahe/curio-dashboard/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"sync"
	"time"
)

type RecordMinerInfo struct {
	fullNode v1api.FullNode
	loader   *loaders.Loader
	appDB    *gorm.DB
}

func NewRecordMinerInfo(fullNode v1api.FullNode, loader *loaders.Loader, appDB *gorm.DB) *RecordMinerInfo {
	return &RecordMinerInfo{
		fullNode: fullNode,
		loader:   loader,
		appDB:    appDB,
	}
}

func (j *RecordMinerInfo) Run() {
	start := time.Now()
	defer func() {
		log.Infof("%s finished in %s", j.Name(), time.Since(start))
	}()
	actors, err := j.loader.Actors(context.TODO())
	if err != nil {
		log.Errorf("failed to load actors: %s", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	currentHour := time.Now().Truncate(time.Hour)
	var records []*model.MinerInfo
	wg := sync.WaitGroup{}
	wg.Add(len(actors))
	for _, act := range actors {
		go func(act *graphModel.Actor) {
			defer wg.Done()
			m := &model.MinerInfo{
				Time:    currentHour,
				Address: act.Address,
			}
			p, err := act.Power(ctx, j.fullNode)
			if err != nil {
				log.Errorf("failed to get power for %s: %s", act.Address, err)
				return
			}
			if p.HasMinPower {
				m.RawBytePower = types.BigInt{Int: p.MinerPower.RawBytePower.Int}
				m.QAPower = types.BigInt{Int: p.MinerPower.QualityAdjPower.Int}
			}
			actor, err := act.ChainActor(ctx, j.fullNode)
			if err != nil {
				log.Errorf("failed to get actor for %s: %s", act.Address, err)
				return
			}
			m.Balance = types.BigInt{Int: actor.Balance.Int}
			actorState, err := act.MinerState(ctx, j.fullNode)
			if err != nil {
				log.Errorf("failed to get miner state for %s: %s", act.Address, err)
				return
			}

			avail, err := actorState.AvailableBalance(actor.Balance)
			if err != nil {
				log.Errorf("failed to get available balance for %s: %s", act.Address, err)
				return
			}
			m.AvailableBalance = types.BigInt{Int: avail.Int}
			mi, err := actorState.Info()
			if err != nil {
				log.Errorf("failed to get miner info for %s: %s", act.Address, err)
				return
			}
			wb, err := j.fullNode.WalletBalance(ctx, mi.Worker)
			if err != nil {
				log.Errorf("failed to get worker balance for %s: %s", act.Address, err)
				return
			}
			m.WorkerAddress = types.Address{Address: mi.Worker}
			m.WorkerBalance = types.BigInt{Int: wb.Int}
			records = append(records, m)
		}(act)
	}
	wg.Wait()

	if len(records) > 0 {
		j.appDB.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Save(records)
	}
}

func (j *RecordMinerInfo) RunWith(time time.Time) {
	// not support passing time
	panic("implement me")
}

func (j *RecordMinerInfo) Name() string {
	return "RecordMinerInfo"
}

func (j *RecordMinerInfo) Spec() string {
	return "@hourly"
}

var _ Job = &RecordMinerInfo{}
