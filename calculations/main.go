package main

import (
	"fmt"
	"math"
)

func main() {
	pi := math.Pi
	r := 2.25
	fmt.Print("Enter height: ")
	var h float64
	fmt.Scanln(&h)

	v := (pi * math.Pow(r, 2) * h)
	fmt.Println("Volume of a Cylinder: ", math.Round(v), "IN")
}
