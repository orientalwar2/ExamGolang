package main

type Barang struct {
	ID        string
	Nama      string
	Kategori  string
	Harga     float64
	Stok      float64
	Deskripsi string
	Ukuran    float64
}

func NewBarang(nama string, kategori string, harga float64, stok float64, deskripsi string, ukuran float64) Barang {
	return Barang{
		Nama:      nama,
		Kategori:  kategori,
		Harga:     harga,
		Stok:      stok,
		Deskripsi: deskripsi,
		Ukuran:    ukuran,
	}
}
