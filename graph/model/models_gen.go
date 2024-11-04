// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/strahe/curio-dashboard/types"
)

type ActorDeadline struct {
	Empty      bool `json:"empty"`
	Current    bool `json:"current"`
	Proven     bool `json:"proven"`
	PartFaulty bool `json:"partFaulty"`
	Faulty     bool `json:"faulty"`
}

type Alert struct {
	ID          int    `json:"id"`
	MachineName string `json:"machineName"`
	Message     string `json:"message"`
}

type GaugeCountValue struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

type MachineDetail struct {
	ID          int       `json:"id"`
	MachineName string    `json:"machineName"`
	Tasks       string    `json:"tasks"`
	TasksArray  []string  `json:"tasksArray"`
	Layers      string    `json:"layers"`
	LayersArray []string  `json:"layersArray"`
	StartupTime time.Time `json:"startupTime"`
	Miners      string    `json:"miners"`
	MinersArray []string  `json:"minersArray"`
	MachineID   int       `json:"machineId"`
}

type MachineMetrics struct {
	CPUUsage                   float64            `json:"cpuUsage"`
	GpuUsage                   float64            `json:"gpuUsage"`
	RAMUsage                   float64            `json:"ramUsage"`
	ActiveTasks                []*GaugeCountValue `json:"activeTasks"`
	AddedTasks                 []*GaugeCountValue `json:"addedTasks"`
	TasksCompleted             []*GaugeCountValue `json:"tasksCompleted"`
	TasksStarted               []*GaugeCountValue `json:"tasksStarted"`
	GoRoutines                 int                `json:"goRoutines"`
	GoVersion                  string             `json:"goVersion"`
	GoThreads                  int                `json:"goThreads"`
	ProcessCPUSecondsTotal     int                `json:"processCpuSecondsTotal"`
	ProcessStartTimeSeconds    int                `json:"processStartTimeSeconds"`
	ProcessVirtualMemoryBytes  int                `json:"processVirtualMemoryBytes"`
	ProcessResidentMemoryBytes int                `json:"processResidentMemoryBytes"`
	ProcessOpenFds             int                `json:"processOpenFds"`
	ProcessMaxFds              int                `json:"processMaxFds"`
}

type MetricsActiveTask struct {
	Name   string      `json:"name"`
	Series [][]float64 `json:"series"`
}

type MinerBeneficiaryTerm struct {
	Quota      types.BigInt `json:"quota"`
	UsedQuota  types.BigInt `json:"usedQuota"`
	Expiration int          `json:"expiration"`
}

type MinerInfo struct {
	Owner                      types.Address                  `json:"owner"`
	Worker                     types.Address                  `json:"worker"`
	NewWorker                  types.Address                  `json:"newWorker"`
	ControlAddresses           []*types.Address               `json:"controlAddresses"`
	WorkerChangeEpoch          int                            `json:"workerChangeEpoch"`
	PeerID                     string                         `json:"peerId"`
	MultiAddrs                 []string                       `json:"multiAddrs"`
	WindowPoStProofType        int                            `json:"windowPoStProofType"`
	SectorSize                 int                            `json:"sectorSize"`
	WindowPoStPartitionSectors int                            `json:"windowPoStPartitionSectors"`
	ConsensusFaultElapsed      int                            `json:"consensusFaultElapsed"`
	PendingOwnerAddress        types.Address                  `json:"pendingOwnerAddress"`
	Beneficiary                types.Address                  `json:"beneficiary"`
	BeneficiaryTerm            *MinerBeneficiaryTerm          `json:"beneficiaryTerm"`
	PendingBeneficiaryChange   *MinerPendingBeneficiaryChange `json:"pendingBeneficiaryChange"`
}

type MinerPendingBeneficiaryChange struct {
	NewBeneficiary        types.Address `json:"newBeneficiary"`
	NewQuota              types.BigInt  `json:"newQuota"`
	NewExpiration         int           `json:"newExpiration"`
	ApprovedByBeneficiary bool          `json:"approvedByBeneficiary"`
	ApprovedByNominee     bool          `json:"approvedByNominee"`
}

