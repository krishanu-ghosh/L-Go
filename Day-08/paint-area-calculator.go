package main

import (
	//"bufio"
	"fmt"
	"math"
	//"os"
)

func painAreaCalculator(h, w, a float64) {
	var cansReq float64
	areaOfWall := h * w
	cansReq = areaOfWall / a
	cansToBuy := int(math.Ceil(cansReq))
	fmt.Printf("You need exactly %.2f cans of paint.\nSo you need to buy %v cans of paint.\n", cansReq, cansToBuy)
	fmt.Printf("This much paint will be left over : %.2f l\n", float64(cansToBuy)-cansReq)
}

func main() {
	var height, weight, area float64
	fmt.Printf("Height of wall (m): ")
	//reader := bufio.NewReader(os.Stdin)
	fmt.Scanf("%f", &height)
	fmt.Printf("Width of wall (m): ")
	fmt.Scanf("%f", &weight)
	fmt.Printf("Area covered by 1 can (m.sq): ")
	fmt.Scanf("%f", &area)

	//fmt.Printf("Height = %v\nWidth = %v\nArea = %v\n", height, weight, area)

	painAreaCalculator(height, weight, area)
}
