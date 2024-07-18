package model

import "time"

type TaskHistoryAggregates struct {
	Time    time.Time `gorm:"primarykey;autoIncrement:false"`
	Task    string    `gorm:"primarykey;autoIncrement:false"`
	Total   int
	Failure int
	Success int
}
