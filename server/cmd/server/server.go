package main

import (
	"context"
	"fmt"
	"github.com/Destaby/architecture-lab3/server/system"
	"net/http"
)

type HttpPortNumber int

// SystemApiServer configures necessary handlers and starts listening on a configured port.
type SystemApiServer struct {
	Port HttpPortNumber

	SystemHandler system.HttpHandlerFunc

	server *http.Server
}

// Start will set all handlers and start listening.
// If this methods succeeds, it does not return until server is shut down.
// Returned error will never be nil.
func (s *SystemApiServer) Start() error {
	if s.SystemHandler == nil {
		return fmt.Errorf("channels HTTP handler is not defined - cannot start")
	}
	if s.Port == 0 {
		return fmt.Errorf("port is not defined")
	}

	handler := new(http.ServeMux)
	handler.HandleFunc("/balancers", s.SystemHandler)

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: handler,
	}

	return s.server.ListenAndServe()
}

// Stops will shut down previously started HTTP server.
func (s *SystemApiServer) Stop() error {
	if s.server == nil {
		return fmt.Errorf("server was not started")
	}
	return s.server.Shutdown(context.Background())
}