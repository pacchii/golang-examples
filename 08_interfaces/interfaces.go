package _8_interfaces

import "fmt"

type person struct {
	firstname string
}

type secretAgent struct {
	person
	active bool
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func (p person) speak() {
	fmt.Println("I am ", p.firstname, "and i am, person")
}

func (sa secretAgent) speak() {
	fmt.Println("I am ", sa.firstname, "and i am secret agent and active ", sa.active)
}

func InterfaceExample() {

	sa := secretAgent{
		person: person{firstname: "Prashant G"},
		active: true,
	}

	p := person{"Prasahnt"}

	fmt.Println(p, sa) //{Prasahnt} {{Prashant Ghiwari} true}

	//sa.speak() //I am  Prashant G and i am secret agent and active  true

	//p.speak() //I am  Prasahnt and i am, person

	saySomething(sa)
	saySomething(p)

}
