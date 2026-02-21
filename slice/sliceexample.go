package slice

import "fmt"

func SliceExample() {

	fmt.Println("-----------------SliceExample--------------")

	//simple slice

	ss := []int{1, 2, 3, 4, 5}

	fmt.Printf("%T - %v\n", ss, ss) //[]int - [1 2 3 4 5]

	for i, v := range ss {
		fmt.Println(i, v)
	}

	// _ is a blank identifier
	for _, v := range ss {
		fmt.Println(v)
	}

	fmt.Printf("Getting slice values using simple for loop with index")
	for i := 0; i < len(ss); i++ {
		fmt.Println(i, ss[i])
	}

	//Slice with make
	s := make([]int, 2, 5)

	sl := make([]int, 3)

	fmt.Println(s, len(s), cap(s))    //[0 0] 2 5
	fmt.Println(sl, len(sl), cap(sl)) //[0 0 0] 3 3

	//Append value to slice

	newsl := []int{}
	fmt.Println("Initial slice ", newsl, len(newsl), cap(newsl)) //[] 0 0
	newsl = append(newsl, 1)
	fmt.Println("Append 1 ", newsl, len(newsl), cap(newsl)) //Append 1  [1] 1 1

	newsl = append(newsl, 2)
	fmt.Println("Append 2 ", newsl, len(newsl), cap(newsl)) //Append 2  [1 2] 2 2

	newsl = append(newsl, 3)
	fmt.Println("Append 3 ", newsl, len(newsl), cap(newsl)) //Append 3  [1 2 3] 3 4

	newsl = append(newsl, 4)
	fmt.Println("Append 4 ", newsl, len(newsl), cap(newsl)) //Append 4  [1 2 3 4] 4 4

	newsl = append(newsl, 5)
	fmt.Println("Append 5 ", newsl, len(newsl), cap(newsl)) //Append 5  [1 2 3 4 5] 5 8

	newsl = append(newsl, 6)
	fmt.Println("Append 6 ", newsl, len(newsl), cap(newsl)) //Append 6  [1 2 3 4 5 6] 6 8

	newsl = append(newsl, 7)
	fmt.Println("Append 7 ", newsl, len(newsl), cap(newsl))

	newsl = append(newsl, 8)
	fmt.Println("Append 8 ", newsl, len(newsl), cap(newsl))

	newsl = append(newsl, 9)
	fmt.Println("Append 9 ", newsl, len(newsl), cap(newsl))

	//for i := 10; i < 514; i++ {
	//newsl = append(newsl, i)
	//fmt.Println("Append ", i, len(newsl), cap(newsl))

	//}

	fmt.Println(len(newsl), cap(newsl))

	newslice := make([]int, 468)
	fmt.Println("---------------------random 468 size checking capacity")
	fmt.Println(len(newslice), cap(newslice))
	newslice = append(newslice, 11)
	fmt.Println(len(newslice), cap(newslice))

	nslice1 := make([]int, 512)

	fmt.Println("---------------------random 512 size checking capacity")
	fmt.Println("Before", len(nslice1), cap(nslice1))
	nslice1 = append(nslice1, 99)
	fmt.Println("After", len(nslice1), cap(nslice1))

	nsl1 := []int{1, 2, 3}
	fmt.Println(len(nsl1), cap(nsl1))
	t := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	nsl1 = append(nsl1, t...)
	fmt.Println(len(nsl1), cap(nsl1))

	fmt.Println("---------------------Slicing slice------------------")

	ss = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(ss, len(ss), cap(ss)) //[1 2 3 4 5 6 7 8 9] 9 9

	fmt.Println(ss[:], len(ss), cap(ss))   //[1 2 3 4 5 6 7 8 9] 9 9
	fmt.Println(ss[4:], len(ss), cap(ss))  //[5 6 7 8 9] 9 9
	fmt.Println(ss[:5], len(ss), cap(ss))  //[1 2 3 4 5] 9 9
	fmt.Println(ss[2:5], len(ss), cap(ss)) //[3 4 5] 9 9

	//ss = ss[2:5]
	//fmt.Println(ss, len(ss), cap(ss)) //[3 4 5] 3 7

	//ss = ss[:5]
	//fmt.Println(ss, len(ss), cap(ss)) //[1 2 3 4 5] 5 9

	//ss = ss[7:]
	//fmt.Println(ss, len(ss), cap(ss)) //[8 9] 2 2

	//ss = ss[:]
	//fmt.Println(ss, len(ss), cap(ss)) //[1 2 3 4 5 6 7 8 9] 9 9

	//Deleting from slice
	fmt.Println("----------------deleting from slice---------------------")
	ss = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(ss, len(ss), cap(ss))
	//I want to delete 5
	ss1 := ss[:4]
	fmt.Println(ss1, len(ss1), cap(ss1)) //[1 2 3 4] 4 9
	ss2 := ss[5:]
	fmt.Println(ss2, len(ss2), cap(ss2)) //[6 7 8 9] 4 4

	ss = append(ss[:4], ss[5:]...)
	fmt.Println(ss, len(ss), cap(ss))

	//Test
	fmt.Println("-------------------testing different size slice with append------------")

	st1 := make([]int, 4, 10)
	fmt.Println(st1, len(st1), cap(st1)) //[0 0 0 0] 4 10
	st1[0] = 1
	st1[1] = 2
	st1[2] = 3
	st1[3] = 4

	fmt.Println(st1, len(st1), cap(st1)) //[1 2 3 4] 4 10
	st2 := []int{5, 6, 7, 8}
	st1 = append(st1, st2...)
	fmt.Println(st1, len(st1), cap(st1)) //[1 2 3 4 5 6 7 8] 8 10
	fmt.Println(st2, len(st2), cap(st2)) //[5 6 7 8] 4 4

	//---------------------------- multidimentional slice -------------------

	fmt.Println("------------multidimnetional slice----------------")

	mds := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(mds, len(mds), cap(mds)) //[[1 2 3] [4 5 6] [7 8 9]] 3 3

	/*

		type slice struct {
			array unsafe.Pointer
		    len int
		    cap int
		}

	*/

	//Underlying array of slice

	fmt.Println("------------ Underlying array in slice-------------")
	//Here even if i changed in a, since a and b shares same backed array, changing with 1 slice will make changes for other slice too
	a := []int{1, 2, 3, 4}
	fmt.Println(a, len(a), cap(a)) //[1 2 3 4] 4 4
	b := a
	fmt.Println(b, len(b), cap(b)) //[1 2 3 4] 4 4
	a[2] = 6
	fmt.Println(a, len(a), cap(a)) //[1 2 6 4] 4 4
	fmt.Println(b, len(b), cap(b)) //[1 2 6 4] 4 4

}
