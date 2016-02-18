package server

import (
	"github.com/gorilla/mux"
	"github.com/werwolfby/movie-sort/guessit"
	"github.com/werwolfby/movie-sort/links"
	"github.com/werwolfby/movie-sort/utils"
	"net/http"
)

type GuessitHandlers struct {
	guessItService guessit.GuessItService
	links          links.Links
}

func NewGuessitHandlers(gs guessit.GuessItService, l links.Links) *GuessitHandlers {
	return &GuessitHandlers{guessItService: gs, links: l}
}

func (g GuessitHandlers) GetHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		path := vars["path"]
		splittedPath := utils.SplitPath(path)
		if len(splittedPath) < 2 {
			writeNotFound(w, "folder is not specified")
		}
		filename := splittedPath[len(splittedPath)-1]
		fi, err := g.guessItService.GuessLink(filename)
		if err != nil {
			writeInternalServerError(w, err)
			return
		}
		writeOk(w, fi)
	})
}
