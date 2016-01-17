package links

import (
	"github.com/werwolfby/movie-sort/settings"
	"io/ioutil"
	"os"
	"path/filepath"
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

type links struct {
	Reader        LinksReader
	InputFolders  *settings.InputFoldersSettings
	OutputFolders *settings.OutputFoldersSettings
	Links         []LinkInfo
}

type Links interface {
	UpdateLinks(extensions []string) error
	GetLinks() []LinkInfo
}

type DirReader interface {
	ReadDir(dirname string) ([]os.FileInfo, error)
}

type SameFileComparer interface {
	SameFile(fi1, fi2 os.FileInfo) bool
}

type DirMaker interface {
	MkdirAll(dirname string) error
}

type Linker interface {
	Link(oldname, newname string) error
}

type LinksReader interface {
	DirReader
	SameFileComparer
	DirMaker
	Linker
}

type OsLinksReader struct {
}

func (OsLinksReader) ReadDir(dirname string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(dirname)
}

func (OsLinksReader) SameFile(fi1, fi2 os.FileInfo) bool {
	return os.SameFile(fi1, fi2)
}

func (OsLinksReader) MkdirAll(dirname string) error {
	return os.MkdirAll(dirname, 0755)
}

func (OsLinksReader) Link(oldname, newname string) error {
	return os.Link(oldname, newname)
}

func NewLinksReader() LinksReader {
	return OsLinksReader{}
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
	l.Links = l.searchLinks(inputFiles, outputFiles)
	return nil
}

func (l *links) GetLinks() []LinkInfo {
	return l.Links
}

type searchFileInfo struct {
	FileInfo
	OsFileInfo os.FileInfo
}

func (l *links) getAllFiles(folders []settings.FolderInfo, extensions []string) ([]searchFileInfo, error) {
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

func (l *links) getAllFilesFromFolder(folder settings.FolderInfo, dirPath []string, extensions []string) ([]searchFileInfo, error) {
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
