package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func inOrderTraversal(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	// Recorrer el subárbol izquierdo
	inOrderTraversal(t.Left, ch)

	// Enviar el valor del nodo actual al canal
	ch <- t.Value

	// Recorrer el subárbol derecho
	inOrderTraversal(t.Right, ch)
}

// Walk recorre el árbol t enviando todos los valores
// del árbol al canal ch.
func Walk(t *tree.Tree, ch chan int) {
	inOrderTraversal(t, ch)
	close(ch)
}

// Same determina si los árboles
// t1 y t2 contienen los mismos valores.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if ok1 != ok2 || v1 != v2 {
			return false
		}

		if !ok1 {
			break
		}

	}

	return true

}

func main() {

	// secondTree := tree.New(1)

	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Println(<-ch)
	// 	}
	// }()

	// Walk(firstTree, ch)

	fmt.Println(Same(tree.New(1), tree.New(2)))
}
