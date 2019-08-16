package handler

import (
	"log"
	"net/http"

	"github.com/Xuanwo/go-mod-redirect/config"
)

// Handler is a normal http handler.
type Handler struct {
	*config.Service
}

// New will create a new Handler.
func New(c *config.Service) (h *Handler, err error) {
	h = &Handler{c}
	return
}

// ServeHTTP implement http.Handler.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Receive request %s", r.URL.Path)
	switch r.URL.Path {
	case "/":
		h.serveIndex(w, r)
	default:
		h.serveImport(w, r)
	}
}

func (h *Handler) serveIndex(w http.ResponseWriter, r *http.Request) {
	paths := make([]string, len(h.Paths))
	for i, v := range h.Paths {
		paths[i] = h.Host + v.Path
	}
	if err := indexTmpl.Execute(w, struct {
		Host  string
		Paths []string
	}{
		Host:  h.Host,
		Paths: paths,
	}); err != nil {
		http.Error(w, "cannot render the page", http.StatusInternalServerError)
	}
}

func (h *Handler) serveImport(w http.ResponseWriter, r *http.Request) {
	m, subpath := h.Find(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return
	}

	if err := importTmpl.Execute(w, struct {
		Import  string
		Subpath string
		Repo    string
		Display string
		VCS     string
	}{
		Import:  h.Host + m.Path,
		Subpath: subpath,
		Repo:    m.Repo,
		VCS:     m.VCS,
	}); err != nil {
		http.Error(w, "cannot render the page", http.StatusInternalServerError)
	}
}
