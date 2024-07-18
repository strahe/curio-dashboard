package resolvers

import (
	"github.com/filecoin-project/lotus/api/v1api"
	"github.com/strahe/curio-dashboard/db"
	"github.com/strahe/curio-dashboard/graph/loaders"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	db       *db.HarmonyDB
	appDB    *gorm.DB
	fullNode v1api.FullNode
	loader   *loaders.Loader
}

func NewResolver(db *db.HarmonyDB, appDB *gorm.DB, fullNode v1api.FullNode) *Resolver {
	return &Resolver{
		db:       db,
		appDB:    appDB,
		fullNode: fullNode,
		loader:   loaders.NewLoader(db, appDB, 1000),
	}
}
