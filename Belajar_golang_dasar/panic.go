package main

import "fmt"

// panic digunkan untuk memberhentikan aplikasi jika terjadi error tapi tetap mengeksekusi defer

func endApp() {
	pesan := recover() //recover digunakan untuk menangkap data panic sehingga aplikasi tetap berjalan
	if pesan != nil {
		fmt.Println("error dengan pesan: ", pesan)

	}
	fmt.Println("äplikasi  selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi Error")
	}
	fmt.Println("aplikasi berjalan")
}

func main() {
	runApp(false)
}
