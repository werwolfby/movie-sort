package main

type linkInfo struct {
	fileInfo
	Links []fileInfo `json:"links"`
}

type links struct {
	Settings *settings
}

func newLinks(s *settings) *links {
	return &links{s}
}

func (l *links) getLinks(reader readDirFunc, sameFile sameFileFunc, folder string, path []string) []linkInfo {
	return searchHardLinks(reader, sameFile,
		l.Settings.InputFoldersSettings.folders,
		l.Settings.OutputFoldersSettings.folders)
}
