package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		fmt.Println("perulangan ke ", i)
	}

	slice := []string{"EKo", "Kurniawan", "khannedy"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for _, value := range slice {
		// fmt.Println("Index", i, "=", value)
		fmt.Println(value)
	}

	var person map[string]string = map[string]string{ //atau person := map[string]string{}
		"name":    "eko",
		"address": "subang",
	}

	for key, value := range person {
		// fmt.Println("Index", i, "=", value)
		fmt.Println(key, "", value)
	}
}
