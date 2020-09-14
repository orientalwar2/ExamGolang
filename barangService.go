package main

import "strings"

type BarangService struct {
	r *BarangReposistory
}

func NewBarangService(repo *BarangReposistory) *BarangService {
	return &BarangService{r: repo}
}

func (bs *BarangService) registerNewBarang(b *Barang) {
	bs.r.AddNewBarang(b)
}

func (bs *BarangService) getAllBarangNewBarang() []*Barang {
	return bs.r.FindAllBarang()
}

func (bs *BarangService) searchBarangID(id string) []*Barang {
	var funcFilter = func(b Barang) bool {
		return b.ID == id
	}
	return bs.r.FindBarangField(funcFilter)

}

func (bs *BarangService) searchBarangByTitle(nama string) []*Barang {
	var funcFilter = func(b Barang) bool {
		return strings.Contains(b.Nama, nama)
	}
	return bs.r.FindBarangField(funcFilter)

}
