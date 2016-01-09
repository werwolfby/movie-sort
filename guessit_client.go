package main

import (
	"github.com/jmcvetta/napping"
	"net/url"
)

type filenameInfo struct {
	Type   string `json:"type"`
	Title  string `json:"title"`
	Season int    `json:"season"`
}

type guessItService struct {
	cfg *config
}

func newGuessItService(c *config) *guessItService {
	return &guessItService{cfg: c}
}

func (g *guessItService) guessIt(filename string) (*filenameInfo, error) {
	params := url.Values{}
	params.Add("filename", filename)
	var info filenameInfo
	_, err := napping.Get(g.cfg.Services.GuessItURL, &params, &info, nil)
	if err != nil {
		return nil, err
	}
	return &info, nil
}