type MinerPower struct {
	ID          string      `json:"id"`
	MinerPower  *PowerClaim `json:"minerPower"`
	TotalPower  *PowerClaim `json:"totalPower"`
	HasMinPower bool        `json:"hasMinPower"`
}

type MiningCount struct {
	Include int `json:"include"`
	Exclude int `json:"exclude"`
}

type MiningSummaryDay struct {
	Day      time.Time     `json:"day"`
	Miner    types.ActorID `json:"miner"`
	WonBlock int           `json:"wonBlock"`
}

type Mutation struct {
}

type NodeInfo struct {
	ID        string   `json:"id"`
	Address   string   `json:"address"`
	Layers    []string `json:"layers"`
	Reachable bool     `json:"reachable"`
	SyncState string   `json:"syncState"`
	Version   string   `json:"version"`
}

type OpenSectorPiece struct {
	SpID                          types.ActorID `json:"spID"`
	SectorNumber                  int           `json:"sectorNumber"`
	PieceIndex                    int           `json:"pieceIndex"`
	PieceCid                      string        `json:"pieceCID"`
	PieceSize                     int           `json:"pieceSize"`
	DataURL                       string        `json:"dataURL"`
	DataHeaders                   string        `json:"dataHeaders"`
	DataRawSize                   int           `json:"dataRawSize"`
	DataDeleteOnFinalize          bool          `json:"dataDeleteOnFinalize"`
	F05PublishCid                 *string       `json:"f05PublishCID"`
	F05DealID                     *int          `json:"f05DealID"`
	F05DealProposal               *string       `json:"f05DealProposal"`
	F05DealStartEpoch             *int          `json:"f05DealStartEpoch"`
	F05DealEndEpoch               *int          `json:"f05DealEndEpoch"`
	DirectStartEpoch              *int          `json:"directStartEpoch"`
	DirectEndEpoch                *int          `json:"directEndEpoch"`
	DirectPieceActivationManifest *string       `json:"directPieceActivationManifest"`
	CreatedAt                     time.Time     `json:"createdAt"`
	IsSnap                        bool          `json:"isSnap"`
}

type Porep struct {
	ID                       string          `json:"id"`
	SpID                     types.ActorID   `json:"spId"`
	SectorNumber             int             `json:"sectorNumber"`
	CreateTime               time.Time       `json:"createTime"`
	RegSealProof             int             `json:"regSealProof"`
	TicketEpoch              *int            `json:"ticketEpoch"`
	TicketValue              types.ByteArray `json:"ticketValue"`
	TaskIDSdr                *int            `json:"taskIdSdr"`
	AfterSdr                 bool            `json:"afterSdr"`
	TreeDCid                 *string         `json:"treeDCid"`
	TaskIDTreeD              *int            `json:"taskIdTreeD"`
	AfterTreeD               bool            `json:"afterTreeD"`
	TaskIDTreeC              *int            `json:"taskIdTreeC"`
	AfterTreeC               bool            `json:"afterTreeC"`
	TreeRCid                 *string         `json:"treeRCid"`
	TaskIDTreeR              *int            `json:"taskIdTreeR"`
	AfterTreeR               bool            `json:"afterTreeR"`
	PrecommitMsgCid          *string         `json:"precommitMsgCid"`
	TaskIDPrecommitMsg       *int            `json:"taskIdPrecommitMsg"`
	AfterPrecommitMsg        bool            `json:"afterPrecommitMsg"`
	SeedEpoch                *int            `json:"seedEpoch"`
	PrecommitMsgTsk          types.ByteArray `json:"precommitMsgTsk"`
	AfterPrecommitMsgSuccess bool            `json:"afterPrecommitMsgSuccess"`
	SeedValue                types.ByteArray `json:"seedValue"`
	TaskIDPorep              *int            `json:"taskIdPorep"`
	PorepProof               types.ByteArray `json:"porepProof"`
	AfterPorep               bool            `json:"afterPorep"`
	TaskIDFinalize           *int            `json:"taskIdFinalize"`
	AfterFinalize            bool            `json:"afterFinalize"`
	TaskIDMoveStorage        *int            `json:"taskIdMoveStorage"`
	AfterMoveStorage         bool            `json:"afterMoveStorage"`
	CommitMsgCid             *string         `json:"commitMsgCid"`
	TaskIDCommitMsg          *int            `json:"taskIdCommitMsg"`
	AfterCommitMsg           bool            `json:"afterCommitMsg"`
	CommitMsgTsk             types.ByteArray `json:"commitMsgTsk"`
	AfterCommitMsgSuccess    bool            `json:"afterCommitMsgSuccess"`
	Failed                   bool            `json:"failed"`
	FailedAt                 *time.Time      `json:"failedAt"`
	FailedReason             string          `json:"failedReason"`
	FailedReasonMsg          string          `json:"failedReasonMsg"`
	TaskIDSynth              *int            `json:"taskIdSynth"`
	AfterSynth               bool            `json:"afterSynth"`
	UserSectorDurationEpochs *int            `json:"userSectorDurationEpochs"`
	Status                   PorepStatus     `json:"status"`
	CurrentTask              *Task           `json:"currentTask"`
}

