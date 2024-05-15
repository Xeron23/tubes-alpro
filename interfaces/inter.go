package interfaces

import "fmt"

func InputInter() {
	fmt.Println("=========================================")
	fmt.Println("Selamat datang di klinik TongFung")
	fmt.Println("=========================================")

	fmt.Print("Silahkan masukan banyak pasien: ")
}


func InputPasien(n int) {
	fmt.Printf("Masukan pasien ke-%d :", n+1)
}

func HapusData() {
	fmt.Print("Masukan nama yang akan dihapus: ")
}

func EditData() {
	fmt.Print("Masukan nama yang akan diedit: ")
}