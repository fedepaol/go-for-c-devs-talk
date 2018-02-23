package main

import "fmt"

func main() {
	source := make(chan int)
	incremented := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			source <- i
		}
		close(source)
	}()
	go func() {
		for s := range source {
			incremented <- s + 1000
		}
		close(incremented)
	}()
	for res := range incremented {
		fmt.Println(res)
	}
}
