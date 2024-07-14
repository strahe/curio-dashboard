package resolvers

import (
	"github.com/filecoin-project/lotus/api/v1api"
	"github.com/strahe/curio-dashboard/db"
	"github.com/strahe/curio-dashboard/graph/loaders"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	db       *db.HarmonyDB
	fullNode v1api.FullNode
	loader   *loaders.Loader
}

func NewResolver(db *db.HarmonyDB, fullNode v1api.FullNode) *Resolver {
	return &Resolver{
		db:       db,
		fullNode: fullNode,
		loader:   loaders.NewLoader(db, 1000),
	}
}
