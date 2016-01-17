package links

import (
	"github.com/werwolfby/movie-sort/settings"
	"os"
	"path/filepath"
	"strings"
)

type FileInfo struct {
	Folder string   `json:"folder"`
	Path   []string `json:"path"`
	Name   string   `json:"name"`
}

type LinkInfo struct {
	FileInfo
	Links []FileInfo `json:"links"`
}

type Links interface {
	UpdateLinks(extensions []string) error
	GetShows() []string
	GetLinks() []LinkInfo
}

type links struct {
	Reader        LinksReader
	InputFolders  *settings.InputFoldersSettings
	OutputFolders *settings.OutputFoldersSettings
	Links         []LinkInfo
	Shows         []string
}

type searchFileInfo struct {
	FileInfo
	OsFileInfo os.FileInfo
}

func NewLinks(r LinksReader, ifs *settings.InputFoldersSettings, ofs *settings.OutputFoldersSettings) Links {
	return &links{Reader: r, InputFolders: ifs, OutputFolders: ofs}
}

func (l *links) UpdateLinks(extensions []string) error {
	inputFiles, err := l.getAllFiles(l.InputFolders.Folders, extensions)
	if err != nil {
		return err
	}
	outputFiles, err := l.getAllFiles(l.OutputFolders.Folders, extensions)
	if err != nil {
		return err
	}
	showsFolders := l.OutputFolders.GetShows()
	l.Shows = l.searchShows(showsFolders, outputFiles)
	l.Links = l.searchLinks(inputFiles, outputFiles)
	return nil
}

func (l links) GetLinks() []LinkInfo {
	return l.Links
}

func (l links) GetShows() []string {
	return l.Shows
}

func (l links) searchShows(folders []settings.FolderInfo, files []searchFileInfo) []string {
	var shows []string
	for _, file := range files {
		showsFolder := false
		for _, folder := range folders {
			if file.Folder == folder.Name {
				showsFolder = true
				break
			}
		}
		if !showsFolder {
			continue
		}
		var fileShow string
		if len(file.Path) > 0 && len(file.Path[0]) > 0 {
			fileShow = file.Path[0]
		}
		if fileShow != "" {
			newShow := true
			for _, show := range shows {
				if strings.EqualFold(show, fileShow) {
					newShow = false
					break
				}
			}
			if newShow {
				shows = append(shows, fileShow)
			}
		}
	}
	return shows
}

func (l links) getAllFiles(folders []settings.FolderInfo, extensions []string) ([]searchFileInfo, error) {
	var result []searchFileInfo
	for _, f := range folders {
		folderFiles, err := l.getAllFilesFromFolder(f, nil, extensions)
		if err != nil {
			return nil, err
		}
		result = append(result, folderFiles...)
	}
	return result, nil
}

func (l links) getAllFilesFromFolder(folder settings.FolderInfo, dirPath []string, extensions []string) ([]searchFileInfo, error) {
	dirname := filepath.Join(append(folder.Path, dirPath...)...)
	files, err := l.Reader.ReadDir(dirname)
	if err != nil {
		return nil, err
	}
	result := make([]searchFileInfo, 0, len(files))

	for _, fi := range files {
		if !fi.IsDir() {
			ext := filepath.Ext(fi.Name())
			if len(ext) == 0 {
				continue
			}
			ext = ext[1:]
			for _, e := range extensions {
				if ext == e {
					sfi := searchFileInfo{FileInfo: FileInfo{Folder: folder.Name, Path: dirPath, Name: fi.Name()}, OsFileInfo: fi}
					result = append(result, sfi)
					break
				}
			}
		} else {
			childDirPath := append(dirPath, fi.Name())
			childFiles, err := l.getAllFilesFromFolder(folder, childDirPath, extensions)
			if err == nil {
				result = append(result, childFiles...)
			}
		}
	}

	return result, nil
}
