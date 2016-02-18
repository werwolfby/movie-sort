package links

import (
	"errors"
	"github.com/werwolfby/movie-sort/settings"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"
)

const (
	seasonPrefix   = "Season "
	SeasonNotFound = -1
	ShowNotFound   = -2
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

type Folders interface {
	GetShowsFolder() (settings.FolderInfo, bool)
	GetMoviesFolder() (settings.FolderInfo, bool)
	GetShow(name string) (FileInfo, bool)
	GetShows() []FileInfo
	GetShowSeason(name string, season int) (FileInfo, int)
}

type Links interface {
	Folders
	UpdateLinks(extensions []string) error
	GetLinks() []LinkInfo
	Link(oldname, newname FileInfo) (LinkInfo, error)
}

type links struct {
	Reader        LinksReader
	InputFolders  *settings.InputFoldersSettings
	OutputFolders *settings.OutputFoldersSettings
	Links         []LinkInfo
	Shows         []FileInfo
	ShowsSeasons  map[string]map[int]FileInfo
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
	l.Shows, l.ShowsSeasons = l.searchShows(showsFolders, outputFiles)
	l.Links = l.searchLinks(inputFiles, outputFiles)
	return nil
}

func (l links) GetLinks() []LinkInfo {
	return l.Links
}

func (l links) GetShows() []FileInfo {
	return l.Shows
}

func (l links) GetShowsFolder() (settings.FolderInfo, bool) {
	shows := l.OutputFolders.GetShows()
	if len(shows) == 0 {
		return settings.FolderInfo{}, false
	}
	return shows[0], true
}

func (l links) GetMoviesFolder() (settings.FolderInfo, bool) {
	movies := l.OutputFolders.GetMovies()
	if len(movies) == 0 {
		return settings.FolderInfo{}, false
	}
	return movies[0], true
}

func (l links) GetShow(name string) (FileInfo, bool) {
	for _, fi := range l.Shows {
		if strings.EqualFold(fi.Name, name) {
			return fi, true
		}
	}
	return FileInfo{}, false
}

func (l links) GetShowSeason(name string, season int) (FileInfo, int) {
	for _, fi := range l.Shows {
		if strings.EqualFold(fi.Name, name) {
			seasonFileInfo, found := l.ShowsSeasons[fi.Name][season]
			if !found {
				return fi, SeasonNotFound
			}
			return seasonFileInfo, season
		}
	}
	return FileInfo{}, ShowNotFound
}

func (l links) Link(oldname, newname FileInfo) (LinkInfo, error) {
	inputFolder, found := l.InputFolders.Find(oldname.Folder)
	if !found {
		return LinkInfo{}, errors.New("Can't find input folder: " + oldname.Folder)
	}
	outputFolder, found := l.OutputFolders.Find(newname.Folder)
	if !found {
		return LinkInfo{}, errors.New("Can't find output folder: " + newname.Folder)
	}

	oldlink := l.findLink(oldname)
	if oldlink == nil {
		return LinkInfo{}, errors.New("Can't find source link")
	}

	oldpath := getFullPath(inputFolder, oldname)
	newpath := getFullPath(outputFolder, newname)

	newdir := filepath.Dir(newpath)
	if exists, _ := isDirExists(newdir); !exists {
		if e := l.Reader.MkdirAll(newdir); e != nil {
			return LinkInfo{}, e
		}
	}

	if e := l.Reader.Link(oldpath, newpath); e != nil {
		return LinkInfo{}, e
	}

	oldlink.Links = append(oldlink.Links, newname)

	return *oldlink, nil
}

func (l *links) findLink(file FileInfo) *LinkInfo {
	for i, link := range l.Links {
		if link.FileInfo.Folder == file.Folder && pathEqual(link.FileInfo.Path, file.Path) && link.FileInfo.Name == file.Name {
			return &l.Links[i]
		}
	}
	return nil
}

func pathEqual(path1, path2 []string) bool {
	if len(path1) != len(path2) {
		return false
	}
	for i, p := range path1 {
		if p != path2[i] {
			return false
		}
	}
	return true
}

func getFullPath(folder settings.FolderInfo, file FileInfo) string {
	return filepath.Join(append(append(folder.Path, file.Path...), file.Name)...)
}

func isDirExists(path string) (bool, error) {
	stat, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return stat != nil, err
}

func (l links) searchShows(folders []settings.FolderInfo, files []searchFileInfo) (shows []FileInfo, showsSeasons map[string]map[int]FileInfo) {
	for _, file := range files {
		var showsFolder string
		for _, folder := range folders {
			if file.Folder == folder.Name {
				showsFolder = folder.Name
				break
			}
		}
		if showsFolder == "" || len(file.Path) == 0 || len(file.Path[0]) == 0 {
			continue
		}
		fileShowPath := file.Path[1:]
		fileShow := file.Path[0]

		var foundShowName string
		for _, show := range shows {
			if strings.EqualFold(show.Name, fileShow) {
				foundShowName = show.Name
				break
			}
		}
		if foundShowName == "" {
			shows = append(shows, FileInfo{Folder: showsFolder, Path: []string{}, Name: fileShow})
			foundShowName = fileShow
		}
		if len(fileShowPath) < 1 {
			continue
		}
		seasonDir := fileShowPath[0]
		// season dir have to be path like "Season %d+"
		if len(seasonDir) < len(seasonPrefix)+1 || !strings.EqualFold(seasonPrefix, seasonDir[0:len(seasonPrefix)]) {
			continue
		}
		i := len(seasonPrefix)
		for ; i < len(seasonDir); i++ {
			if !unicode.IsSpace([]rune(seasonDir)[i]) {
				break
			}
		}
		seasonNumber, err := strconv.Atoi(seasonDir[i:])
		if err != nil || seasonNumber < 0 {
			continue
		}
		if showsSeasons == nil {
			showsSeasons = make(map[string]map[int]FileInfo)
		}
		showSeasons, found := showsSeasons[foundShowName]
		if !found {
			showSeasons = make(map[int]FileInfo)
			showsSeasons[foundShowName] = showSeasons
		}
		showSeasons[seasonNumber] = FileInfo{Folder: showsFolder, Path: []string{seasonDir}, Name: foundShowName}
	}
	return
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
