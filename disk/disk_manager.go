package disk

import (
	"io"
	"os"
)

const (
	PageSize = 4096
)

type Page struct {
	Data [PageSize]byte
}

type Manager struct {
	file *os.File
}

func NewDiscManager(file *os.File) (*Manager, error) {
	return &Manager{file: file}, nil
}

func (dm *Manager) ReadPage(pageID int64) (*Page, error) {
	offset := pageID * PageSize
	_, err := dm.file.Seek(offset, io.SeekStart)
	if err != nil {
		return nil, err
	}

	page := &Page{}
	_, err = dm.file.Read(page.Data[:])
	if err != nil {
		return nil, err
	}
	return page, nil
}

func (dm *Manager) WritePage(pageID int64, page *Page) error {
	offset := pageID * PageSize
	_, err := dm.file.Seek(offset, io.SeekStart)
	if err != nil {
		return err
	}

	_, err = dm.file.Write(page.Data[:])
	return err
}
