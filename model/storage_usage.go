package model

import "time"

type StorageUsage struct {
	Time        time.Time `gorm:"primarykey;autoIncrement:false"`
	StorageID   string    `gorm:"primarykey;autoIncrement:false"`
	Available   int
	FsAvailable int
	Reserved    int
	Used        int
}
