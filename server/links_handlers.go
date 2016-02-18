package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/werwolfby/movie-sort/links"
	"github.com/werwolfby/movie-sort/utils"
	"net/http"
)

type LinksHandlers struct {
	links links.Links
}

func NewLinksHandlers(l links.Links) *LinksHandlers {
	return &LinksHandlers{links: l}
}

func (l LinksHandlers) GetHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l.links.UpdateLinks([]string{"mkv", "avi", "mp4"})
		writeOk(w, l.links.GetLinks())
	})
}

func (l LinksHandlers) GetPutHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		path := vars["path"]
		splittedPath := utils.SplitPath(path)
		if len(splittedPath) < 2 {
			writeNotFound(w, "folder is not specified")
			return
		}
		inputFolder := splittedPath[0]
		inputFileInfo := links.FileInfo{Folder: inputFolder, Path: splittedPath[1 : len(splittedPath)-1], Name: splittedPath[len(splittedPath)-1]}

		var linkFileInfo links.FileInfo
		if json.NewDecoder(r.Body).Decode(&linkFileInfo) != nil {
			writeBadRequest(w, "Can't parse body")
			return
		}

		l, e := l.links.Link(inputFileInfo, linkFileInfo)

		if e != nil {
			writeInternalServerError(w, e)
		}

		writeOk(w, l)
	})
}
