package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	// var slice2 = months[4:]
	fmt.Println(slice1)
	// fmt.Println(slice2)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// months[5] = "diubah"
	// fmt.Println(slice1)

	// slice1[0] = "mei lagi"
	// fmt.Println(months)

	slice2 := months[10:]
	fmt.Println(slice2)

	slice3 := append(slice2, "Eko")
	fmt.Println(slice3)
	slice3[1] = "bukan desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Fermi"
	newSlice[1] = "Naufal"
	fmt.Println(newSlice)

	// copy slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	// perbedaan array dengan slice
	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
