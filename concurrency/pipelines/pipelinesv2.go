// package main

// import (
// 	"fmt"
// )

// // generator: crea un channel que emite los números del 1 al 5
// func generator() <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		for i := 1; i <= 5; i++ {
// 			out <- i
// 		}
// 		close(out)
// 	}()
// 	return out
// }

// // square: recibe números de un channel, los eleva al cuadrado y los envía a otro channel
// func square(in <-chan int) <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		for n := range in {
// 			out <- n * n
// 		}
// 		close(out)
// 	}()
// 	return out
// }

// func main() {
// 	// Construir el pipeline
// 	c := generator() // Genera números
// 	out := square(c) // Eleva al cuadrado los números

// 	// Consumir el output
// 	for n := range out {
// 		fmt.Println(n)
// 	}
// }
