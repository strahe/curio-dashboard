package aggregator

import (
	"fmt"

	"github.com/filecoin-project/lotus/api/v1api"
	logging "github.com/ipfs/go-log/v2"
	"github.com/robfig/cron/v3"
	"github.com/strahe/curio-dashboard/aggregator/jobs"
	"github.com/strahe/curio-dashboard/aggregator/query"
	"github.com/strahe/curio-dashboard/db"
	"github.com/strahe/curio-dashboard/graph/loaders"
	"gorm.io/gorm"
)

var log = logging.Logger("aggregator")

type Manager struct {
	fullNode v1api.FullNode
	db       *db.HarmonyDB
	appDB    *gorm.DB
	cron     *cron.Cron
}

func NewManager(fullNode v1api.FullNode, db *db.HarmonyDB, appDB *gorm.DB) (*Manager, error) {
	nm := &Manager{
		fullNode: fullNode,
		db:       db,
		appDB:    appDB,
		cron:     cron.New(),
	}
	if err := nm.initAggregators(); err != nil {
		return nil, err
	}
	return nm, nil
}

func (m *Manager) Start() {
	m.cron.Start()
}

func (m *Manager) Stop() {
	m.cron.Stop()
}

func (m *Manager) initAggregators() error {
	allJobs := AllJobs(query.NewQuery(m.db), m.db, m.appDB,
		m.fullNode, loaders.NewLoader(m.db, m.appDB, 1000))
	for _, job := range allJobs {
		log.Infof("Adding job %s with spec: %s", job.Name(), job.Spec())
		_, err := m.cron.AddJob(job.Spec(), job)
		if err != nil {
			return fmt.Errorf("failed to add job %s: %s", job.Name(), err)
		}
	}
	return nil
}

func AllJobs(query *query.Query, db *db.HarmonyDB, appDB *gorm.DB, fullNode v1api.FullNode, loader *loaders.Loader) []jobs.Job {
	res := []jobs.Job{
		jobs.NewAggTaskHistory(query, db, appDB),
	}
	res = append(res, AllOnlyRealtimeJobs(appDB, fullNode, loader)...)
	return res
}

func AllOnlyRealtimeJobs(appDB *gorm.DB, fullNode v1api.FullNode, loader *loaders.Loader) []jobs.Job {
	return []jobs.Job{
		jobs.NewRecordMinerInfo(fullNode, loader, appDB),
	}
}
