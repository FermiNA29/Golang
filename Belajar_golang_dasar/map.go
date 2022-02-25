package main

import "fmt"

func main() {
	var person map[string]string = map[string]string{ //atau person := map[string]string{}
		"name":    "eko",
		"address": "subang",
	}

	person["title"] = "programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go-lang"
	book["author"] = "Eko"
	book["ups"] = "salah"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")
	fmt.Println(book)
	fmt.Println(len(book))
}
