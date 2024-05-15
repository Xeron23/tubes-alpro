package fileTest

import "fmt"

type data struct {
	nama string
	umur int
}

type Arr [100]data

func InputData(arr *Arr, n int) {
	fmt.Scan(&arr[n].nama, &arr[n].umur)
}
var Datas Arr

func DeleteData(Arr *Arr, n *int, nama string) {
	var test bool
	for i := 0; i < *n ||i >= len(Arr); i++ {
		if Arr[i].nama == nama {
			copy(Arr[i:], Arr[i+1:*n])
			Arr[*n-1] = data{} // Menghapus data terakhir setelah geser
			*n--
			test = true
			break
		}
		if !test && i==*n-1{
			fmt.Println("Nama tidak ditemukan.")
		}
	}
}

func EditData(Arr *Arr, n int, namaS string) {
	var nama string
	var umur int
	for i := 0; i < n; i++ {
		if Arr[i].nama==namaS {
			fmt.Scan(&nama, &umur)
			Arr[i].nama = nama
			Arr[i].umur = umur
		}
	}
}
func PrintData(arr Arr, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("Nama: %s Umur: %d\n", arr[i].nama, arr[i].umur)
	}
}
