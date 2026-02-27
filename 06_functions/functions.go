package _6_functions

import "fmt"

func FunctionExamples() {
	variadicFunc()
	deferExample()
	ananymous()
	returningFunc()
	callback()
	closure()
	recursion()
}

func recursion() {
	fmt.Println(factorial(5))
	fmt.Println(normFactorial(5))
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func normFactorial(n int) int {
	total := 1
	for i := 1; i <= n; i++ {
		total *= i
	}
	return total
}

func closure() {

	f := incrementor()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func callback() {
	doMathOperation(2, 3, add)
	doMathOperation(6, 2, subtarct)
}

func doMathOperation(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(a int, b int) int {
	return a + b
}

func subtarct(a int, b int) int {
	return a - b
}

func returningFunc() {

	f := bar()

	fmt.Printf("%T\n", f)   //func() int
	fmt.Printf("%T\n", bar) //func() func() int

}

func bar() func() int {
	return func() int {
		return 33
	}
}

func ananymous() {

	func() {
		fmt.Println("Hello Ananymous")
	}()

	func(s string) {
		fmt.Println("Ananymous with parameter", s)
	}("Hello World")
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
