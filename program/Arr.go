package program

import (
	"fmt"
)

const NMAX int = 1000
type Data struct {
	nama, asal, tanggalLahir, golongan_riwayat_penyakit, bulan string
	keparahan                                                  int
	umur, tanggal, tahun                                       int
}

type Arr [1000]Data

// keparahan pasien
func parah(parah string)int {
	var hasil int
	if parah == "Parah" {
		hasil = 3
	}else if parah == "Sedang" {
		hasil = 2
	}else {
		hasil = 1
	}
	return hasil
}

// input menu 1
func InputData(arr *Arr, k *int) {
	var valid bool // mencetak data valid
	var valids bool // for stop loop
	valids = true
	valid = true
	for *k<NMAX == valids{
		fmt.Print("Masukan nama pasien: ")
		fmt.Scan(&arr[*k].nama)
		fmt.Print("Masukan asal pasien: ")
		fmt.Scan(&arr[*k].asal)
		fmt.Print("Masukan umur pasien: ")
		fmt.Scan(&arr[*k].umur)
		fmt.Println("Masukan tanggal lahir pasien")
		fmt.Print("Tanggal: ")
		fmt.Scan(&arr[*k].tanggal)
		fmt.Print("Bulan: ")
		fmt.Scan(&arr[*k].bulan)
		fmt.Print("Tahun: ")
		fmt.Scan(&arr[*k].tahun)
		arr[*k].tanggalLahir = fmt.Sprintf("%d-%s-%d", arr[*k].tanggal, arr[*k].bulan, arr[*k].tahun)
		fmt.Print("Masukan golongan riwayat penyakit pasien: ")
		fmt.Scan(&arr[*k].golongan_riwayat_penyakit)

		arr[*k].keparahan = parah(arr[*k].golongan_riwayat_penyakit)
		*k++
		valids = false
		for i := 0; i <*k; i++ {
			if i > 0 && arr[i].nama == arr[i-1].nama {
				valid = false
				*k--
			}
		}
		if valid {
			fmt.Print("Data yang anda masukan telah terkirim ke sistem\n")
		} else {
			fmt.Print("Data sudah ada, silahkan masukan data lain\n")
		}
		fmt.Print("\n")
	}
}

// find menu 2
func FindDataNama(Arr *Arr, n int, namaS string) {
	var hasil int
	var pas bool
	pas = false
	hasil = Find(*Arr, n, namaS)
	if hasil != -1 {
		fmt.Println("Data ditemukan")
			fmt.Printf("%s %d %s %s %s", Arr[hasil].nama, Arr[hasil].umur, Arr[hasil].asal, Arr[hasil].tanggalLahir, Arr[hasil].golongan_riwayat_penyakit)
			pas = true
	}
	if !pas {
		fmt.Println("Data tidak ditemukan")
	}
}

func FindDataGol(Arr *Arr, n int, gol string) {
	var x, pas bool
	pas = false
	x = true
	for i := 0; i < n; i++ {
		if Arr[i].golongan_riwayat_penyakit == gol {
			if x {
				fmt.Println("Data ditemukan")
				x = false
			}
			fmt.Printf("%s %d %s %s %s\n", Arr[i].nama, Arr[i].umur, Arr[i].asal, Arr[i].tanggalLahir, Arr[i].golongan_riwayat_penyakit)
			pas = true
		}
	}
	if !pas {
		fmt.Println("Data tidak ditemukan")
	}
}

func FindDataUmur(Arr *Arr, n,umur int) {
	var x, pas bool
	x = true
	for i := 0; i < n; i++ {
		if Arr[i].umur == umur {
			if x {
				fmt.Println("Data ditemukan")
				x = false
			}
			fmt.Printf("%s %d %s %s %s\n", Arr[i].nama, Arr[i].umur, Arr[i].asal, Arr[i].tanggalLahir, Arr[i].golongan_riwayat_penyakit)
			pas = true
		}
	}
	if !pas {
		fmt.Println("Data tidak ditemukan")
	}
}

