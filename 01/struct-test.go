package main

import "fmt"

type Person struct {
	firstname, lastname string
	favIceCream         []string
}

func (p *Person) printData() {
	fmt.Printf("%s %s\n", p.firstname, p.lastname)
	for _, v := range p.favIceCream {
		fmt.Println(v)
	}
}

func main() {
	p1 := Person{
		firstname:   "John",
		lastname:    "Doe",
		favIceCream: []string{"mango", "chocolate", "butterscotch"},
	}

	p2 := Person{
		firstname:   "Jane",
		lastname:    "Doe",
		favIceCream: []string{"banana", "apple", "strawberry"},
	}

	p1.printData()
	p2.printData()

}
