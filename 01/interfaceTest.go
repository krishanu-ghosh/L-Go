package main

import "fmt"

type Specs interface {
	ShowSpecs()
}

type Bike struct {
	model        string
	wheel        int
	cylinders    int
	displacement int
}

type Car struct {
	model  string
	engine string
	tyre   string
}

func (b Bike) ShowSpecs() {
	fmt.Println("Bike Model:", b.model)
	fmt.Println("Wheel:", b.wheel)
	fmt.Println("Cylinder:", b.cylinders)
	fmt.Println("Displacement:", b.displacement)
}

func (c Car) ShowSpecs() {
	fmt.Println("Car Model:", c.model)
	fmt.Println("Engine:", c.engine)
	fmt.Println("Tyre:", c.tyre)
}

func main() {
	bike := Bike{model: "Kawasaki", wheel: 2, cylinders: 4, displacement: 998}
	car := Car{model: "Shelby", engine: "V12", tyre: "Slicks"}
	specs := []Specs{bike, car}
	for _, spec := range specs {
		fmt.Println(spec)
		spec.ShowSpecs()
	}
}
