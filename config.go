package main

import "gopkg.in/gcfg.v1"

type configPaths struct {
	Source     string `gcfg:"src"`
	DestMovies string `gcfg:"dest-movies"`
	DestShows  string `gcfg:"dest-shows"`
}

type configServices struct {
	GuessItURL string `gcfg:"guessit"`
}

type config struct {
	Paths    configPaths
	Services configServices
}

func readConfig(filename string) (*config, error) {
	cfg := new(config)
	if e := gcfg.ReadFileInto(cfg, "config.ini"); e != nil {
		return nil, e
	}
	return cfg, nil
}
