package guessit

import (
	"github.com/jmcvetta/napping"
	"net/url"
)

type FilenameInfo struct {
	Type   string `json:"type"`
	Title  string `json:"title"`
	Season int    `json:"season"`
}

type guessItService struct {
	Session    *napping.Session
	ServiceURL string
}

type GuessItService interface {
	GuessIt(filename string) (*FilenameInfo, error)
}

func NewGuessItService(serviceUrl string) GuessItService {
	return &guessItService{ServiceURL: serviceUrl}
}

func (g *guessItService) GuessIt(filename string) (*FilenameInfo, error) {
	if g.Session == nil {
		g.Session = &napping.Session{}
	}

	params := url.Values{}
	params.Add("filename", filename)
	var info FilenameInfo
	_, err := g.Session.Get(g.ServiceURL, &params, &info, nil)
	if err != nil {
		return nil, err
	}
	return &info, nil
}
