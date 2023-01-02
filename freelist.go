package main

// metaPage is the maximum pgnum of all pages used by the db for its own purposes. For now, only page 0 is used as the
// header page. It means all other page numbers can be used, starting from page 1.
const maxMetaPgNum = 0

type FreeList struct {
	curMaxPage		PgNum
	releasedPages 	[]PgNum
}

func newFreeList() *FreeList {
	return &FreeList{
		curMaxPage: maxMetaPgNum,
		releasedPages: []PgNum{},
	}
}

func (fr *FreeList) getNextPage() PgNum {
	if len(fr.releasedPages) != 0 {
		pageId := fr.releasedPages[len(fr.releasedPages)-1]
		// remove the last element from the list
		fr.releasedPages = fr.releasedPages[:len(fr.releasedPages)-1]
		return pageId
	}

	fr.curMaxPage += 1
	return fr.curMaxPage
}

func (fr *FreeList) releasePage(page PgNum) {
	fr.releasedPages = append(fr.releasedPages, page)
}