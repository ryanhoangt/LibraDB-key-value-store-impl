package main

import (
	"os"
)

func main() {
	// fmt.Println("Hello world!")
	dAL, _ := newDAL("db.db", os.Getpagesize())

	p := dAL.allocateEmptyPage()
	p.num = dAL.getNextPage()
	copy(p.data, "somedummydataabcdef")

	_ = dAL.writePage(p)
}