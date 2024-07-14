package model

import (
	"fmt"
	"github.com/filecoin-project/go-address"
	"io"
)

type Address struct {
	address.Address
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (b *Address) UnmarshalGQL(v interface{}) error {
	switch value := v.(type) {
	case string:
		addr, err := address.NewFromString(value)
		if err != nil {
			return fmt.Errorf("invalid address: %w", err)
		}
		*b = Address{addr}
	default:
		return fmt.Errorf("invalid address")
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (b Address) MarshalGQL(w io.Writer) {
	_, _ = w.Write([]byte(`"` + b.String() + `"`))
}
