package main

import "fmt"

func getHello(name string) string {
	return "hallo " + name
}

func main() {
	asd := getHello("Eko")
	fmt.Println(asd)
}
