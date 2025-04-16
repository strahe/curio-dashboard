package loaders

import (
	"context"

	"github.com/BurntSushi/toml"
	"github.com/filecoin-project/go-address"
	"github.com/web3tea/curio-dashboard/graph/model"
	"github.com/web3tea/curio-dashboard/types"
	"golang.org/x/xerrors"
)

type ActorLoader interface {
	Actors(ctx context.Context) ([]*model.Actor, error)
	Actor(ctx context.Context, address types.Address) (*model.Actor, error)
}

type ActorLoaderImpl struct {
	loader *Loader
}

func NewActorLoader(loader *Loader) ActorLoader {
	return &ActorLoaderImpl{loader}
}

func (l *ActorLoaderImpl) Actors(ctx context.Context) ([]*model.Actor, error) {
	cfgs, err := l.loader.Configs(ctx)
	if err != nil {
		return nil, err
	}
	type minimalActorInfo struct {
		Addresses []struct {
			MinerAddresses []string
		}
	}
	confNameToAddr := map[address.Address][]string{} // address -> config names
	for _, cfg := range cfgs {
		var info minimalActorInfo
		if err := toml.Unmarshal([]byte(cfg.Config), &info); err != nil {
			return nil, xerrors.Errorf("unmarshaling actor info: %s", err)
		}
		for _, ass := range info.Addresses {
			for _, addr := range ass.MinerAddresses {
				a, err := address.NewFromString(addr)
				if err != nil {
					return nil, err
				}
				confNameToAddr[a] = append(confNameToAddr[a], cfg.Title)
			}
		}
	}
	var out []*model.Actor
	for addr, names := range confNameToAddr {
		out = append(out, &model.Actor{
			Address: types.Address{Address: addr},
			Layers:  names,
		})
	}
	return out, nil
}

func (l *ActorLoaderImpl) Actor(ctx context.Context, address types.Address) (*model.Actor, error) {
	cfgs, err := l.loader.Configs(ctx)
	if err != nil {
		return nil, err
	}
	var layers []string
	type minimalActorInfo struct {
		Addresses []struct {
			MinerAddresses []string
		}
	}
loopCfg:
	for _, cfg := range cfgs {
		var info minimalActorInfo
		if err := toml.Unmarshal([]byte(cfg.Config), &info); err != nil {
			return nil, xerrors.Errorf("unmarshaling actor info: %s", err)
		}
		for _, ass := range info.Addresses {
			for _, addr := range ass.MinerAddresses {
				if addr == address.String() {
					layers = append(layers, cfg.Title)
					continue loopCfg
				}
			}
		}
	}

	return &model.Actor{
		Address: address,
		Layers:  layers,
	}, nil
}
