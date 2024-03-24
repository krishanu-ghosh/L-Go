package main

import "fmt"

type person struct{
	firstname string
	lastname string
	age string
}

type pet struct{
	name string
	age string
}

func (p person) sound (){
	fmt.Printf("%v says Hello\n",p.firstname)
}

func (pe pet) sound(){
	fmt.Printf("%v says vow vow\n", pe.name)
}

type household interface{
	sound()
}

func main() {
	// var a = 80
	// fmt.Println(a)
	// fmt.Printf("%T", a)
	p1 := person{
		firstname: "James",
		lastname: "Bond",
		age: "34",
	}

	pe1 := pet{
		name: "Escobar",
		age: "7",
	}

	p1.sound()
	pe1.sound()


}
