package api

import (
	genapi "bank-app/internal/api/gen"
)

// will be used to create http.Handler using chi utilities.
//
// implements all handlers for this REST endpoint.
type Server struct{}

func NewServer() *Server {
	return &Server{}
}

var _ genapi.ServerInterface = &Server{}