type PowerClaim struct {
	RawBytePower    *types.BigInt `json:"rawBytePower"`
	QualityAdjPower *types.BigInt `json:"qualityAdjPower"`
}

type Query struct {
}

type SectorMetaPiece struct {
	SpID              types.ActorID `json:"spID"`
	SectorNum         int           `json:"sectorNum"`
	PieceNum          int           `json:"pieceNum"`
	PieceCid          string        `json:"pieceCID"`
	PieceSize         int           `json:"pieceSize"`
	RequestedKeepData bool          `json:"requestedKeepData"`
	RawDataSize       *int          `json:"rawDataSize"`
	StartEpoch        *int          `json:"startEpoch"`
	OrigEndEpoch      *int          `json:"origEndEpoch"`
	F05DealID         *int          `json:"f05DealID"`
	DdoPam            types.JSONB   `json:"ddoPam"`
	F05DealProposal   types.JSONB   `json:"f05DealProposal"`
}

type StorageLiveness struct {
	StorageID      string     `json:"storageId"`
	URL            string     `json:"url"`
	LastChecked    time.Time  `json:"lastChecked"`
	LastLive       *time.Time `json:"lastLive"`
	LastDead       *time.Time `json:"lastDead"`
	LastDeadReason *string    `json:"lastDeadReason"`
}

type StorageStats struct {
	Type             StorageType `json:"type"`
	TotalCapacity    int         `json:"totalCapacity"`
	TotalAvailable   int         `json:"totalAvailable"`
	TotalUsed        int         `json:"totalUsed"`
	TotalReserved    int         `json:"totalReserved"`
	TotalFsAvailable int         `json:"totalFsAvailable"`
}

type StorageUsage struct {
	Time        time.Time `json:"time"`
	Available   int       `json:"available"`
	Used        int       `json:"used"`
	Reserved    int       `json:"reserved"`
	FsAvailable int       `json:"fsAvailable"`
}

type Subscription struct {
}

type TaskNameAggregate struct {
	Name    string `json:"name"`
	Total   int    `json:"total"`
	Success int    `json:"success"`
	Failure int    `json:"failure"`
}

type TaskStats struct {
	Name    string `json:"name"`
	Total   int    `json:"total"`
	Success int    `json:"success"`
	Failure int    `json:"failure"`
}

type TaskSummary struct {
	Name       string `json:"name"`
	TrueCount  int    `json:"trueCount"`
	FalseCount int    `json:"falseCount"`
	TotalCount int    `json:"totalCount"`
}

type TaskSummaryDay struct {
	Day        time.Time `json:"day"`
	TrueCount  int       `json:"trueCount"`
	FalseCount int       `json:"falseCount"`
	TotalCount int       `json:"totalCount"`
}

type PorepStatus string

