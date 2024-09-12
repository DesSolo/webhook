package responser

import (
	"net/http"
	"webhook/internal/entities"
)

type Responser interface {
	Response(w http.ResponseWriter, r *entities.Request) error
}

type ResponserFunc func(w http.ResponseWriter, r *entities.Request) error

func (f ResponserFunc) Response(w http.ResponseWriter, r *entities.Request) error {
	return f(w, r)
}