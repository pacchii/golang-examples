package _6_functions

import "fmt"

func FunctionExamples() {
	variadicFunc()
	deferExample()

}

func variadicFunc() {
	fmt.Println("Hello World", 25, true, 2.386) //Hello World 25 true 2.386

	n, err := fmt.Println("Hello") //Hello

	fmt.Println(n, err) //6 <nil>

	fmt.Println(sum(1, 3, 5, 7))

	fmt.Println(sum(2, 3, 5, 8, 13, 21, 34))

	sl := []int{2, 4, 5, 6, 8}

	fmt.Println(sum(sl...)) //Unfurling slice

}

func sum(i ...int) int { //Variadic function meaning it will take any number of parameters like fmt.Println
	fmt.Println(i)
	sum := 0
	for _, k := range i {
		sum += k
	}
	return sum
}

func deferExample() {

	/* OUTPUT
	log 1
	defer 1
	*/

	defer fmt.Println("defer 1 ")
	fmt.Println("log 1")
	x := 2

	if x%2 == 0 {
		return
	}
	fmt.Println("log 2")
	defer fmt.Println("defer 2")
}
