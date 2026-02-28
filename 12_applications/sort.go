package _2_applications

//   https://go.dev/play/p/Mxd5M0-uUfs

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

//sort.Interface has below 3 methods, Len, Less, Swap so these user for sorting

func (p Person) String() string {
	return fmt.Sprintf("Person Name %s with age %d", p.Name, p.Age)
}

type ByAge []Person

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func SortExample() {

	s := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(s)

	sort.Ints(s)
	fmt.Println(s)

	CustomSort()

}

func CustomSort() {

	persons := []Person{{"Prashant", 34}, {"Aishwarya", 29}, {"Arshith", 4}, {"Shashikala", 56}}

	fmt.Println("Unsorted", persons)
	sort.Sort(ByAge(persons))

	fmt.Println("Sorted By Age", persons)

	//Also we can sort without implementing those 3 methods

	sort.Slice(persons, func(i, j int) bool {
		return persons[i].Name < persons[j].Name
	})

	fmt.Println("Sorted By Name", persons)

}
