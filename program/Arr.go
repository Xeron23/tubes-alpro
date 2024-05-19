package program

import (
	"fmt"
)

type Data struct {
	nama, asal, tanggalLahir, golongan_riwayat_penyakit, bulan string
	keparahan                                                  int
	umur, tanggal, tahun                                       int
}

type Arr [1000]Data

func InputData(arr *Arr, n int, k *int) {
	var valid bool
	valid = true
	var i int
	for i = *k; i < n+*k && valid; i++ {
		fmt.Println("===================")
		fmt.Printf("    Pasien ke-%d   \n", i+1)
		fmt.Println("===================")
		fmt.Print("Masukan nama pasien: ")
		fmt.Scan(&arr[i].nama)
		fmt.Print("Masukan asal pasien: ")
		fmt.Scan(&arr[i].asal)
		fmt.Print("Masukan umur pasien: ")
		fmt.Scan(&arr[i].umur)
		fmt.Println("Masukan tanggal lahir pasien")
		fmt.Print("Tanggal: ")
		fmt.Scan(&arr[i].tanggal)
		fmt.Print("Bulan: ")
		fmt.Scan(&arr[i].bulan)
		fmt.Print("Tahun: ")
		fmt.Scan(&arr[i].tahun)
		arr[i].tanggalLahir = fmt.Sprintf("%d-%s-%d", arr[i].tanggal, arr[i].bulan, arr[i].tahun)
		fmt.Print("Masukan golongan riwayat penyakit pasien: ")
		fmt.Scan(&arr[i].golongan_riwayat_penyakit)
		if arr[i].golongan_riwayat_penyakit == "Parah" {
			arr[i].keparahan = 3
		} else if arr[i].golongan_riwayat_penyakit == "Sedang" {
			arr[i].keparahan = 2
		} else {
			arr[i].keparahan = 1
		}
		if i>0 && arr[i].nama == arr[i-1].nama {
			valid = false
			n = 0
		}
		fmt.Print("\n")
		
	}
	*k += n
}

var Datas Arr

func DeleteData(Arr *Arr, n *int, nama string) {
	var test bool
	for i := 0; i < *n || i >= len(Arr); i++ {
		if Arr[i].nama == nama {
			copy(Arr[i:], Arr[i+1:*n])
			Arr[*n-1] = Data{} // Menghapus data terakhir setelah geser
			*n--
			test = true
			break
		}
		if !test && i == *n-1 {
			fmt.Println("Nama tidak ditemukan.")
		}
	}
}
func DeleteDat(Arr *Arr, n *int, nama string, k *int) {
	var data int
	data = Find(*Arr, *n, k, nama)
	if data != -1 {
		for i := data; i < *n; i++ {
			Arr[i] = Arr[i+1]
		}
		*n--
	} else {

	}
}

func FindData(Arr *Arr, n int, namaS string, pas *bool) {
	for i := 0; i < n && !*pas; i++ {
		if Arr[i].nama == namaS {
			fmt.Printf("%s %d %s %s %s", Arr[i].nama, Arr[i].umur, Arr[i].asal, Arr[i].tanggalLahir, Arr[i].golongan_riwayat_penyakit)
			*pas = true
		}
	}
}

func EditData(arr *Arr, n int, namaS string, pas *bool) {
	var inStr string
	var inInt int
	// scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n && !*pas; i++ {
		// var inInt int
		if arr[i].nama == namaS {
			fmt.Print("Masukan nama pasien: ")
			fmt.Scan(&inStr)
			if inStr != "-" {
				arr[i].nama = inStr
			}
			fmt.Print("Masukan asal pasien: ")
			fmt.Scan(&inStr)
			if inStr != "-" {
				arr[i].asal = inStr
			}

			fmt.Print("Masukan umur pasien: ")
			fmt.Scan(&inInt)
			if inStr != "-" {
				arr[i].umur = inInt
			}
			fmt.Println("Masukan tanggal lahir pasien")
			fmt.Print("Tanggal: ")
			fmt.Scan(&inInt)
			if inStr != "-" {
				arr[i].tanggal = inInt
			}
			fmt.Print("Bulan: ")
			fmt.Scan(&inStr)
			if inStr != "-" {
				arr[i].bulan = inStr
			}
			fmt.Print("Tahun: ")
			fmt.Scan(&inInt)
			if inStr != "-" {
				arr[i].tahun = inInt
			}
			arr[i].tanggalLahir = fmt.Sprintf("%d-%s-%d", arr[i].tanggal, arr[i].bulan, arr[i].tahun)
			fmt.Print("Masukan golongan riwayat penyakit pasien: ")
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
			fmt.Print("\n")
		}
	}
}

func PrintData(arr Arr, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d. Nama: %s Umur: %d Asal: %s Tanggal Lahir: %s Gol-riwayat: %s\n", i+1, arr[i].nama, arr[i].umur, arr[i].asal, arr[i].tanggalLahir, arr[i].golongan_riwayat_penyakit)
	}
}

func printNama(Arr Arr, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d. Nama: %s Gol: %s\n", i+1, Arr[i].nama, Arr[i].golongan_riwayat_penyakit)
	}
}
func UrutAntrian(Arr *Arr, n int) {
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if Arr[i].keparahan < Arr[j].keparahan {
				temps := Arr[i]
				Arr[i] = Arr[j]
				Arr[j] = temps
			}
		}
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
	var result int
	result = Find(*Arr, *n, k, nmas)
	if result != -1 {
		fmt.Println("Data berhasil di hapus\n")
		for i := result; i < *n-1; i++ {
			Arr[i] = Arr[i+1]
		}
		*n--
	} else {
		fmt.Print("Data tidak di temukan\n")
	}
}

func Find(Arr Arr, n int, k *int, nmas string) int {
	*k = -1
	for i := 0; i < n; i++ {
		if Arr[i].nama == nmas {
			*k = i
			break
		}
	}
	return *k
}

