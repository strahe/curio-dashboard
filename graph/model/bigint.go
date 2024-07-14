package model

import (
	"fmt"
	"io"
	"math/big"
)

type BigInt struct {
	*big.Int
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (b *BigInt) UnmarshalGQL(v interface{}) error {
	switch value := v.(type) {
	case string:
		v := big.Int{}
		v.SetString(value, 10)
		*b = BigInt{&v}
	default:
		return fmt.Errorf("invalid bigint")
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (b BigInt) MarshalGQL(w io.Writer) {
	_, _ = w.Write([]byte(b.String()))
}
