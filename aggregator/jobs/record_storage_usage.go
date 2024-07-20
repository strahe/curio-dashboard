package jobs

import (
	"context"
	"time"

	"github.com/strahe/curio-dashboard/graph/loaders"
	"github.com/strahe/curio-dashboard/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RecordStorageUsage struct {
	loader *loaders.Loader
	appDB  *gorm.DB
}

func (j RecordStorageUsage) Run() {
	start := time.Now()
	defer func() {
		log.Infof("%s finished in %s", j.Name(), time.Since(start))
	}()

	paths, err := j.loader.StoragePaths(context.TODO())
	if err != nil {
		log.Errorf("failed to load storage paths: %s", err)
		return
	}

	now := time.Now()

	var records []*model.StorageUsage
	for _, p := range paths {
		records = append(records, &model.StorageUsage{
			Time:        now,
			StorageID:   p.ID,
			Available:   p.Available,
			FsAvailable: p.FsAvailable,
			Reserved:    p.Reserved,
			Used:        p.Used,
		})
	}
	if len(records) > 0 {
		j.appDB.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Save(records)
	}
}

func (j RecordStorageUsage) Spec() string {
	return "@hourly"
}

func (j RecordStorageUsage) Name() string {
	return "RecordStorageUsage"
}

func NewRecordStorageUsage(loader *loaders.Loader, appDB *gorm.DB) *RecordStorageUsage {
	return &RecordStorageUsage{
		loader: loader,
		appDB:  appDB,
	}
}

var _ Job = (*RecordStorageUsage)(nil)
