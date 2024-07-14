package model

type Sector struct {
	SpID      ActorID
	SectorNum int
	Meta      *SectorMeta `json:"meta,omitempty"`
}

type SectorMeta struct {
	SpID            ActorID   `json:"spId"`
	SectorNum       int       `json:"sectorNum"`
	RegSealProof    int       `json:"regSealProof"`
	TicketEpoch     int       `json:"ticketEpoch"`
	TicketValue     ByteArray `json:"ticketValue,omitempty"`
	OrigSealedCid   string    `json:"origSealedCid"`
	OrigUnsealedCid string    `json:"origUnsealedCid"`
	CurSealedCid    string    `json:"curSealedCid"`
	CurUnsealedCid  string    `json:"curUnsealedCid"`
	MsgCidPrecommit *string   `json:"msgCidPrecommit,omitempty"`
	MsgCidCommit    *string   `json:"msgCidCommit,omitempty"`
	MsgCidUpdate    *string   `json:"msgCidUpdate,omitempty"`
	SeedEpoch       int       `json:"seedEpoch"`
	SeedValue       ByteArray `json:"seedValue,omitempty"`
}
