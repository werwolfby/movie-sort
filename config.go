package main

import "gopkg.in/gcfg.v1"

type config struct {
	Paths struct {
		Source     string `gcfg:"src"`
		DestMovies string `gcfg:"dest-movies"`
		DestShows  string `gcfg:"dest-shows"`
	}
}

func readConfig(filename string) *config {
	cfg := new(config)
	gcfg.ReadFileInto(cfg, "config.ini")
	return cfg
}
