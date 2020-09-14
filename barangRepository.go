package main

import (
	"crypto/md5"
	"fmt"
)

// var barangCollection = make([]Barang, 0)

type BarangReposistory struct {
	barangCollection      *[]*Barang
	barangCollectionInfra *BarangRepositoryInfrastructure
}

func NewBarangRepository(dataPath string, barangCollection []*Barang) *BarangReposistory {
	barangRepoInfra := NewBarangRepoInfra(dataPath)
	return &BarangReposistory{barangCollection: &barangCollection, barangCollectionInfra: barangRepoInfra}
}

func (br *BarangReposistory) AddNewBarang(barang *Barang) {
	data := []byte(barang.Nama)
	barang.ID = fmt.Sprintf("%x", md5.Sum(data))
	*br.barangCollection = append(*br.barangCollection, barang)
	br.barangCollectionInfra.saveToFile(br.barangCollection)
}

func (br *BarangReposistory) FindAllBarang() []*Barang {
	br.barangCollectionInfra.readFile(br.barangCollection)
	return *br.barangCollection
}
func (br BarangReposistory) FindBarangField(fnFilter func(Barang) bool) []*Barang {
	br.barangCollectionInfra.readFile(br.barangCollection)
	var searchResult = make([]*Barang, 0)
	for _, b := range *br.barangCollection {
		if fnFilter(*b) {
			searchResult = append(searchResult, b)
		}
	}
	return searchResult
}
