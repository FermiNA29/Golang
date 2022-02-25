package main

import "fmt"

func main() {
	// array
	// languages := [...]string{
	// 	"rubu",
	// 	"php",
	// 	"java",
	// 	"java",
	// }

	// for _, v := range languages {
	// 	fmt.Println(v)
	// }

	// slice
	// var gamingConsole []string
	// gaming := append(gamingConsole, "playstation4", "nintendo switch") //memasukan ke dalam slice
	// fmt.Println(gaming)

	// for _, v := range gaming {
	// 	fmt.Println(v)
	// }

	// gamingC := []string{"playstation4", "Xbox"}
	// fmt.Println(gamingC)

	// map
	var myMap map[string]int
	myMap = map[string]int{}

	myMap["ruby"] = 9
	myMap["go"] = 10
	myMap["javascript"] = 8

	fmt.Println(myMap["go"])

	myMups := map[string]string{
		"ruby":       "is beautiful",
		"go":         "is superfast",
		"JavaScript": "hype",
	}
	fmt.Println(myMups)

	for key, v := range myMups {
		fmt.Println("key :", key, "value :", v)
	}

	fmt.Println("===============")
	delete(myMups, "ruby")
	for key, v := range myMups {
		fmt.Println("key :", key, "value :", v)
	}
	// cek apakah map mengandung suatu key tertentu
	value, isAvailable := myMups["python"]
	fmt.Println(value)
	fmt.Println(isAvailable)

}
