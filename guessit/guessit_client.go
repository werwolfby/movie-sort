package guessit

import (
	"errors"
	"fmt"
	"github.com/jmcvetta/napping"
	"github.com/werwolfby/movie-sort/links"
	"net/http"
	"net/url"
)

type filenameInfo struct {
	Type   string
	Title  string
	Season int
}

type guessItService struct {
	session    *napping.Session
	folders    links.Folders
	serviceURL string
}

type GuessItService interface {
	GuessLink(filename string) (*links.FileInfo, error)
}

func NewGuessItService(serviceUrl string, folders links.Folders) GuessItService {
	return &guessItService{serviceURL: serviceUrl, folders: folders}
}

func (g guessItService) guessIt(filename string) (*filenameInfo, error) {
	if g.session == nil {
		g.session = &napping.Session{}
	}

	params := url.Values{}
	params.Add("filename", filename)
	var info filenameInfo
	resp, err := g.session.Get(g.serviceURL, &params, &info, nil)
	if err != nil {
		return nil, err
	}
	if resp.Status() != http.StatusOK {
		return nil, fmt.Errorf("Unexpected status: %d", resp.Status())
	}
	return &info, nil
}

func (g guessItService) GuessLink(filename string) (*links.FileInfo, error) {
	res, err := g.guessIt(filename)
	if err != nil {
		return nil, err
	}
	var fileInfo *links.FileInfo
	if res.Type == "episode" {
		show, season := g.folders.GetShowSeason(res.Title, res.Season)
		if season == links.ShowNotFound {
			showsFolder, set := g.folders.GetShowsFolder()
			if !set {
				return nil, errors.New("Shows folders doens't set")
			}
			title := res.Title
			seasonPath := fmt.Sprintf("Season %d", res.Season)
			fileInfo = &links.FileInfo{Folder: showsFolder.Name, Path: []string{title, seasonPath}, Name: filename}
		} else if season == links.SeasonNotFound {
			seasonPath := fmt.Sprintf("Season %d", res.Season)
			fileInfo = &links.FileInfo{Folder: show.Folder, Path: append(show.Path, show.Name, seasonPath), Name: filename}
		} else {
			fileInfo = &links.FileInfo{Folder: show.Folder, Path: append(show.Path, show.Name), Name: filename}
		}
	} else {
		moviesFolder, set := g.folders.GetMoviesFolder()
		if !set {
			return nil, errors.New("Movies folders doens't set")
		}
		fileInfo = &links.FileInfo{Folder: moviesFolder.Name, Path: []string{}, Name: filename}
	}
	return fileInfo, nil
}
