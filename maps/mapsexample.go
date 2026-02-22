package maps

import (
	"fmt"
	"os"
	"strings"
)

func MapsExample() {
	m := map[string]int{"Prashant": 34, "Aishwarya": 29}

	fmt.Println(m)

	m1 := map[string]int{}

	m1["Alpha"] = 1

	fmt.Println(m1)

	var m2 map[string]int

	m2 = m

	fmt.Println(m2)

	m2["Beeta"] = 2

	m3 := make(map[string]int)

	m3["beeta"] = 3

	fmt.Println(m3)

	//Iterating over for loop

	/*
		Below for loop will print below output. Because on line m2 = m, m and m2 referring to same map, if m2 altered m also will get changed
		vice versa
		Prashant 34
		Aishwarya 29
		Beeta 2
	*/

	for k, v := range m {
		fmt.Println(k, v)
	}

	for k, v := range m2 {
		fmt.Println(k, v)
	}

	delete(m, "Beeta")

	fmt.Println(m)  //map[Aishwarya:29 Prashant:34]
	fmt.Println(m2) //map[Aishwarya:29 Prashant:34]

	delete(m, "dummy") // this is not panic/error it will not throw error saying key not found

	fmt.Println(m["keynotfound"]) // 0  -- gets zero value of Value type here its int

	v, ok := m["keynotfound"]
	fmt.Println(v, ok) //0 false
	if ok {
		fmt.Println("Found")
	}
	fmt.Println("-----------------------CountWordsinSlice() --------------")
	CountWordsinSlice()

	fmt.Println("-----------------------CountWordsinFile() --------------")
	CountWordsinFile()
}

func CountWordsinSlice() {

	m := map[string]int{}

	words := []string{
		"apple", "banana", "apple", "orange", "banana",
		"apple", "grape", "orange", "banana", "grape",
		"kiwi", "apple", "kiwi", "banana", "mango",
		"mango", "apple", "orange", "grape", "banana",
	}

	for _, val := range words {
		v, ok := m[val]
		if ok {
			m[val] = v + 1
		} else {
			m[val] = 1
		}
	}

	fmt.Println(m)

}

func CountWordsinFile() {

	m := map[string]int{}

	data, err := os.ReadFile("data.txt")

	if err != nil {
		panic(err)
	}

	words := strings.Fields(string(data))

	for _, val := range words {
		v, ok := m[val]
		if ok {
			m[val] = v + 1
		} else {
			m[val] = 1
		}
	}

	fmt.Println(m)

}
