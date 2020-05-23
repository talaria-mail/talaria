package web

import (
	"context"
	"net/http"
	"time"
)

//go:generate go-bindata -prefix "../../../frontend/dist/" -fs ../../../frontend/dist/...

// fallbackFS is a http.FileSystem that falls back to a different path when
// opening an attempted path fails. Useful for serving single path
// applications.
type fallbackFS struct {
	path string
	next http.FileSystem
}

func (f fallbackFS) Open(name string) (http.File, error) {
	file, err := f.next.Open(name)
	if err != nil {
		return f.next.Open(f.path)
	}
	return file, nil
}

type Config struct {
	Addr string
}

type Server struct {
	server *http.Server
}

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

func (s Server) Run() error {
	return s.server.ListenAndServe()
}

func (s Server) Shutdown(error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	server.Shutdown(ctx)
}
