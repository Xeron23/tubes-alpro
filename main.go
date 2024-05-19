package main

import (
	"fmt"
	"test-tubes/program"
)

var aksi, banPasien int
var dataPasien program.Arr
var namPasien string
var adaPasien bool
var henti bool

func main() {
	henti = true
	banPasien = 0
	for henti {
		program.Interface()
		fmt.Scan(&aksi)
		if aksi == 1 {
			fmt.Println("\nSILAHKAN MASUKAN DATA PASIEN")
			program.InputData(&dataPasien,1, &banPasien)

		} else if aksi == 2 {
			adaPasien = false
			fmt.Print("Masukan Nama Pasien yang ingin dicari: ")
			fmt.Scan(&namPasien)
			program.FindData(&dataPasien, banPasien, namPasien, &adaPasien)
			fmt.Println()
			
		} else if aksi == 3 {
			adaPasien = false
			fmt.Print("Masukan Nama Pasien yang ingin diubah: ")
			fmt.Scan(&namPasien)
			program.EditData(&dataPasien, banPasien, namPasien, &adaPasien)
			if adaPasien {
				fmt.Print("Data berhasil di ubah\n")
			} else {
				fmt.Print("Data tidak ditemukan\n")
			}
		} else if aksi == 4 {
			var k int
			fmt.Print("Masukan Nama Pasien yang ingin dihapus: ")
			fmt.Scan(&namPasien)
			program.HapusData(&dataPasien, &banPasien, &k, namPasien)
		} else if aksi == 5 {
			fmt.Println("Urutan periksa pasien")
			program.UrutAntrian(&dataPasien, banPasien)
		} else if aksi==6{
			henti = false
		}
	}

}
