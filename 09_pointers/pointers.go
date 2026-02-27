package _9_pointers

import "fmt"

func PointerExample() {

	x := 20

	y := &x
	fmt.Println(x, &x)

	fmt.Printf("%v\t%T\n", &x, &x) //0xc00005a270   *string

	fmt.Println(x, y, *y) //20, 0xc00005a270 20

	*y = 10

	fmt.Println(x, y, *y) // 10, 0xc00005a270 10

	x = 30
	fmt.Println(x, &x) //30, 0xc00005a270
	deltaChange(&x)
	fmt.Println(x, &x) //20, 0xc00005a270
}

func deltaChange(x *int) {
	*x = 20
}
