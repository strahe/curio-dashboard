package loaders

import (
	"context"

	"github.com/web3tea/curio-dashboard/graph/model"
)

type MachineLoader interface {
	Machine(ctx context.Context, id int) (*model.Machine, error)
	MachineByHostPort(ctx context.Context, hostPort string) (*model.Machine, error)
	Machines(ctx context.Context) ([]*model.Machine, error)
	MachineDetails(ctx context.Context) ([]*model.MachineDetail, error)
	MachineStorages(ctx context.Context, hostPort string) ([]*model.StoragePath, error)
}

type MachineLoaderImpl struct {
	loader *Loader
}

func NewMachineLoader(loader *Loader) MachineLoader {
	return &MachineLoaderImpl{loader}
}

func (l *MachineLoaderImpl) MachineStorages(ctx context.Context, hostPort string) ([]*model.StoragePath, error) {
	var ss []*model.StoragePath
	err := l.loader.db.Select(ctx, &ss, `SELECT
    storage_id,
    urls,
    weight,
    max_storage,
    can_seal,
    can_store,
    groups,
    allow_to,
    allow_types,
    deny_types,
    capacity,
    available,
    fs_available,
    reserved,
    used,
    allow_miners,
    deny_miners,
    last_heartbeat,
    heartbeat_err
FROM storage_path
WHERE urls
          LIKE '%' || $1 || '%'`, hostPort)
	if err != nil {
		return nil, err
	}
	return ss, nil
}

func (l *MachineLoaderImpl) Machine(ctx context.Context, id int) (*model.Machine, error) {
	var out model.Machine
	err := l.loader.db.QueryRow(ctx, `SELECT
    id,
    last_contact,
    host_and_port,
    cpu,
    ram,
    gpu
FROM
    harmony_machines
WHERE id = $1`, id).
		Scan(&out.ID, &out.LastContact, &out.HostAndPort, &out.CPU, &out.RAM, &out.Gpu)
	return &out, err
}

func (l *MachineLoaderImpl) Machines(ctx context.Context) ([]*model.Machine, error) {
	var out []*model.Machine
	if err := l.loader.db.Select(ctx, &out, `SELECT
    id,
    last_contact,
    host_and_port,
    cpu,
    ram,
    gpu
FROM
    harmony_machines`); err != nil {
		return nil, err
	}
	return out, nil
}

func (l *MachineLoaderImpl) MachineByHostPort(ctx context.Context, hostPort string) (*model.Machine, error) {
	var res []*model.Machine
	if err := l.loader.db.Select(ctx, &res, `SELECT
    id,
    last_contact,
    host_and_port,
    cpu,
    ram,
    gpu
FROM
    harmony_machines
WHERE host_and_port = $1`, hostPort); err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, ErrorNotFound
	}
	return res[0], nil
}

func (l *MachineLoaderImpl) MachineDetails(ctx context.Context) ([]*model.MachineDetail, error) {
	var out []*model.MachineDetail
	if err := l.loader.db.Select(ctx, &out, `SELECT
    id,
    tasks,
    layers,
    startup_time,
    miners,
    machine_id,
    machine_name
FROM
    harmony_machine_details
`); err != nil {
		return nil, err
	}
	return out, nil
}
