package comparison

import (
	"fmt"
	"math/rand"
)

func IfCondition(x int, y int, message string) {
	if x > y {
		fmt.Println(x, " is ", message, " than ", y)
	} else if y > x {
		fmt.Println(y, " is ", message, " than ", x)
	} else {
		fmt.Println("Both are equal")
	}

	if z := 2 * rand.Intn(x); z > x {
		fmt.Println("Random ", z, " is ", message, " than ", x)
	}
}
