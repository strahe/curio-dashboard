package db

import (
	"context"
	"fmt"

	"github.com/strahe/curio-dashboard/model"
	"github.com/xo/dburl"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewAppDb(ctx context.Context, dbURL string) (*gorm.DB, error) {
	u, err := dburl.Parse(dbURL)
	if err != nil {
		return nil, err
	}

	sqlDB, err := dburl.Open(dbURL)
	if err != nil {
		return nil, err
	}

	if err := sqlDB.PingContext(ctx); err != nil {
		return nil, err
	}

	var dialector gorm.Dialector

	switch u.Scheme {
	case "postgres":
		dialector = postgres.New(postgres.Config{
			Conn: sqlDB,
		})
	case "sqlite", "sqlite3":
		dialector = sqlite.Dialector{
			Conn: sqlDB,
		}
	default:
		return nil, fmt.Errorf("unsupported database type: %s", u.Scheme)
	}

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.TaskHistoryAggregates{}, &model.MinerInfo{})

	return db, err
}
