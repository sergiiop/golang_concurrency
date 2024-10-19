package main

// # Pipeline Estilo 2 en Go

// ## Características Principales
// 1. Cada función crea y retorna su propio channel
// 2. Usa channels unidireccionales (<-chan)
// 3. Implementa goroutines internamente
// 4. Cada etapa es responsable de cerrar su propio channel

// ## Estructura Básica

// ## Ventajas
// 1. Mayor encapsulamiento (cada función maneja sus channels)
// 2. Más seguro (channels unidireccionales)
// 3. Fácil de componer (encadenamiento simple)
// 4. Responsabilidades claras (cierre de channels)
// 5. Menos acoplamiento entre componentes

// ## Puntos Clave
// 1. Siempre usa goroutines dentro de las funciones
// 2. Cierra los channels usando defer
// 3. Usa channels unidireccionales (<-chan) para mayor seguridad
// 4. Retorna inmediatamente el channel

import (
	"fmt"
)

func generator() <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 1; i <= 3; i++ {
			// fmt.Printf("Generando: %d\n", i)
			out <- i
		}
	}()
	return out
}

// in channel de lectura, retorna un chan de lectura
func double(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			// fmt.Printf("Duplicando: %d\n", n)
			out <- n * 2
		}
	}()
	return out
}

func main() {
	fmt.Println("Iniciando pipeline...")

	// Cada función retorna inmediatamente
	nums := generator()
	doubled := double(nums)

	// Consumimos los resultados
	for n := range doubled {
		fmt.Printf("Recibido: %d\n", n)
	}
}
