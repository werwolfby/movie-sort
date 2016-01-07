package main

import "gopkg.in/gcfg.v1"

type configPaths struct {
	Source     string `gcfg:"src"`
	DestMovies string `gcfg:"dest-movies"`
	DestShows  string `gcfg:"dest-shows"`
}

type config struct {
	Paths configPaths
}

func readConfig(filename string) *config {
	cfg := new(config)
	gcfg.ReadFileInto(cfg, "config.ini")
	return cfg
}
