package main

import "fmt"

//interface adalah tipe data abstract. Berisikan definisi method. Biasanya digunakan sebagai kontrak
type HasName interface {
	GetName() string
}

func sayHello(hasname HasName) {
	fmt.Println("Hello ", hasname.GetName())
}

//struct
type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var eko Person
	eko.Name = "Eko"

	sayHello(eko)

	cat := Animal{
		Name: "pushy",
	}
	sayHello(cat)
}
