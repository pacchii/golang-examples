package _7_methods

import "fmt"

type person struct {
	fName string
	lName string
	age   int
}

func (p person) print() {
	fmt.Printf("%s %s %d\n", p.fName, p.lName, p.age)
}

func (p person) changeAge(newAge int) { //The value of age will change onl in this method scope it will not get changed in called func
	p.age = newAge
}
func (p *person) changeAgePointer(newAge int) { // The new age value is changed in memory location so caller object will get new age value
	p.age = newAge
}

func MethodExample() {
	p1 := person{"Prashanr", "Ghiwari", 34}
	p2 := person{"Aishwarya", "Ghiwari", 29}

	p1.changeAge(32)
	p2.changeAgePointer(27)

	p1.print()
	p2.print()

}
