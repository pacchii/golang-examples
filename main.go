package main

import (
	"fmt"

	"github.com/pacchii/golang-examples/slice"
)

func main() {

	//array.ArrayExample()
	slice.SliceExample()

	//fmt.Println(nextslicecap(513, 512))

	/*
		x := 10
		y := &x
		fmt.Println(y)
		fmt.Println(x)
		fmt.Println(y)
	*/
}

func nextslicecap(newLen, oldCap int) int {
	newcap := oldCap
	doublecap := newcap + newcap
	if newLen > doublecap {
		return newLen
	}

	const threshold = 256
	if oldCap < threshold {
		return doublecap
	}
	for {
		// Transition from growing 2x for small slices
		// to growing 1.25x for large slices. This formula
		// gives a smooth-ish transition between the two.
		fmt.Println("Before", newcap)
		newcap += (newcap + 3*threshold) >> 2
		fmt.Println("After", newcap)

		// We need to check `newcap >= newLen` and whether `newcap` overflowed.
		// newLen is guaranteed to be larger than zero, hence
		// when newcap overflows then `uint(newcap) > uint(newLen)`.
		// This allows to check for both with the same comparison.
		if uint(newcap) >= uint(newLen) {
			break
		}
	}

	// Set newcap to the requested cap when
	// the newcap calculation overflowed.
	if newcap <= 0 {
		return newLen
	}
	return newcap
}
