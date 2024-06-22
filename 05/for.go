/*
Módulo 5: ciclos
--------------------------------------------------------------------------
! Break nos permite detener la ejecución del ciclo

! Continue nos permite saltar una iteración del ciclo cuando se cumple una condición
*/

package main

import "fmt"

func main() {
	
	// Ciclo for clásico

	/*
	for i := 0; i < 5; i++ {
		if i == 4 {
			break
		}
		fmt.Println("Valor de i:",i)
	}
	
	for i := 4; i >= 0; i-- {
		if i==3 {
			continue 
		}
		fmt.Println("Valor de i:",i)
	}
	*/

	// Matrices

	matrix := [3][3]int{}

	// vamos a llenar la matriz con un ciclo for
	valor := 0 // este es el valor que vamos a ir incrementando y guardando en la matriz

	// Recorremos la matriz
	// Primero filas
	for i := 0; i <= 2; i++ {
		// Después columnas
		for j := 0; j <= 2; j++ {
			matrix[i][j] = valor
			valor++
		}
	}

	// Mostramos la matriz
	// Primero filas
	for i := 0; i <= 2; i++ {
		// Después columnas
		for j := 0; j <= 2; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Printf("\n")
	}

	
	// O mostramos el array directamente
	fmt.Println(matrix)
}