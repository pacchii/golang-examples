package array

import "fmt"

func ArrayExample() {

	//array literal

	a := [5]int{1, 2, 3, 5, 4}

	fmt.Println(a)

	b := [...]string{"alpha", "beta", "gamma", "delta"}
	fmt.Println(b)

	var c [5]int
	c[4] = 100
	c[2] = 20
	c[1] = 10

	fmt.Println(c)

	fmt.Printf("%T\n %T\n %T\n", a, b, c)

	//we can assign one array to another if their size is same
	c = a
	fmt.Println(c)
}
