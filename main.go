package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)

	go doSomething(c)

	<-c
	fmt.Println("Doing something else")

}

func doSomething(c chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Done doing something")

	c <- true
}
