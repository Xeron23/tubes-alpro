package main

import (
	"fmt"
	"test-tubes/fileTest"
)
func main() {
	var aksi,banPasien int
	var dataPasien fileTest.Arr
	var namPasien string
	var adaPasien bool

	for {
		fileTest.Interface()
		fmt.Scan(&aksi)
		if aksi ==1 {
			fmt.Print("Berapa banyak pasien yang ingin di tambahkan: ")
			fmt.Scan(&banPasien)
			fileTest.InputData(&dataPasien, banPasien)
			fmt.Print("Data yang masukan telah terkirim ke sistem\n\n")
		}else if aksi ==2 {
			adaPasien = false
			fmt.Print("Masukan Nama Pasien yang ingin dicari: ")
			fmt.Scan(&namPasien)
			fileTest.FindData(&dataPasien, banPasien, namPasien, &adaPasien)
			if adaPasien {
				fmt.Println("Data pasien ditemukan")
			}else if !adaPasien && banPasien==0 {
				fmt.Println("Data pasien tidak ada, silahkan masukan pasien")
			}else {
				fmt.Println("Data pasien tidak ada")
			}
		}else if aksi==3 {
			adaPasien = false
			fmt.Print("Masukan Nama Pasien yang ingin diubah: ")
			fmt.Scan(&namPasien)
			fileTest.EditData(&dataPasien, banPasien, namPasien, &adaPasien)
			if adaPasien {
				fmt.Print("Data berhasil di ubah\n")
			}else {
				fmt.Print("Data tidak ditemukan\n")
			}
		}else if aksi==4 {
			var k int
			fmt.Print("Masukan Nama Pasien yang ingin diubah: ")
			fmt.Scan(&namPasien)
			fileTest.HapusData(&dataPasien, &banPasien, &k, namPasien)
		}else {
			fileTest.PrintData(dataPasien, banPasien)
		}
	}

}
