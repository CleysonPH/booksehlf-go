package dto

import (
	"encoding/json"
	"time"
)

type NullString struct {
	Value string
}

func (n *NullString) IsValid() bool {
	return n.Value != ""
}

func (n NullString) MarshalJSON() ([]byte, error) {
	if n.IsValid() {
		return json.Marshal(n.Value)
	}
	return []byte("null"), nil
}

type NullInt32 struct {
	Value int32
}

func (n *NullInt32) IsValid() bool {
	return n.Value != 0
}

func (n NullInt32) MarshalJSON() ([]byte, error) {
	if n.IsValid() {
		return json.Marshal(n.Value)
	}
	return []byte("null"), nil
}

type NullTime struct {
	Value time.Time
}

func (n *NullTime) IsValid() bool {
	return !n.Value.IsZero()
}

func (n NullTime) MarshalJSON() ([]byte, error) {
	if n.IsValid() {
		return json.Marshal(n.Value)
	}
	return []byte("null"), nil
}
