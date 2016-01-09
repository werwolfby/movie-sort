package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type guessitHandler struct {
	Settings       *settings
	GuessItService *guessItService
}

func newGuessitHandler(s *settings, g *guessItService) *guessitHandler {
	return &guessitHandler{s, g}
}

func (g *guessitHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	path := vars["path"]
	splittedPath := splitPath(path)
	filename := splittedPath[len(splittedPath)-1]
	guessInfo, e := g.GuessItService.guessIt(filename)
	if e != nil {
		w.WriteHeader(500)
		return
	}
	var folder string
	var guessPath []string
	if guessInfo.Type == "episode" {
		folder = "Shows"
		guessPath = []string{guessInfo.Title, fmt.Sprintf("Season %d", guessInfo.Season)}
	} else {
		folder = "Movies"
		guessPath = []string{}
	}
	result := fileInfo{
		Folder: folder,
		Path:   guessPath,
		Name:   filename,
	}
	json.NewEncoder(w).Encode(result)
}
