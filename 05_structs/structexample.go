package _5_structs

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
	Address   Address // if We give variable name, then we cant access p.AddressLine1,
	//Address         // If we dont give variable name, then we can access address details using person streuct directly
}
type Address struct {
	AddrLine1 string
	AddrLine2 string
	Pincode   int
}

func StructExample() {

	a := Address{"#731, 60th Cross, 5th Block", "Near Bashyam Circle Rajajinagar Bangalore", 560010}
	p := Person{"Prashant", "Ghiwari", 34, a}

	fmt.Printf("%#v\n", p) //main.Person{FirstName:"Prashant", LastName:"Ghiwari", Age:34}

	//fmt.Println(p.AddrLine1, p.AddrLine2) // this works if you dont give variable name Address in Person struct
	AmbiguityExample()
	AnanymousStruct()

}

func AnanymousStruct() {

	p := struct {
		fName string
		lName string
	}{}

	fmt.Printf("%#v\n", p)
	p.fName = "Prashant"
	p.lName = "Ghiwari"
	fmt.Printf("%#v\n", p)

	p2 := struct {
		fName string
		lName string
	}{"Prashant", "Ghiwari"}

	fmt.Printf("%#v\n", p2)
	fmt.Printf("%#v\n", p2)
}

func AmbiguityExample() {
	type Result struct {
		marks int
		grade string
	}
	type Rank struct {
		marks int
		rank  int
	}
	type Student struct {
		Name string
		Result
		Rank
	}

	r := Result{444, "A"}
	rank := Rank{444, 15}

	s := Student{"Prashant", r, rank}
	fmt.Printf("%#v\n", s)
	//fmt.Println(s.marks) //compile error - Ambiguous reference 'marks'
	fmt.Println(s.Rank.marks) // this will work
}
