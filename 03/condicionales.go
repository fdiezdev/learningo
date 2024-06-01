/*
--------------------------------------------------------------------------
CONDICIONALES IF
! En go podemos crear instrucciones dentro de las sentencias if
Esto nos permite crear variables en la misma sentencia if y asignarle valor al mismo tiempo que analizamos la variable
Las variables que creamos en la sentencia if estan definidas SOLO dentro del if, no las podemos usar por fuera de estos!
--------------------------------------------------------------------------
CONDICIONALES SWITCH
No es necesario usar break al final de cada case como en C++
Fallthrough usamos cuando quiero saltear un case
--------------------------------------------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world!")

	if 5 > 10 {
		fmt.Println("5 > 10")
	} else {
		fmt.Println("Falso")
	}

	if isValid := true; isValid {
		fmt.Println("Is valid es verdadero")
	} else {
		fmt.Println("Is valid NO es verdadero")
	}

	// fmt.Println(isValid) --> esto no es legal pq' solo está definido para el if anterior

	if edad, nombre := 22, "Francisco"; edad < 12 {
		fmt.Printf("%s es un niño", nombre)
	} else if edad < 18 {
		fmt.Printf("%s es un adolescente", nombre)
	} else if edad < 30 {
		fmt.Printf("%s es un adulto", nombre)
	} else {
		fmt.Printf("%s es un adulto grandecito", nombre)
	}

	// Switch
	var edad uint8 = 22
	switch edad {
	case 1:
		fmt.Println("Es un bebé")
	case 2:
		fmt.Println("Es un bebé todavía")
	default:
		fmt.Println("Ni idea")

	}
}
