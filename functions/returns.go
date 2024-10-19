package main

import "fmt"

func sum(values ...int) int {
	total := 0

	for _, num := range values {
		total += num
	}

	return total
}

// func printNames(names ...string) {
// 	for _, name := range names {
// 		fmt.Println(name)
// 	}
// }

func getValues(x int) (double int, y int, z int) {
	double = 2 * x
	y = 3 * x
	z = 4 * x

	return
}

func main() {
	fmt.Println(sum(1, 2, 3))
}
