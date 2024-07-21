package loaders

import (
	"context"
	"time"

	"github.com/strahe/curio-dashboard/graph/model"
	model2 "github.com/strahe/curio-dashboard/model"
)

func (l *Loader) StoragePaths(ctx context.Context) ([]*model.StoragePath, error) {
	var m []*model.StoragePath
	if err := l.db.Select(ctx, &m, "SELECT * FROM storage_path"); err != nil {
		return nil, err
	}
	for _, p := range m {
		if p.CanStore && p.CanSeal {
			p.Type = model.StorageTypeHybrid
		} else if p.CanSeal {
			p.Type = model.StorageTypeSeal
		} else if p.CanStore {
			p.Type = model.StorageTypeStore
		} else {
			p.Type = model.StorageTypeReadonly
		}
		p.Used = p.Capacity - p.Available - p.Reserved // todo: debug this
	}

	return m, nil
}

func (l *Loader) StorageStats(ctx context.Context) ([]*model.StorageStats, error) {
	paths, err := l.StoragePaths(ctx)
	if err != nil {
		return nil, err
	}

	statsMap := make(map[model.StorageType]*model.StorageStats)
	var countFor = func(group *model.StorageStats, path *model.StoragePath) {
		group.TotalAvailable += path.Available
		group.TotalCapacity += path.Capacity
		group.TotalFsAvailable += path.FsAvailable
		group.TotalUsed += path.Used
		group.TotalReserved += path.Reserved
	}

	for _, path := range paths {
		var storageType model.StorageType
		if path.CanSeal && path.CanStore {
			storageType = model.StorageTypeHybrid
		} else if path.CanStore {
			storageType = model.StorageTypeStore
		} else if path.CanSeal {
			storageType = model.StorageTypeSeal
		} else {
			storageType = model.StorageTypeReadonly
		}

		group, exists := statsMap[storageType]
		if !exists {
			group = &model.StorageStats{Type: storageType}
			statsMap[storageType] = group
		}

		countFor(group, path)
	}

	// Convert the map to a slice
	var res []*model.StorageStats
	for _, group := range statsMap {
		res = append(res, group)
	}

	return res, nil
}

func (l *Loader) StorageUsages(_ context.Context, storageID *string, start, end time.Time) ([]*model2.StorageUsage, error) {
	var args []interface{}
	args = append(args, start, end)
	var sql string
	if storageID != nil {
		// Aggregate data by hour for a specific storageID
		sql = `SELECT to_char(time, 'YYYY-MM-DD HH24:00:00') AS time, storage_id, 
                SUM(available) AS available, SUM(fs_available) AS fs_available, 
                SUM(reserved) AS reserved, SUM(used) AS used
                FROM storage_usages 
                WHERE storage_id = ? AND time BETWEEN ? AND ? 
                GROUP BY time,storage_id`
		args = append([]interface{}{*storageID}, args...)
	} else {
		sql = `SELECT to_char(time, 'YYYY-MM-DD HH24:00:00') AS time, storage_id, 
                SUM(available) AS available, SUM(fs_available) AS fs_available, 
                SUM(reserved) AS reserved, SUM(used) AS used
                FROM storage_usages 
                WHERE time BETWEEN ? AND ? 
                GROUP BY time,storage_id
                ORDER BY time`
	}

	type tempStorageUsage struct {
		Time        string
		StorageID   string
		Available   int
		FsAvailable int
		Reserved    int
		Used        int
	}

	var tempUsages []*tempStorageUsage

	err := l.appDB.Raw(sql, args...).Scan(&tempUsages).Error
	if err != nil {
		return nil, err
	}
	var usages []*model2.StorageUsage
	const layout = "2006-01-02 15:04:05"

	for _, tu := range tempUsages {
		parsedTime, err := time.Parse(layout, tu.Time)
		if err != nil {
			return nil, err // Handle parsing error
		}
		usages = append(usages, &model2.StorageUsage{
			Time:        parsedTime,
			StorageID:   tu.StorageID,
			Available:   tu.Available,
			FsAvailable: tu.FsAvailable,
			Reserved:    tu.Reserved,
			Used:        tu.Used,
		})
	}

	return usages, nil
}
