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

// input menu 1
func InputData(arr *Arr, k *int) {
	var valid bool
	var valids bool
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
		if arr[*k].golongan_riwayat_penyakit == "Parah" {
			arr[*k].keparahan = 3
		} else if arr[*k].golongan_riwayat_penyakit == "Sedang" {
			arr[*k].keparahan = 2
		} else {
			arr[*k].keparahan = 1
		}
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
func FindDataNama(Arr *Arr, n int, namaS string, pas *bool) {
	for i := 0; i < n && !*pas; i++ {
		if Arr[i].nama == namaS {
			fmt.Println("Data ditemukan")
			fmt.Printf("%s %d %s %s %s", Arr[i].nama, Arr[i].umur, Arr[i].asal, Arr[i].tanggalLahir, Arr[i].golongan_riwayat_penyakit)
			*pas = true
		}
	}
	if !*pas {
		fmt.Println("Data tidak ditemukan")
	}
}

func FindDataGol(Arr *Arr, n int, gol string, pas *bool) {
	var x bool
	x = true
	for i := 0; i < n; i++ {
		if Arr[i].golongan_riwayat_penyakit == gol {
			if x {
				fmt.Println("Data ditemukan")
				x = false
			}
			fmt.Printf("%s %d %s %s %s\n", Arr[i].nama, Arr[i].umur, Arr[i].asal, Arr[i].tanggalLahir, Arr[i].golongan_riwayat_penyakit)
			*pas = true
		}
	}
	if !*pas {
		fmt.Println("Data tidak ditemukan")
	}
}



func FindDataUmur(Arr *Arr, n,umur int, pas *bool) {
	var x bool
	x = true
	for i := 0; i < n; i++ {
		if Arr[i].umur == umur {
			if x {
				fmt.Println("Data ditemukan")
				x = false
			}
			fmt.Printf("%s %d %s %s %s\n", Arr[i].nama, Arr[i].umur, Arr[i].asal, Arr[i].tanggalLahir, Arr[i].golongan_riwayat_penyakit)
			*pas = true
		}
	}
	if !*pas {
		fmt.Println("Data tidak ditemukan")
	}
}

// edit menu 3
func EditData(arr *Arr, n int, namaS string, pas *bool) {
	var inStr string
	var inInt int
	// scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n && !*pas; i++ {
		// var inInt int
		if arr[i].nama == namaS {
			fmt.Print("Masukan nama pasien (masukan '-' jika tidak ingin dirubah): ")
			fmt.Scan(&inStr)
			if inStr != "-" {
				arr[i].nama = inStr
			}
			fmt.Print("Masukan asal pasien (masukan '-' jika tidak ingin dirubah):")
			fmt.Scan(&inStr)
			if inStr != "-" {
				arr[i].asal = inStr
			}

			fmt.Print("Masukan umur pasien (masukan '0' jika tidak ingin dirubah):")
			fmt.Scan(&inInt)
			if inInt != 0 {
				arr[i].umur = inInt
			}
			fmt.Println("Masukan tanggal lahir pasien")
			fmt.Print("Tanggal (masukan '0' jika tidak ingin dirubah): ")
			fmt.Scan(&inInt)
			if inInt != 0 {
				arr[i].tanggal = inInt
			}
			fmt.Print("Bulan (masukan '-' jika tidak ingin dirubah):")
			fmt.Scan(&inStr)
			if inStr != "-" {
				arr[i].bulan = inStr
			}
			fmt.Print("Tahun (masukan '0' jika tidak ingin dirubah): ")
			fmt.Scan(&inInt)
			if inInt != 0 {
				arr[i].tahun = inInt
			}
			arr[i].tanggalLahir = fmt.Sprintf("%d-%s-%d", arr[i].tanggal, arr[i].bulan, arr[i].tahun)
			fmt.Print("Masukan golongan riwayat penyakit pasien (masukan '-' jika tidak ingin dirubah): ")
			fmt.Scan(&inStr)
			if inStr != "-" {
				arr[i].golongan_riwayat_penyakit = inStr
			}
			if arr[i].golongan_riwayat_penyakit == "Parah" {
				arr[i].keparahan = 3
			} else if arr[i].golongan_riwayat_penyakit == "Sedang" {
				arr[i].keparahan = 2
			} else {
				arr[i].keparahan = 1
			}
			*pas = true
			fmt.Print("\n")
		}
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
			if Arr[idx].keparahan < Arr[i].keparahan {
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


func HapusData(Arr *Arr, n *int, k *int, nmas string) {
	Find(*Arr, *n, k, nmas)
	if *k != -1 {
		fmt.Println("Data berhasil di hapus")
		for i := *k; i < *n-1; i++ {
			Arr[i] = Arr[i+1]
		}
		*n--
	} else {
		fmt.Print("Data tidak di temukan\n")
	}
}

func Find(Arr Arr, n int, k *int, nmas string) int {
	*k = -1
	for i := 0; i < n && *k==-1; i++ {
		if Arr[i].nama == nmas {
			*k = i
		}
	}
	return *k
}

