package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

const (
	pathSep = "\\/"
)

type settings struct {
	PathSeparator string `json:"pathSeparator"`
}

type folderInfo struct {
	Name string   `json:"name"`
	Path []string `json:"path"`
}

type globalSettingsHandler struct {
}

type inputFoldersHandler struct {
	cfg *config
}

type outputFoldersHandler struct {
	cfg *config
}

func splitPath(s string) []string {
	f := func(r rune) bool {
		for _, c := range pathSep {
			if r == c {
				return true
			}
		}
		return false
	}

	return strings.FieldsFunc(s, f)
}

func (globalSettingsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	s := settings{PathSeparator: fmt.Sprintf("%c", os.PathSeparator)}
	json.NewEncoder(w).Encode(s)
}

func (h inputFoldersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := splitPath(h.cfg.Paths.Source)

	folders := []folderInfo{{Name: "Downloads", Path: path}}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(folders)
}

func (h outputFoldersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	moviesPath := splitPath(h.cfg.Paths.DestMovies)
	showsPath := splitPath(h.cfg.Paths.DestShows)

	folders := []folderInfo{{Name: "Movies", Path: moviesPath}, {Name: "Shows", Path: showsPath}}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(folders)
}
