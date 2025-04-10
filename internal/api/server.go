package api

import (
	genapi "bank-app/internal/api/gen"
	repo "bank-app/internal/repository"
)

// will be used to create http.Handler using chi utilities.
//
// implements all handlers for this REST endpoint.
type Server struct {
	repo repo.Repository
}

type OptFunc func(*Server)

func NewServer(r repo.Repository) *Server {
	s := &Server{repo: r}

	return s
}

var _ genapi.ServerInterface = &Server{}