const (
	PorepStatusSdr              PorepStatus = "SDR"
	PorepStatusTreeD            PorepStatus = "TreeD"
	PorepStatusTreeRc           PorepStatus = "TreeRC"
	PorepStatusSynthetic        PorepStatus = "Synthetic"
	PorepStatusPreCommitMsg     PorepStatus = "PreCommitMsg"
	PorepStatusPreCommitMsgWait PorepStatus = "PreCommitMsgWait"
	PorepStatusWaitSeed         PorepStatus = "WaitSeed"
	PorepStatusPoRep            PorepStatus = "PoRep"
	PorepStatusClearCache       PorepStatus = "ClearCache"
	PorepStatusMoveStorage      PorepStatus = "MoveStorage"
	PorepStatusCommitMsg        PorepStatus = "CommitMsg"
	PorepStatusCommitMsgWait    PorepStatus = "CommitMsgWait"
	PorepStatusFailed           PorepStatus = "Failed"
	PorepStatusSuccess          PorepStatus = "Success"
	PorepStatusOnChain          PorepStatus = "OnChain"
	PorepStatusActive           PorepStatus = "Active"
	PorepStatusUnknown          PorepStatus = "Unknown"
)

var AllPorepStatus = []PorepStatus{
	PorepStatusSdr,
	PorepStatusTreeD,
	PorepStatusTreeRc,
	PorepStatusSynthetic,
	PorepStatusPreCommitMsg,
	PorepStatusPreCommitMsgWait,
	PorepStatusWaitSeed,
	PorepStatusPoRep,
	PorepStatusClearCache,
	PorepStatusMoveStorage,
	PorepStatusCommitMsg,
	PorepStatusCommitMsgWait,
	PorepStatusFailed,
	PorepStatusSuccess,
	PorepStatusOnChain,
	PorepStatusActive,
	PorepStatusUnknown,
}

func (e PorepStatus) IsValid() bool {
	switch e {
	case PorepStatusSdr, PorepStatusTreeD, PorepStatusTreeRc, PorepStatusSynthetic, PorepStatusPreCommitMsg, PorepStatusPreCommitMsgWait, PorepStatusWaitSeed, PorepStatusPoRep, PorepStatusClearCache, PorepStatusMoveStorage, PorepStatusCommitMsg, PorepStatusCommitMsgWait, PorepStatusFailed, PorepStatusSuccess, PorepStatusOnChain, PorepStatusActive, PorepStatusUnknown:
		return true
	}
	return false
}

func (e PorepStatus) String() string {
	return string(e)
}

func (e *PorepStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PorepStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PorepStatus", str)
	}
	return nil
}

func (e PorepStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type StorageType string

const (
	StorageTypeHybrid   StorageType = "Hybrid"
	StorageTypeSeal     StorageType = "Seal"
	StorageTypeStore    StorageType = "Store"
	StorageTypeReadonly StorageType = "Readonly"
)

var AllStorageType = []StorageType{
	StorageTypeHybrid,
	StorageTypeSeal,
	StorageTypeStore,
	StorageTypeReadonly,
}

func (e StorageType) IsValid() bool {
	switch e {
	case StorageTypeHybrid, StorageTypeSeal, StorageTypeStore, StorageTypeReadonly:
		return true
	}
	return false
}

func (e StorageType) String() string {
	return string(e)
}

func (e *StorageType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = StorageType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid StorageType", str)
	}
	return nil
}

func (e StorageType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TaskHistoriesAggregateInterval string

const (
	TaskHistoriesAggregateIntervalDay  TaskHistoriesAggregateInterval = "day"
	TaskHistoriesAggregateIntervalHour TaskHistoriesAggregateInterval = "hour"
)

var AllTaskHistoriesAggregateInterval = []TaskHistoriesAggregateInterval{
	TaskHistoriesAggregateIntervalDay,
	TaskHistoriesAggregateIntervalHour,
}

func (e TaskHistoriesAggregateInterval) IsValid() bool {
	switch e {
	case TaskHistoriesAggregateIntervalDay, TaskHistoriesAggregateIntervalHour:
		return true
	}
	return false
}

func (e TaskHistoriesAggregateInterval) String() string {
	return string(e)
}

func (e *TaskHistoriesAggregateInterval) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TaskHistoriesAggregateInterval(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TaskHistoriesAggregateInterval", str)
	}
	return nil
}

func (e TaskHistoriesAggregateInterval) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
