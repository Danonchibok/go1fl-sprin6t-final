package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Loger  *log.Logger
	Server *http.Server
}

func NewServer(loger *log.Logger) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)

	return &Server{
		Loger: loger,
		Server: &http.Server{
			Addr:         ":8080",
			Handler:      mux,
			ErrorLog:     loger,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
	}
}

func (server *Server) StartServer() {
	err := server.Server.ListenAndServe()
	if err != nil {
		server.Loger.Fatalf("failed to start server: %v", err)
	}
}
