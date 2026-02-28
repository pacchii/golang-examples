package _1_json

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Person struct {
	Name string
	Age  int
}

func JsonExample() {

	var jsonBlob = []byte(`[
	{"Name": "Prashant", "Age":34},
	{"Name": "Aishwarya", "Age":29}
	]`)

	var persons []Person

	json.Unmarshal(jsonBlob, &persons)
	fmt.Println(persons) //[{Prashant 34} {Aishwarya 29}]

	p1 := Person{Name: "Prashant", Age: 23}
	p2 := Person{Name: "Aishwarya", Age: 18}
	var ps []Person
	ps = append(ps, p1, p2)

	bs, err := json.Marshal(persons)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bs)) //[{"Name":"Prashant","Age":34},{"Name":"Aishwarya","Age":29}]

	fmt.Fprintln(os.Stdout, "Hello World")
	io.WriteString(os.Stdout, "Prashant")

}
