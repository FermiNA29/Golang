package main

import "fmt"

func main() {
	var name = "rudi"

	if name == "eko" {
		fmt.Println("hallo eko")
	} else if name == "rudi" {
		fmt.Println("hallo " + name)
	} else {
		fmt.Println("aleluya")
	}

	// var length := len(name)
	if length := len(name); length >= 5 { //sort statement
		fmt.Println("name lebih dari 5 karakter")
	} else {
		fmt.Println("nama kurang dari 5 karakter")
	}
}
