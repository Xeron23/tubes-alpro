package main

import (
	"fmt"
	"test-tubes/program"
)

var aksi, banPasien, data, umur int
var dataPasien program.Arr
var namPasien string
var adaPasien, henti bool

func main() {
	henti = true
	banPasien = 0
	for henti {
		program.Interface()
		fmt.Scan(&aksi)
		if aksi == 1 {
			fmt.Println("\nSILAHKAN MASUKAN DATA PASIEN")
			program.InputData(&dataPasien, &banPasien)
		} else if aksi == 2 {
			fmt.Print("Silahkan pilih berdasarkan 1). nama 2). riwayat penyakit 3).umur\n: ")
			fmt.Scan(&data)
			if data == 1 {
				fmt.Print("Masukan Nama Pasien yang ingin dicari: ")
				fmt.Scan(&namPasien)
				program.FindDataNama(&dataPasien, banPasien, namPasien)
				fmt.Println()
			}else if data == 2{
				fmt.Print("Masukan riwayat penyakit Pasien yang ingin dicari (Parah/Sedang/Ringan): ")
				fmt.Scan(&namPasien)
				program.FindDataGol(&dataPasien, banPasien, namPasien)
			}else {
				fmt.Print("Masukan umur pasien yang ingin dicari:  ")
				fmt.Scan(&umur)
				program.FindDataUmur(&dataPasien, banPasien, umur)
			}
		} else if aksi == 3 {
			adaPasien = false
			fmt.Print("Masukan Nama Pasien yang ingin diubah: ")
			fmt.Scan(&namPasien)
			program.EditData(&dataPasien, banPasien,namPasien, &adaPasien)
			if adaPasien {
				fmt.Print("Data berhasil di ubah\n")
			} else {
				fmt.Print("Data tidak ditemukan\n")
			}
		} else if aksi == 4 {
			fmt.Print("Masukan Nama Pasien yang ingin dihapus: ")
			fmt.Scan(&namPasien)
			program.HapusData(&dataPasien, &banPasien, namPasien)
		} else if aksi == 5 {
			fmt.Println("Urutan periksa pasien")
			program.UrutAntrian(&dataPasien, banPasien)
		} else if aksi==6 {
			henti = false
			fmt.Print("Terimakasi")
		}else {
			program.PrintData(dataPasien, banPasien)
		}
	}
}
