package server

import (
	"net/http"
	"time"

	"github.com/Asemokamichi/Forum/internal/delivery"
)

// hjba
type Server struct {
	S *http.Server
}

func NewServer(handler *delivery.Handler) *Server {
	mux := http.NewServeMux()
	handler.InitRoutes(mux)
	return &Server{
		S: &http.Server{
			Addr:           ":8080",
			Handler:        mux,
			MaxHeaderBytes: 1 << 20,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
		},
	}
}
