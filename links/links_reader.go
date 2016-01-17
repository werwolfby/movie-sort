package links

import (
	"io/ioutil"
	"os"
)

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

type osLinksReader struct {
}

func (osLinksReader) ReadDir(dirname string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(dirname)
}

func (osLinksReader) SameFile(fi1, fi2 os.FileInfo) bool {
	return os.SameFile(fi1, fi2)
}

func (osLinksReader) MkdirAll(dirname string) error {
	return os.MkdirAll(dirname, 0755)
}

func (osLinksReader) Link(oldname, newname string) error {
	return os.Link(oldname, newname)
}

func NewLinksReader() LinksReader {
	return osLinksReader{}
}
