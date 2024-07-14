package model

import (
	"fmt"
	"github.com/filecoin-project/go-address"
	"io"
)

type ActorID uint64

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (b *ActorID) UnmarshalGQL(v interface{}) error {
	switch value := v.(type) {
	case int:
		*b = ActorID(value)
	case string:
		addr, err := address.NewFromString(value)
		if err != nil {
			return fmt.Errorf("invalid actor id: %w", err)
		}
		aid, err := address.IDFromAddress(addr)
		if err != nil {
			return fmt.Errorf("invalid actor id: %w", err)
		}
		*b = ActorID(aid)
	default:
		return fmt.Errorf("invalid actor id type")
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (b ActorID) MarshalGQL(w io.Writer) {
	addr, _ := address.NewIDAddress(uint64(b))
	_, _ = w.Write([]byte(`"` + addr.String() + `"`))
}
