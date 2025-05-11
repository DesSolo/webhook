package responser

import (
	"net/http"

	"webhook/internal/pkg/entities"
)

// DumpRestorer ...
type DumpRestorer interface {
	Kind() string
	MarshalBinary() ([]byte, error)
	UnmarshalBinary([]byte) error
}

// Responser ...
type Responser interface {
	Response(w http.ResponseWriter, r *entities.Request) error
	DumpRestorer
}
