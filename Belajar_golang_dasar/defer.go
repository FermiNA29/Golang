package main

import "fmt"

// defer yaitu memanggil function setelah menjkalankan suatu function walaupun error

func logging() {
	fmt.Println("selesai memanggil function")
}

func runApplication() {
	defer logging()
	fmt.Println("run apllication")
}

func main() {
	runApplication()
}
