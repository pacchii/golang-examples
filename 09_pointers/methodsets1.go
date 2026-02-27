package _9_pointers

import "fmt"

type dog struct {
	name string
}

func (d dog) run() {
	fmt.Println(d.name, "is running")
}

func (d *dog) walk() {
	d.name = "ALTERED"
	fmt.Println(d.name, "is walking")
}

func MethodSetExample() {

	d1 := dog{"Hutch"}
	d1.run()        //Hutch is running
	fmt.Println(d1) //{Hutch}
	d1.walk()       //ALTERED is walking
	fmt.Println(d1) //{ALTERED}

	d2 := &dog{"Charlie"}
	d2.run()        //Charlie is running
	fmt.Println(d2) //&{Charlie}
	d2.walk()       //ALTERED is walking
	fmt.Println(d2) //&{ALTERED}
}
