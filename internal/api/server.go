package api

import (
	genapi "bank-app/internal/api/gen"
	"bank-app/internal/db"
)

// will be used to create http.Handler using chi utilities.
//
// implements all handlers for this REST endpoint.
type Server struct {
	repo db.Repository
}

type OptFunc func(*Server)

func NewServer(r db.Repository) *Server {
	s := &Server{repo: r}

	return s
}

var _ genapi.ServerInterface = &Server{}
