package _0_generics

import (
	"log"
)

func AddI(a, b int) int {
	return a + b
}
func AddF(a, b float64) float64 {
	return a + b
}

// Type Constraint
func AddTC[T int | float64](a, b T) T {
	return a + b
}

// Type Set interfacce
type numbers interface {
	int | float64
}

func AddTSI[T numbers](a, b T) T {
	return a + b
}

func Print[T any](a T) {
	log.Println("INFO", a)
}

// Underlying type constraint
type underlyingtypenumbers interface {
	~int | ~float64
}

func AddUTSI[T underlyingtypenumbers](a, b T) T {
	return a + b
}

func TypeConstraint() {

	Print(AddI(2, 3))     // 2009/11/10 23:00:00 INFO 5
	Print(AddF(2.2, 3.9)) // 2009/11/10 23:00:00 INFO 6.1

	Print(AddTC(2, 3))     // 2009/11/10 23:00:00 INFO 5
	Print(AddTC(2.2, 3.9)) // 2009/11/10 23:00:00 INFO 6.1

	type myint int
	type myfloat float64

	//ERROR
	//Print(AddTC(myint(2), myint(3)))     //Cannot use myint as the type interface{ int | float64 } Type does not implement constraint interface{ int | float64 } because type is not included in type set (int, float64)
	//Print(AddTC(myfloat(2.2), myfloat(3.9))) // Cannot use myfloat as the type interface{ int | float64 } Type does not implement constraint interface{ int | float64 } because type is not included in type set (int, float64)
	//Print(AddTSI(myint(2), myint(3)))
	//Print(AddTSI(myfloat(2.2), myfloat(3.9)))

	Print(AddTSI(2, 3))     // 2009/11/10 23:00:00 INFO 5
	Print(AddTSI(2.2, 3.9)) // 2009/11/10 23:00:00 INFO 6.1

	Print(AddUTSI(myint(2), myint(3)))         //2009/11/10 23:00:00 INFO 5
	Print(AddUTSI(myfloat(2.2), myfloat(3.9))) //2009/11/10 23:00:00 INFO 6.1

}
