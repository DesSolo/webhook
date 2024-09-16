package responser

import (
	"net/http"
	"webhook/internal/entities"
)

type DumpRestorer interface {
	Kind() string
	MarshalBinary() ([]byte, error)
	UnmarshalBinary([]byte) error
}

type Responser interface {
	Response(w http.ResponseWriter, r *entities.Request) error
	DumpRestorer
}
