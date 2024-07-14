package model

import (
	"encoding/base64"
	"fmt"
	"io"
)

type ByteArray []byte

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (b *ByteArray) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("ByteArray must be a string")
	}

	arr, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return fmt.Errorf("ByteArray must be valid base64: %w", err)
	}

	*b = arr
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (b ByteArray) MarshalGQL(w io.Writer) {
	jsonStr := fmt.Sprintf(`"%s"`, base64.StdEncoding.EncodeToString(b))
	_, _ = w.Write([]byte(jsonStr))
}
