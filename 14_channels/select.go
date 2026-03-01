package _4_channels

import "fmt"

func SelectExample() {

	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(even, odd, quit)

	receive(even, odd, quit)

	fmt.Println("Done!")

}

func receive(e, o, q <-chan int) {
	for {

		select {

		case v, ok := <-e:
			if ok {
				fmt.Println("EVE-", v)
			}
		case v, ok := <-o:
			if ok {
				fmt.Println("ODD-", v)
			}
		case <-q:
			fmt.Println("Quit")
			return

		}
	}

}

func send(e, o, q chan<- int) {
	for i := 0; i < 25; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
	q <- 0
	close(q)
}
