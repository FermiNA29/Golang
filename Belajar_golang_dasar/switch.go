package main

import "fmt"

func main() {
	name := "eko"

	switch name {
	case "Eko":
		fmt.Println("Hello " + name)
	case "joko":
		fmt.Println("Hello " + name)
	default:
		fmt.Println("Hello default")
	}

	// switch length := len(name); length {
	// case 3:
	// 	fmt.Println("Nama Terlalu panjang")
	// case 2:
	// 	fmt.Println("nama sudah benar")
	// }

	length := len(name)
	switch {
	case length > 5:
		fmt.Println("nama terlalu panjang")
	case length <= 5:
		fmt.Println("nama kurang dari 5 hruruf")
	default:
		fmt.Println("nama sudah benar")
	}
}
