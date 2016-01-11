package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

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

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	http.Handle("/", r)
	log.Println("Starting server on http://localhost:88")
	log.Fatal(http.ListenAndServe(":88", nil))
}
