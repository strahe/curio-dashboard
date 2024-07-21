package db

import (
	"github.com/strahe/curio-dashboard/model"
	"github.com/xo/dburl"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewAppDb(dbURL string) (*gorm.DB, error) {
	u, err := dburl.Parse(dbURL)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(postgres.Open(u.DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(
		&model.TaskHistoryAggregates{},
		&model.MinerInfo{},
		&model.StorageUsage{},
	)

	return db, err
}
