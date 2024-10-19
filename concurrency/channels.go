package main

import "fmt"

func main2() {

	// Unbuffered channel
	c := make(chan int)

	c <- 1

	fmt.Println(<-c)
}
