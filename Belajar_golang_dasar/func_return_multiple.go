package main

import "fmt"

func getFullName() (string, string, int) {
	return "eko", "khanedy", 2
}

func main() {
	firstname, lastname, _ := getFullName()
	fmt.Println(firstname)
	fmt.Println(lastname)
	// fmt.Println(angka)
}
