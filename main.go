package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.Headers("Content-Type", "application/json")
	r.HandleFunc("/api/settings", globalSettings).Methods("GET")
	r.HandleFunc("/api/settings/input-folders", inputFolders).Methods("GET")
	r.HandleFunc("/api/settings/output-folders", outputFolders).Methods("GET")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":88", nil))
}
