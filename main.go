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
	for henti {
		program.Interface()
		fmt.Scan(&aksi)
		if aksi == 1 {
			fmt.Print("Berapa banyak pasien yang ingin di tambahkan: ")
			fmt.Scan(&banPasien)
			program.InputData(&dataPasien, banPasien)
			fmt.Print("\n\nData yang masukan telah terkirim ke sistem\n\n")
		} else if aksi == 2 {
			adaPasien = false
			fmt.Print("Masukan Nama Pasien yang ingin dicari: ")
			fmt.Scan(&namPasien)
			program.FindData(&dataPasien, banPasien, namPasien, &adaPasien)
			fmt.Println()
			if adaPasien {
				fmt.Println("Data pasien ditemukan")
			} else if !adaPasien && banPasien == 0 {
				fmt.Println("Data pasien tidak ada, silahkan masukan pasien")
				fmt.Print("Berapa banyak pasien yang ingin di tambahkan: ")
				fmt.Scan(&banPasien)
				program.InputData(&dataPasien, banPasien)
				fmt.Print("\n\nData yang masukan telah terkirim ke sistem\n\n")
			} else {
				fmt.Println("Data pasien tidak ada")

			}
		} else if aksi == 3 {
			adaPasien = false
			fmt.Print("Masukan Nama Pasien yang ingin dihapus: ")
			fmt.Scan(&namPasien)
			program.EditData(&dataPasien, banPasien, namPasien, &adaPasien)
			if adaPasien {
				fmt.Print("Data berhasil di ubah\n")
			} else {
				fmt.Print("Data tidak ditemukan\n")
			}
		} else if aksi == 4 {
			var k int
			fmt.Print("Masukan Nama Pasien yang ingin diubah: ")
			fmt.Scan(&namPasien)
			program.HapusData(&dataPasien, &banPasien, &k, namPasien)
		} else if aksi == 5 {
			program.UrutAntrian(&dataPasien, banPasien)
		} else {
			henti = false
		}
	}

}
