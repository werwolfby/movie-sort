package main

import (
	"errors"
	"github.com/gorilla/mux"
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
	c, e := readConfig("config.ini")
	if e != nil {
		log.Fatal(e)
	}

	s := newSettings(c)
	g := newGuessItService(c)
	l := newLinks(s)
	sh := newSettingsHandler(s)
	lh := newLinksHandlers(l)
	gh := newGuessitHandler(s, g)

	r := mux.NewRouter()

	r.Handle("/api/settings", sh.getGlobalSettingsHandler()).Methods("GET")
	r.Handle("/api/settings/input-folders", sh.getInputFoldersSettingsHandler()).Methods("GET")
	r.Handle("/api/settings/output-folders", sh.getOutputFoldersSettingsHandler()).Methods("GET")

	r.Handle("/api/links", lh.getLinksHandler()).Methods("GET")
	r.Handle("/api/links/{path:.*}", lh.getPutLinksHandler(os.Link)).Methods("PUT")

	r.Handle("/api/guess/{folder}/{path:.*}", gh.getHandler())

	r.PathPrefix("/").Handler(http.FileServer(indexIfNotExist("static")))

	http.Handle("/", r)
	log.Println("Starting server on http://localhost:88")
	log.Fatal(http.ListenAndServe(":88", nil))
}
