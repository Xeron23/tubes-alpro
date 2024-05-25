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

// func FindDataUmur(Arr *Arr, n,umur int, pas *bool) {
// 	var result int
// 	result = findBinary(Arr,n,umur)
// 	if result!= -1 {
// 		fmt.Println("Data ditemukan")
// 		fmt.Printf("%s %d %s %s %s\n", Arr[result].nama, Arr[result].umur, Arr[result].asal, Arr[result].tanggalLahir, Arr[result].golongan_riwayat_penyakit)
// 			*pas = true
// 	}
// 	if !*pas {
// 		fmt.Println("Data tidak ditemukan")
// 	}
// }

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

// func FindDataUmur(Arr *Arr, n,umur int, pas *bool) {
// 	findBinary(Arr, n, umur, pas)
// 	if !*pas {
// 		fmt.Println("Data tidak ditemukan")
// 	}
// }

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

func urutUmur(Arr *Arr, n int) {
	var pass, i int
	var temps Data
	pass = 1
	for pass < n {
		i = pass
		temps = Arr[i]
		for i>0 && temps.umur < Arr[i-1].umur {
			Arr[i] = Arr[i-1]
			i--
		}
		Arr[i] = temps
		pass++
	}
}
func UrutAntrian(Arr *Arr, n int) {
	var temps Data
	for i := 0; i < n-1; i++ {
		// var j int
		// i = i+1
		for j := i +1; j < n; j++ {
			if Arr[i].keparahan < Arr[j].keparahan {
				temps = Arr[i]
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

func findBinary(Arr *Arr, n int, umur int)int {
	urutUmur(Arr, n)
	var low, mid, high int
	var result int
	result = -1
	low = 0
	high= n-1
	for low<=high && result ==-1{
		mid = (low + high) /2
		if Arr[mid].umur == umur{
			result = mid 
		}else if Arr[mid].umur < umur {
			low = mid + 1
		}else {
			high = mid - 1
		}
	}
	return result
}

// func findBinary(Arr *Arr, n int, umur int, pas *bool) {
// 	urutUmur(Arr, n)
// 	var low, mid, high int
// 	var stop bool
// 	stop = true
// 	low = 0
// 	high= n-1
// 	for low<=high && stop{
// 		mid = (low + high) /2
// 		if Arr[mid].umur == umur{
// 			fmt.Println("Data ditemukan")
// 			fmt.Printf("%s %d %s %s %s\n", Arr[mid].nama, Arr[mid].umur, Arr[mid].asal, Arr[mid].tanggalLahir, Arr[mid].golongan_riwayat_penyakit)
// 			*pas = true
// 			if Arr[mid+1].umur == umur {
// 				fmt.Printf("%s %d %s %s %s\n", Arr[mid].nama, Arr[mid].umur, Arr[mid].asal, Arr[mid].tanggalLahir, Arr[mid].golongan_riwayat_penyakit)
// 			}else{
// 				stop = false
// 			}
// 		}else if Arr[mid].umur < umur {
// 			low = mid + 1
// 		}else {
// 			high = mid - 1
// 		}
// 	}
// }