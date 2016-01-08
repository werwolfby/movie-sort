package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

const linksAPI = "/api/links"

type linksHandlers struct {
	Settings *settings
}

func newLinksHandlers(s *settings) *linksHandlers {
	return &linksHandlers{s}
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
		links := l.getLinks("", nil)
		json.NewEncoder(w).Encode(links)
	})
}

func (l *linksHandlers) getLinks(folder string, path []string) []linkInfo {
	return searchHardLinks(ioutil.ReadDir, os.SameFile,
		l.Settings.InputFoldersSettings.folders,
		l.Settings.OutputFoldersSettings.folders)
}
