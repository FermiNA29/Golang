package main

import "fmt"

func getComplateName() (firstname, middlename, lastname string) {
	firstname = "eko"
	middlename = "kurniawan"
	lastname = "khannedy"
	return
}

func main() {
	firstname, middlename, lastname := getComplateName()
	fmt.Println(firstname)
	fmt.Println(middlename)
	fmt.Println(lastname)
}
