package _9_pointers

import "fmt"

type cat struct {
	name string
}

func (d cat) run() {
	fmt.Println(d.name, "is running")
}

func (d *cat) walk() {
	fmt.Println(d.name, "is walking")
}

type animal interface {
	run()
	walk()
}

func AnimalWalk(a animal) {
	a.walk()
}

func MethodSets2Example() {
	d1 := cat{"cat1"}
	d1.run()
	d1.walk()
	//AnimalWalk(d1) //Cannot use d1 (type cat) as the type animal. Type does not implement animal as the walk method has a pointer receiver

	d2 := &cat{"cat2"}
	d2.run()
	d2.walk()
	AnimalWalk(d2)

}
