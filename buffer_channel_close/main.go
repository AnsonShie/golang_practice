package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	ch <- 2
	ch <- 3
	ch <- 5
	v, ok := <-ch
	fmt.Println(v)
	fmt.Println(ok)
	v, ok = <-ch
	fmt.Println(v)
	fmt.Println(ok)
	v, ok = <-ch
	fmt.Println(v)
	fmt.Println(ok)

	close(ch)
	fmt.Println("Channel closed.")

	for i := 1; i <= 10; i++ {
		v, ok := <-ch
		fmt.Println(v)
		fmt.Println(ok)
	}
	fmt.Println("Channel closed!")
}
