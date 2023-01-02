package main

import (
	"fmt"
	"os"
)

type PgNum uint64

type Page struct {
	num 		PgNum
	data 		[]byte
}

type DAL struct {
	file 		*os.File
	pageSize 	int
	*FreeList	
}

func newDAL(path string, pageSize int) (*DAL, error) {
	file, err := os.OpenFile(path, os.O_RDWR | os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	dAL := &DAL{
		file,
		pageSize,
		newFreeList(),
	}
	return dAL, nil
}

func (dAL *DAL) close() error {
	if dAL.file != nil {
		err := dAL.file.Close()
		if err != nil {
			return fmt.Errorf("could not close file: %s", err)
		}
		dAL.file = nil
	}
	return nil
}

func (dAl *DAL) allocateEmptyPage() *Page {
	return &Page{
		data: make([]byte, dAl.pageSize),
	}
}

func (dAL *DAL) readPage(pageNum PgNum) (*Page, error) {
	p := dAL.allocateEmptyPage()

	offset := int(pageNum) * dAL.pageSize

	_, err := dAL.file.ReadAt(p.data, int64(offset))
	if err != nil {
		return nil, err
	}
	return p, err
}

func (dAl *DAL) writePage(p *Page) error {
	offset := int(p.num) * dAl.pageSize
	_, err := dAl.file.WriteAt(p.data, int64(offset))
	return err
}