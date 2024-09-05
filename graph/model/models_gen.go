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

type MetricsActiveTask struct {
	Name   string      `json:"name"`
	Series [][]float64 `json:"series"`
}

type MinerBeneficiaryTerm struct {
	Quota      types.BigInt `json:"quota"`
	UsedQuota  int          `json:"usedQuota"`
	Expiration int          `json:"expiration"`
}

type MinerInfo struct {
	Owner                      types.Address                  `json:"owner"`
	Worker                     types.Address                  `json:"worker"`
	NewWorker                  types.Address                  `json:"newWorker"`
	ControlAddresses           []*types.Address               `json:"controlAddresses"`
	WorkerChangeEpoch          int                            `json:"workerChangeEpoch"`
	PeerID                     *string                        `json:"peerId"`
	MultiAddrs                 []string                       `json:"multiAddrs"`
	WindowPoStProofType        int                            `json:"windowPoStProofType"`
	SectorSize                 int                            `json:"sectorSize"`
	WindowPoStPartitionSectors int                            `json:"windowPoStPartitionSectors"`
	ConsensusFaultElapsed      int                            `json:"consensusFaultElapsed"`
	PendingOwnerAddress        *types.Address                 `json:"pendingOwnerAddress"`
	Beneficiary                types.Address                  `json:"beneficiary"`
	BeneficiaryTerm            *MinerBeneficiaryTerm          `json:"beneficiaryTerm"`
	PendingBeneficiaryChange   *MinerPendingBeneficiaryChange `json:"pendingBeneficiaryChange"`
}

type MinerPendingBeneficiaryChange struct {
	NewBeneficiary        *types.Address `json:"newBeneficiary"`
	NewQuota              types.BigInt   `json:"newQuota"`
	NewExpiration         int            `json:"newExpiration"`
	ApprovedByBeneficiary bool           `json:"approvedByBeneficiary"`
	ApprovedByNominee     bool           `json:"approvedByNominee"`
}

type MinerPower struct {
	ID          string      `json:"id"`
	MinerPower  *PowerClaim `json:"minerPower"`
	TotalPower  *PowerClaim `json:"totalPower"`
	HasMinPower bool        `json:"hasMinPower"`
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

type Pipeline struct {
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
	Status                   PipelineStatus  `json:"status"`
	CurrentTask              *Task           `json:"currentTask"`
}

type PowerClaim struct {
	RawBytePower    *types.BigInt `json:"rawBytePower"`
	QualityAdjPower *types.BigInt `json:"qualityAdjPower"`
}

type Query struct {
}

type SectorLocation struct {
	MinerID        types.ActorID `json:"minerId"`
	SectorNum      int           `json:"sectorNum"`
	SectorFiletype int           `json:"sectorFiletype"`
	StorageID      string        `json:"storageId"`
	IsPrimary      *bool         `json:"isPrimary"`
	ReadTs         *string       `json:"readTs"`
	ReadRefs       int           `json:"readRefs"`
	WriteTs        *string       `json:"writeTs"`
	WriteLockOwner *string       `json:"writeLockOwner"`
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

type StoragePath struct {
	ID            string      `json:"id"`
	StorageID     string      `json:"storageId"`
	Type          StorageType `json:"type"`
	Urls          string      `json:"urls"`
	Weight        int         `json:"weight"`
	MaxStorage    int         `json:"maxStorage"`
	CanSeal       bool        `json:"canSeal"`
	CanStore      bool        `json:"canStore"`
	Groups        *string     `json:"groups"`
	AllowTo       *string     `json:"allowTo"`
	AllowTypes    *string     `json:"allowTypes"`
	DenyTypes     *string     `json:"denyTypes"`
	Capacity      int         `json:"capacity"`
	Available     int         `json:"available"`
	FsAvailable   int         `json:"fsAvailable"`
	Reserved      int         `json:"reserved"`
	Used          int         `json:"used"`
	LastHeartbeat time.Time   `json:"lastHeartbeat"`
	HeartbeatErr  *string     `json:"heartbeatErr"`
	AllowMiners   string      `json:"allowMiners"`
	DenyMiners    string      `json:"denyMiners"`
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

type TaskHistory struct {
	ID                     int       `json:"id"`
	TaskID                 int       `json:"taskId"`
	Name                   string    `json:"name"`
	Posted                 time.Time `json:"posted"`
	WorkStart              time.Time `json:"workStart"`
	WorkEnd                time.Time `json:"workEnd"`
	Result                 bool      `json:"result"`
	Err                    *string   `json:"err"`
	CompletedByHostAndPort string    `json:"completedByHostAndPort"`
}

type TaskNameAggregate struct {
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

type PipelineStatus string

const (
	PipelineStatusSdr              PipelineStatus = "SDR"
	PipelineStatusTreeD            PipelineStatus = "TreeD"
	PipelineStatusTreeRc           PipelineStatus = "TreeRC"
	PipelineStatusSynthetic        PipelineStatus = "Synthetic"
	PipelineStatusPreCommitMsg     PipelineStatus = "PreCommitMsg"
	PipelineStatusPreCommitMsgWait PipelineStatus = "PreCommitMsgWait"
	PipelineStatusWaitSeed         PipelineStatus = "WaitSeed"
	PipelineStatusPoRep            PipelineStatus = "PoRep"
	PipelineStatusClearCache       PipelineStatus = "ClearCache"
	PipelineStatusMoveStorage      PipelineStatus = "MoveStorage"
	PipelineStatusCommitMsg        PipelineStatus = "CommitMsg"
	PipelineStatusCommitMsgWait    PipelineStatus = "CommitMsgWait"
	PipelineStatusFailed           PipelineStatus = "Failed"
	PipelineStatusSuccess          PipelineStatus = "Success"
	PipelineStatusUnknown          PipelineStatus = "Unknown"
)

var AllPipelineStatus = []PipelineStatus{
	PipelineStatusSdr,
	PipelineStatusTreeD,
	PipelineStatusTreeRc,
	PipelineStatusSynthetic,
	PipelineStatusPreCommitMsg,
	PipelineStatusPreCommitMsgWait,
	PipelineStatusWaitSeed,
	PipelineStatusPoRep,
	PipelineStatusClearCache,
	PipelineStatusMoveStorage,
	PipelineStatusCommitMsg,
	PipelineStatusCommitMsgWait,
	PipelineStatusFailed,
	PipelineStatusSuccess,
	PipelineStatusUnknown,
}

func (e PipelineStatus) IsValid() bool {
	switch e {
	case PipelineStatusSdr, PipelineStatusTreeD, PipelineStatusTreeRc, PipelineStatusSynthetic, PipelineStatusPreCommitMsg, PipelineStatusPreCommitMsgWait, PipelineStatusWaitSeed, PipelineStatusPoRep, PipelineStatusClearCache, PipelineStatusMoveStorage, PipelineStatusCommitMsg, PipelineStatusCommitMsgWait, PipelineStatusFailed, PipelineStatusSuccess, PipelineStatusUnknown:
		return true
	}
	return false
}

func (e PipelineStatus) String() string {
	return string(e)
}

func (e *PipelineStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PipelineStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PipelineStatus", str)
	}
	return nil
}

func (e PipelineStatus) MarshalGQL(w io.Writer) {
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
