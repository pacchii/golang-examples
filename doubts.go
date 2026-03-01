package main

import "fmt"

func DeferExample() {

	defer fmt.Println("Last Statement")

	var a, b float64

	a = 12.0
	b = 0.0

	fmt.Println(a / b)

	panic("Unexpected occured")

}
