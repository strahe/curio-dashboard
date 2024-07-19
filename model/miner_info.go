package model

import (
	"github.com/strahe/curio-dashboard/types"
	"time"
)

type MinerInfo struct {
	Time             time.Time     `gorm:"primarykey;autoIncrement:false"`
	Address          types.Address `gorm:"primarykey;autoIncrement:false"`
	WorkerAddress    types.Address
	RawBytePower     types.BigInt
	QAPower          types.BigInt
	Balance          types.BigInt
	AvailableBalance types.BigInt
	WorkerBalance    types.BigInt
}
