package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type settings struct {
	PathSeparator string `json:"pathSeparator"`
}

type folderInfo struct {
	Name string   `json:"name"`
	Path []string `json:"path"`
}

func globalSettings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	s := settings{PathSeparator: fmt.Sprintf("%c", os.PathSeparator)}
	json.NewEncoder(w).Encode(s)
}

func inputFolders(w http.ResponseWriter, r *http.Request) {
	folders := []folderInfo{{Name: "Downloads", Path: []string{"D:", "Video", "Complete"}}}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(folders)
}

func outputFolders(w http.ResponseWriter, r *http.Request) {
	folders := []folderInfo{{Name: "Movies", Path: []string{"D:", "Video", "Films"}}, {Name: "Shows", Path: []string{"D:", "Video", "Serials"}}}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(folders)
}
