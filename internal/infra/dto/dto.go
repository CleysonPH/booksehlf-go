package dto

import (
	"encoding/json"
	"time"
)

type NullString struct {
	Value string
}

func (n NullString) MarshalJSON() ([]byte, error) {
	if n.Value != "" {
		return json.Marshal(n.Value)
	}
	return []byte("null"), nil
}

type NullInt32 struct {
	Value int32
}

func (n NullInt32) MarshalJSON() ([]byte, error) {
	if n.Value != 0 {
		return json.Marshal(n.Value)
	}
	return []byte("null"), nil
}

type NullTime struct {
	Value time.Time
}

func (n NullTime) MarshalJSON() ([]byte, error) {
	if !n.Value.IsZero() {
		return json.Marshal(n.Value)
	}
	return []byte("null"), nil
}
