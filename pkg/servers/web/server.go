package web

import (
	"context"
	"net/http"
	"time"
)

// Config is the data needed to configure a web client server. Server binds to
// Addr and serves the filesystem FileSystem.
type Config struct {
	Addr       string
	FileSystem http.FileSystem
}

// Server is a web-client server
type Server struct {
	server *http.Server
}

// New generates a new web client server. Start with Server.Run() and shutdown
// with Server.Shutdown().
func New(conf Config) Server {
	mux := http.NewServeMux()
	fs := AssetFile()
	mux.Handle("/", http.FileServer(fallbackFS{"index.html", fs}))

	server := &http.Server{
		Addr:    conf.Addr,
		Handler: mux,
	}
	return Server{server}
}

// Run starts a web client server. Server binds to address from config and serve
// filesystem provided.
func (s Server) Run() error {
	return s.server.ListenAndServe()
}

// Shutdown gracefully stops the server
func (s Server) Shutdown(error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	s.server.Shutdown(ctx)
}
