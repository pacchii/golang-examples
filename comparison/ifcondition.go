package comparison

import "fmt"

func IfCondition(x int, y int, message string) {

	if x > y {
		fmt.Println(x, " is ", message, " than ", y)
	} else if y > x {
		fmt.Println(y, " is ", message, " than ", x)
	} else {
		fmt.Println("Both are equal")
	}
}
