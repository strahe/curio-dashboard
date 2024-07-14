package loaders

import (
	"time"

	"github.com/hashicorp/golang-lru/v2/expirable"
	"github.com/strahe/curio-dashboard/db"
)

type Loader struct {
	db    *db.HarmonyDB
	cache *expirable.LRU[string, any] // low level cache
}

func NewLoader(db *db.HarmonyDB, cacheSize int) *Loader {
	return &Loader{
		db:    db,
		cache: expirable.NewLRU[string, any](cacheSize, nil, time.Minute),
	}
}
