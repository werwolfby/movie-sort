package main

type linkInfo struct {
	fileInfo
	Links []fileInfo `json:"links"`
}
