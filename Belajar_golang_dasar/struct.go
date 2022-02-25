package main

import "fmt"

//struct merupakan template data yg digunakan untuk menggabungkan nol atau lebih tipe data
type Customer struct {
	Name, Address string
	Age           int
}

// struct method
func (customer Customer) sayHi(name string) {
	fmt.Println("hello " + name + " my name is " + customer.Name)
}
func (a Customer) sayHuuu() {
	fmt.Println("hello my name is " + a.Name)
}

func main() {
	var eko Customer
	eko.Name = "eki"
	eko.Address = "jl bdg"
	eko.Age = 1

	eko.sayHi("joko")
	eko.sayHuuu()
	fmt.Println(eko)

	joko := Customer{
		Name:    "joko",
		Address: "cirebon",
		Age:     2,
	}

	fmt.Println(joko)

	budi := Customer{"budi", "jl sumatra", 3}
	fmt.Println(budi)

}
