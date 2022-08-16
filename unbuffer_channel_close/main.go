package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		v, ok := <-ch
		fmt.Println(v)
		fmt.Println(ok)
	}()
	ch <- 2

	close(ch)
	fmt.Println("Channel closed.")

	for i := 1; i <= 10; i++ {
		v, ok := <-ch
		fmt.Println(v)
		fmt.Println(ok)
	}
	fmt.Println("Channel closed!")
}
