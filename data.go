package main

import (
	"fmt"
	"test-tubes/fileTest"
	"test-tubes/interfaces"
)

func main() {
	var nama string
	// input data
	interfaces.InputInter()
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		// Dari interface
		interfaces.InputPasien(i)

		// Dari fileTest
		fileTest.InputData(&fileTest.Datas, i)
	}

	// print data
	fileTest.PrintData(fileTest.Datas, n)

	// hapus data
	interfaces.HapusData()
	fmt.Scan(&nama)
	fileTest.DeleteData(&fileTest.Datas, &n, nama)
	fileTest.PrintData(fileTest.Datas, n)
	
	// edit data
	interfaces.EditData()
	fmt.Scan(&nama)
	fileTest.EditData(&fileTest.Datas, n, nama)
	fileTest.PrintData(fileTest.Datas, n)
}
