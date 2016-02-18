package server

import (
	"github.com/werwolfby/movie-sort/links"
	"net/http"
)

type LinksHandlers struct {
	links links.Links
}

func NewLinksHandlers(l links.Links) *LinksHandlers {
	return &LinksHandlers{links: l}
}

func (l LinksHandlers) GetHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeOk(w, l.links.GetLinks())
	})
}
