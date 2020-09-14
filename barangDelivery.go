package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type BarangDelivery struct {
	barangService *BarangService
}

func NewBarangDelivery(c *AppConfig) *BarangDelivery {
	repo := NewBarangRepository(c.getDbPath(), c.getDb())
	barangService := NewBarangService(repo)
	return &BarangDelivery{barangService}
}

func (bd *BarangDelivery) create() {
	var isExist = false
	var userChoice string

	bd.mainMenuForm()
	for isExist == false {
		fmt.Printf("\n%s", "Pilih nomor : ")
		fmt.Scan(&userChoice)
		switch userChoice {
		case "01":
			bd.registerationBarangForm()
		case "02":
			bd.viewBarangCollectionForm()
		case "q":
			isExist = true
		default:
			fmt.Println("Menu tidak diketahui")
		}
	}
}

func (bd *BarangDelivery) menuChoiceOrdered(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (bd *BarangDelivery) mainMenuForm() {
	var appMenu = map[string]string{
		"01": "tambah barang ke daftar barang",
		"02": "lihat daftar barang",
		"q":  "keluar dari aplikasi",
	}
	fmt.Printf("%s\n", strings.Repeat("*", 30))
	fmt.Printf("%26s\n", "Daftar Barang")
	fmt.Printf("%s\n", strings.Repeat("*", 30))
	for _, menuCode := range bd.menuChoiceOrdered(appMenu) {
		fmt.Printf("%s. %s\n", menuCode, appMenu[menuCode])
	}
}

func (bd *BarangDelivery) registerationBarangForm() {
	consoleClear()
	var nama string
	var kategori string
	var harga float64
	var stok float64
	var deskripsi string
	var ukuran float64
	var confirmation string
	fmt.Println()
	fmt.Printf("%s\n", "Form Registrasi Barang")
	fmt.Printf("%s\n", strings.Repeat("-", 30))
	scanner := bufio.NewReader(os.Stdin)
	sAsd, _ := scanner.ReadString('\n')
	nama = strings.TrimSpace(sAsd)
	fmt.Print("Nama : ")
	sNama, _ := scanner.ReadString('\n')
	nama = strings.TrimSpace(sNama)
	fmt.Print("Kategori : ")
	sKategori, _ := scanner.ReadString('\n')
	kategori = strings.TrimSpace(sKategori)
	fmt.Print("Harga : ")
	sHarga, _ := scanner.ReadString('\n')
	harga, _ = strconv.ParseFloat(strings.TrimSpace(sHarga), 64)
	fmt.Print("Stok : ")
	sStok, _ := scanner.ReadString('\n')
	stok, _ = strconv.ParseFloat(strings.TrimSpace(sStok), 64)
	fmt.Print("Deskripsi : ")
	sDeskripsi, _ := scanner.ReadString('\n')
	deskripsi = strings.TrimSpace(sDeskripsi)
	fmt.Print("Ukuran : ")
	sUkuran, _ := scanner.ReadString('\n')
	ukuran, _ = strconv.ParseFloat(strings.TrimSpace(sUkuran), 64)

	fmt.Println("Simpan ke daftar barang ? : (y/n)")
	fmt.Scanln(&confirmation)

	if confirmation == "y" {
		newBarang := NewBarang(nama, kategori, harga, stok, deskripsi, ukuran)
		bd.barangService.registerNewBarang(&newBarang)
	}
	consoleClear()
	bd.mainMenuForm()
}

func (bd *BarangDelivery) viewBarangCollectionForm() {
	consoleClear()
	barangs := bd.barangService.getAllBarangNewBarang()
	fmt.Println("")
	fmt.Printf("Daftar Barang\n")
	fmt.Printf("%s\n", strings.Repeat("=", 150))
	fmt.Printf("%-40s%-30s%-15s%-15s%-15s%-15s%15s\n", "ID", "Nama", "Kategori", "Harga", "Stok", "Deskripsi", "Ukuran")
	fmt.Printf("%s\n", strings.Repeat("=", 150))
	for _, b := range barangs {
		fmt.Printf("%-40s%-30s%-15s%-15.2f%-15.2f%-15s%15.2f\n", b.ID, b.Nama, b.Kategori, b.Harga, b.Stok, b.Deskripsi, b.Ukuran)
	}

	var confirmation2 string
	fmt.Print("Kembali ke menu utama ?")
	fmt.Scanln(&confirmation2)

	if confirmation2 == "y" {
		bd.mainMenuForm()
	}

}
