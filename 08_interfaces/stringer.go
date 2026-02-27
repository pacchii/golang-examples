package _8_interfaces

//go Playground https://go.dev/play/p/yljl83X7qvo

import (
	"fmt"
	"log"
	"strconv"
)

type Account struct {
	typ string
}

type Book struct {
	author string
}

type count int

func (c count) String() string {
	return fmt.Sprint("The number is", strconv.Itoa(int(c)))
}

func (a Account) String() string {
	return fmt.Sprint("Hey I am Account of type", a.typ)
}

func StringerExample() {
	a := Account{
		"SAVINGS",
	}
	b := Book{"S. L. Byrappa"}

	fmt.Println(a) //Hey I am Account of type SAVINGS

	fmt.Println(b) //{S. L. Byrappa}

	var c count = 25

	//fmt.Println(c) // The number is 25

	LogInfo(c) //2009/11/10 23:00:00 INFO  Hey I am number45

}

//Using log

func LogInfo(s fmt.Stringer) {
	log.Println("INFO ", s.String())
}
