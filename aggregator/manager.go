package aggregator

import (
	"fmt"

	"github.com/filecoin-project/lotus/api/v1api"
	logging "github.com/ipfs/go-log/v2"
	"github.com/robfig/cron/v3"
	"github.com/samber/lo"
	"github.com/strahe/curio-dashboard/aggregator/jobs"
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
	allJobs := AllJobs(m.db, m.appDB,
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

func AllJobs(db *db.HarmonyDB, appDB *gorm.DB, fullNode v1api.FullNode, loader *loaders.Loader) []jobs.Job {
	return []jobs.Job{
		jobs.NewAggTaskHistory(loader, db, appDB),
		jobs.NewRecordMinerInfo(fullNode, loader, appDB),
		jobs.NewRecordStorageUsage(loader, appDB),
	}
}

func AllBackFillingJobs(js []jobs.Job) []jobs.BackFillingJob {
	var backFillingJobs []jobs.BackFillingJob
	for _, job := range js {
		if bfJob, ok := job.(jobs.BackFillingJob); ok {
			backFillingJobs = append(backFillingJobs, bfJob)
		}
	}
	return backFillingJobs
}

func NoBackFillingJobs(js []jobs.Job) []jobs.Job {
	return lo.Filter(js, func(j jobs.Job, index int) bool {
		_, ok := j.(jobs.BackFillingJob)
		return !ok
	})
}
