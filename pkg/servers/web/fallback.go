package web

import (
	"net/http"
)

type fallbackFS struct {
	path string
	next http.FileSystem
}

// WithFallback is an http.FileSystem middleware that serves the file
// at path when the requested file does not exist. Usefule for serving sing page
// applications.
func WithFallback(path string, next http.FileSystem) http.FileSystem {
	return fallbackFS{path, next}
}

func (f fallbackFS) Open(name string) (http.File, error) {
	file, err := f.next.Open(name)
	if err != nil {
		return f.next.Open(f.path)
	}
	return file, nil
}
