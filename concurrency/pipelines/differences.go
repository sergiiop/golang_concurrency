// package main

// import "fmt"

// // Version INCORRECTA - Sin goroutine
// func generatorIncorrecto() <-chan int {
// 	out := make(chan int)
// 	// Sin go func()
// 	for i := 1; i <= 5; i++ {
// 		out <- i // ¡ESTO BLOQUEARÁ!
// 	}
// 	close(out)
// 	return out
// }

// // Version CORRECTA - Con goroutine
// func generatorCorrecto() <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		for i := 1; i <= 5; i++ {
// 			out <- i // Esto funciona porque está en una goroutine
// 		}
// 		close(out)
// 	}()
// 	return out
// }

// func main() {
// 	ch := generatorIncorrecto()

// 	// Este código nunca se ejecutará porque generatorIncorrecto
// 	// se bloqueará en el primer intento de enviar al channel
// 	for n := range ch {
// 		fmt.Println(n)
// 	}
// }

// func main() {
// 	ch := generatorCorrecto()
// 	// Ahora sí funciona porque:
// 	// 1. La función retorna inmediatamente el channel
// 	// 2. La goroutine se ejecuta en segundo plano
// 	for n := range ch {
// 		fmt.Println(n)
// 	}
// }
