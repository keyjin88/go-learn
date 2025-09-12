package bins

import "time"

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type BinList struct {
	bins []Bin
}

func (b *Bin) NewBin(id string, private bool, createdAt time.Time, name string) {
	b.id = id
	b.private = private
	b.createdAt = createdAt
	b.name = name
}

func (b *BinList) NewBinList(bins []Bin) {
	b.bins = bins
}
