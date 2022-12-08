package server

import (
	"net/http"
	"time"

	"github.com/Asemokamichi/Forum/internal/delivery"
)

// type App struct {
// 	db         *sql.DB
// 	httpServer *http.Server
// }

// func NewApp() *App {
// 	return &App{}
// }

// func (a *App) Run() error {
// 	return a.httpServer.ListenAndServe()
// }

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