// edit menu 3
func EditData(arr *Arr, n int, namaS string, pas *bool) {
	var inStr string
	var inInt, hasil int
	hasil = Find(*arr, n, namaS)

	if hasil != -1 {
		fmt.Print("Masukan nama pasien (masukan '-' jika tidak ingin dirubah): ")
			fmt.Scan(&inStr)
			if inStr != "-" {
				arr[hasil].nama = inStr
			}
			fmt.Print("Masukan asal pasien (masukan '-' jika tidak ingin dirubah):")
			fmt.Scan(&inStr)
			if inStr != "-" {
				arr[hasil].asal = inStr
			}

			fmt.Print("Masukan umur pasien (masukan '0' jika tidak ingin dirubah):")
			fmt.Scan(&inInt)
			if inInt != 0 {
				arr[hasil].umur = inInt
			}
			fmt.Println("Masukan tanggal lahir pasien")
			fmt.Print("Tanggal (masukan '0' jika tidak ingin dirubah): ")
			fmt.Scan(&inInt)
			if inInt != 0 {
				arr[hasil].tanggal = inInt
			}
			fmt.Print("Bulan (masukan '-' jika tidak ingin dirubah):")
			fmt.Scan(&inStr)
			if inStr != "-" {
				arr[hasil].bulan = inStr
			}
			fmt.Print("Tahun (masukan '0' jika tidak ingin dirubah): ")
			fmt.Scan(&inInt)
			if inInt != 0 {
				arr[hasil].tahun = inInt
			}
			arr[hasil].tanggalLahir = fmt.Sprintf("%d-%s-%d", arr[hasil].tanggal, arr[hasil].bulan, arr[hasil].tahun)
			fmt.Print("Masukan golongan riwayat penyakit pasien (masukan '-' jika tidak ingin dirubah): ")
			fmt.Scan(&inStr)
			if inStr != "-" {
				arr[hasil].golongan_riwayat_penyakit = inStr
			}
			arr[hasil].keparahan = parah(arr[hasil].golongan_riwayat_penyakit)
			*pas = true
	}
}

// print data
func PrintData(arr Arr, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d. Nama: %s Umur: %d Asal: %s Tanggal Lahir: %s Gol-riwayat: %s\n", i+1, arr[i].nama, arr[i].umur, arr[i].asal, arr[i].tanggalLahir, arr[i].golongan_riwayat_penyakit)
	}
}

// test file
func printNama(Arr Arr, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d. Nama: %s Gol: %s\n", i+1, Arr[i].nama, Arr[i].golongan_riwayat_penyakit)
	}
}

// menurutkan prioritas
func UrutAntrian(Arr *Arr, n int) {
	var temps Data
	var pass, idx, i int
	pass = 1
	for pass<n {
		idx = pass-1
		i = pass
		for i<n {
			if (Arr[idx].keparahan < Arr[i].keparahan) || (i>0 && Arr[idx].keparahan == Arr[i].keparahan && Arr[i].umur > Arr[idx].umur){
				idx = i
			}
			i++
		}
		temps = Arr[pass-1]
		Arr[pass-1] = Arr[idx]
		Arr[idx] = temps
		pass++
	}
	printNama(*Arr, n)
}



func HapusData(Arr *Arr, n *int, nmas string) {
	var hasil int
	hasil = Find(*Arr, *n, nmas)

	if hasil != -1 {
		fmt.Println("Data berhasil di hapus")
		for i := hasil; i < *n-1; i++ {
			Arr[i] = Arr[i+1]
		}
		*n--
	} else {
		fmt.Print("Data tidak di temukan\n")
	}
}

// search
func Find(Arr Arr, n int, nmas string) int {
	var hasil, i int
	hasil = -1
	i = 0
	for i < n && hasil==-1{
		if Arr[i].nama == nmas {
			hasil = i
		}
		i++
	}
	return hasil
}

// interface
func Interface() {
	fmt.Println("*************************************")
	fmt.Println("*                                   *")
	fmt.Println("*       ======================      *")
	fmt.Println("*           Selamat Datang          *")
	fmt.Println("*         Di Klinik Tong Fung       *")
	fmt.Println("*       =======================     *")
	fmt.Println("*       1. Pendaftaran Pasien       *")
	fmt.Println("*       2. Cari Pasien              *")
	fmt.Println("*       3. Edit Data Pasien         *")
	fmt.Println("*       4. Hapus Data Pasien        *")
	fmt.Println("*       5. Urutan pemeriksaan       *")
	fmt.Println("*       6. Hentikan program         *")
	fmt.Println("*       ========================    *")
	fmt.Println("*                                   *")
	fmt.Println("*************************************")
	fmt.Print("Silahkan Pilih Menu : ")
}
