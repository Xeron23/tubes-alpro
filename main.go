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
var data string
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
			fmt.Print("Silahkan pilih berdasarkan nama atau riwayat penyakit: ")
			fmt.Scan(&data)
			if data == "nama" {
				fmt.Print("Masukan Nama Pasien yang ingin dicari: ")
				fmt.Scan(&namPasien)
				program.FindDataNama(&dataPasien, banPasien, namPasien, &adaPasien)
				fmt.Println()
			}else {
				fmt.Print("Masukan riwayat penyakit Pasien yang ingin dicari: ")
				fmt.Scan(&namPasien)
				program.FindDataGol(&dataPasien, banPasien, namPasien, &adaPasien)
			}
			
			
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
