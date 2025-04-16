package loaders

import (
	"context"
	"strings"

	"github.com/web3tea/curio-dashboard/graph/model"
)

type ConfigLoader interface {
	Config(ctx context.Context, layer string) (*model.Config, error)
	Configs(ctx context.Context) ([]*model.Config, error)
	ConfigUsed(ctx context.Context, layer string) ([]*model.MachineDetail, error)
}

type ConfigLoaderImpl struct {
	loader *Loader
}

func NewConfigLoader(loader *Loader) ConfigLoader {
	return &ConfigLoaderImpl{loader}
}

func (l *ConfigLoaderImpl) Config(ctx context.Context, layer string) (*model.Config, error) {
	var m model.Config
	err := l.loader.db.QueryRow(ctx, "SELECT id,title,config FROM harmony_config WHERE title = $1", layer).
		Scan(&m.ID, &m.Title, &m.Config)
	return &m, err
}

// Configs is the resolver for the configs field.
func (l *ConfigLoaderImpl) Configs(ctx context.Context) ([]*model.Config, error) {
	var m []*model.Config
	if err := l.loader.db.Select(ctx, &m, `SELECT
    id,
    title,
    config
FROM
    harmony_config`); err != nil {
		return nil, err
	}
	return m, nil
}

func (l *ConfigLoaderImpl) configUsedMap(ctx context.Context) (map[string][]*model.MachineDetail, error) {
	cacheKey := "configUsedMap"
	m, ok := l.loader.cache.Get(cacheKey)
	if ok {
		return m.(map[string][]*model.MachineDetail), nil
	}

	mds, err := l.loader.MachineDetails(ctx)
	if err != nil {
		return nil, err
	}
	confMap := make(map[string][]*model.MachineDetail)
	for _, md := range mds {
		layers := strings.Split(md.Layers, ",")
		for _, layer := range layers {
			confMap[layer] = append(confMap[layer], md)
		}
	}

	l.loader.cache.Add(cacheKey, confMap)

	return confMap, nil
}

func (l *ConfigLoaderImpl) ConfigUsed(ctx context.Context, layer string) ([]*model.MachineDetail, error) {
	confMap, err := l.configUsedMap(ctx)
	if err != nil {
		return nil, err
	}

	mds, ok := confMap[layer]
	if !ok {
		return nil, nil
	}

	return mds, nil
}
