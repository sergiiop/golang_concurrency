// package main

// import "fmt"

// type Result struct {
// 	Value int
// 	Err   error
// }

// func GeneratorWithError() <-chan Result {
// 	out := make(chan Result)
// 	go func() {
// 		defer close(out)
// 		for i := 0; i <= 10; i++ {
// 			// Simular posible error
// 			if i == 5 {
// 				out <- Result{Err: fmt.Errorf("error en número %d", i)}
// 				return
// 			}
// 			out <- Result{Value: i}
// 		}
// 	}()
// 	return out
// }

// func DoubleWithError(in <-chan Result) <-chan Result {
// 	out := make(chan Result)
// 	go func() {
// 		defer close(out)
// 		for r := range in {
// 			if r.Err != nil {
// 				out <- r
// 				return
// 			}
// 			out <- Result{Value: r.Value * 2}
// 		}
// 	}()
// 	return out
// }

// func main() {
// 	numbers := GeneratorWithError()
// 	doubles := DoubleWithError(numbers)

// 	for result := range doubles {
// 		if result.Err != nil {
// 			fmt.Printf("Error: %v\n", result.Err)
// 			break
// 		}
// 		fmt.Println(result.Value)
// 	}
// }
