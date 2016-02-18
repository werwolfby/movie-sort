package main

import (
	"errors"
	"github.com/gorilla/mux"
	"github.com/werwolfby/movie-sort/guessit"
	"github.com/werwolfby/movie-sort/links"
	"github.com/werwolfby/movie-sort/server"
	"github.com/werwolfby/movie-sort/settings"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type indexIfNotExist string

func (d indexIfNotExist) Open(name string) (http.File, error) {
	if filepath.Separator != '/' && strings.IndexRune(name, filepath.Separator) >= 0 ||
		strings.Contains(name, "\x00") {
		return nil, errors.New("http: invalid character in file path")
	}
	dir := string(d)
	if dir == "" {
		dir = "."
	}
	f, err := os.Open(filepath.Join(dir, filepath.FromSlash(path.Clean("/"+name))))
	if os.IsNotExist(err) {
		return d.Open("index.html")
	}
	if err != nil {
		return nil, err
	}
	return f, nil
}

func main() {
	s, e := settings.ReadSettingsFromFile("config.ini")
	if e != nil {
		log.Fatal(e)
	}

	lr := links.NewLinksReader()
	l := links.NewLinks(lr, &s.InputFolders, &s.OutputFolders)
	g := guessit.NewGuessItService(s.Services.GuessItURL, l)

	sh := server.NewSettingsHandlers(s)
	lh := server.NewLinksHandlers(l)
	gh := server.NewGuessitHandlers(g, l)

	r := mux.NewRouter()

	r.Handle("/api/settings", sh.GetGlobalSettingsHandler()).Methods("GET")
	r.Handle("/api/settings/input-folders", sh.GetInputFoldersSettingsHandler()).Methods("GET")
	r.Handle("/api/settings/output-folders", sh.GetOutputFoldersSettingsHandler()).Methods("GET")

	r.Handle("/api/links", lh.GetHandler()).Methods("GET")
	//r.Handle("/api/links/{path:.*}", lh.getPutLinksHandler(os.Link)).Methods("PUT")

	r.Handle("/api/guess/{path:.*}", gh.GetHandler())

	r.PathPrefix("/").Handler(http.FileServer(indexIfNotExist("static")))

	http.Handle("/", r)

	l.UpdateLinks([]string{"mkv", "avi", "mp4"})

	log.Println("Starting server on http://localhost:88")
	log.Fatal(http.ListenAndServe(":88", nil))
}
