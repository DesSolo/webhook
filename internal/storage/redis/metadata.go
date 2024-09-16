package redis

import (
	"encoding/json"
	"fmt"
	"webhook/internal/responser"
)

// metadata is internal entity for serialization
type metadata struct {
	Kind      string
	Responser responser.Responser
}

// newMetadata constructor for build metadata from responser
func newMetadata(r responser.Responser) *metadata {
	return &metadata{
		Kind:      r.Kind(),
		Responser: r,
	}
}

// MarshalBinary marshal
func (m *metadata) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

// UnmarshalBinary unmarshal
func (m *metadata) UnmarshalBinary(data []byte) error {
	var alias struct {
		Kind      string
		Responser json.RawMessage
	}

	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}

	r, ok := responser.Get(alias.Kind)
	if !ok {
		return fmt.Errorf("unknown responser kind: %s", alias.Kind)
	}

	if err := r.UnmarshalBinary(alias.Responser); err != nil {
		return err
	}

	m.Kind = alias.Kind
	m.Responser = r

	return nil
}
