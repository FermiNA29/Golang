package main

import "fmt"

func main() {

	var names [4]string

	names[0] = "Asd"
	names[1] = "sdf"
	names[2] = "asd"

	fmt.Println(names[3])

	var values = [3]int{
		90,
		95,
		80,
	}

	fmt.Println(values)
	fmt.Println(len(values)) //menghitung panjang array
}
