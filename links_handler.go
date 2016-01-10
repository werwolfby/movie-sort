package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

const linksAPI = "/api/links"

type linksHandlers struct {
	Links *links
}

func newLinksHandlers(l *links) *linksHandlers {
	return &linksHandlers{l}
}

func (l *linksHandlers) getLinksHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if len(path) < len(linksAPI) {
			w.WriteHeader(404)
			return
		}
		path = path[len(linksAPI):]
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		links := l.Links.getLinks(ioutil.ReadDir, os.SameFile, "", nil)
		json.NewEncoder(w).Encode(links)
	})
}
