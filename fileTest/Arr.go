package fileTest

import "fmt"

type Data struct {
	nama,asal, tanggalLahir, golongan_riwayat_penyakit string
	umur int
}

type Arr [100]Data

func InputData(arr *Arr, n int) {
	for i := 0; i < n; i++ {
		fmt.Println("===================")
		fmt.Printf("    Pasien ke-%d   \n", i+1)
		fmt.Println("===================")
		fmt.Print("Masukan nama pasien: ")
		fmt.Scan(&arr[i].nama)
		fmt.Print("Masukan asal pasien: ")
		fmt.Scan(&arr[i].asal)
		fmt.Print("Masukan umur pasien: ")
		fmt.Scan(&arr[i].umur)
		fmt.Print("Masukan tanggal lahir pasien: ")
		fmt.Scan(&arr[i].tanggalLahir)
		fmt.Print("Masukan golongan riwayat penyakit pasien: ")
		fmt.Scan(&arr[i].golongan_riwayat_penyakit)
		fmt.Print("\n")
		
	}
}
var Datas Arr

func DeleteData(Arr *Arr, n *int, nama string) {
	var test bool
	for i := 0; i < *n ||i >= len(Arr); i++ {
		if Arr[i].nama == nama {
			copy(Arr[i:], Arr[i+1:*n])
			Arr[*n-1] = Data{} // Menghapus data terakhir setelah geser
			*n--
			test = true
			break
		}
		if !test && i==*n-1{
			fmt.Println("Nama tidak ditemukan.")
		}
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
	for i := 0; i < n && !*pas; i++ {
		if arr[i].nama == namaS {
			fmt.Print("Masukan nama pasien: ")
			fmt.Scan(&arr[i].nama)
			fmt.Print("Masukan asal pasien: ")
			fmt.Scan(&arr[i].asal)
			fmt.Print("Masukan umur pasien: ")
			fmt.Scan(&arr[i].umur)
			fmt.Print("Masukan tanggal lahir pasien: ")
			fmt.Scan(&arr[i].tanggalLahir)
			fmt.Print("Masukan golongan riwayat penyakit pasien: ")
			fmt.Scan(&arr[i].golongan_riwayat_penyakit)
			fmt.Print("\n")
			*pas = true
		}
	}
}

func PrintData(arr Arr, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d. Nama: %s Umur: %d Asal: %s Tanggal Lahir: %s Gol-riwayat: %s\n",i+1, arr[i].nama, arr[i].umur, arr[i].asal, arr[i].tanggalLahir, arr[i].golongan_riwayat_penyakit)
	}
}

func cetakAntrian() {

}

func Interface() {
		fmt.Println("*************************************")
		fmt.Println("*                                   *")
		fmt.Println("*       ======================      *")
		fmt.Println("*           Selamat Datang          *")
		fmt.Println("*         Di Klinik Tong Fung       *")
		fmt.Println("*       =======================     *")
		fmt.Println("*       1. Pendaftaran Pasien       *")
		fmt.Println("*       2. Mencari Pasien           *")
		fmt.Println("*       3. Mengedit Data Pasien     *")
		fmt.Println("*       4. Menghapus Data Pasien    *")
		fmt.Println("*       5. Mengurutkan Pasien       *")
		fmt.Println("*       ========================    *")
		fmt.Println("*                                   *")
		fmt.Println("*************************************")
		fmt.Print("Silahkan Pilih Menu : ")
}


func HapusData(Arr *Arr, n int, k *int, nmas string) {
	Find(Arr, n, k,nmas)
	if *k!=-1 {
		fmt.Println("Data ditemukan")
		for i := *k; i < n-1; i++ {
			Arr[i] = Arr[i+1]
		}
	}else {
		fmt.Print("Data tidak di temukan")
	}
}

func Find(Arr *Arr, n int, k *int, nmas string) {
	*k = -1
	for i := 0; i < n; i++ {
		if Arr[i].nama==nmas {
			*k = i
			break
		}
	}
}