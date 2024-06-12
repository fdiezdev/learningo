/*
Módulo 4: estructuras de datos
--------------------------------------------------------------------------
ARRAYS:
! Todos los valores en un array tienen el mismo tipo de dato
! Los arrays tienen una tamaño fijo

Para seleccionar un rango de valores dentro de un array:
array[inicio:final_excluyente]
*/
package main

import "fmt"

func main()  {
	var names [5] string
	names[0] = "Marge"
	names[1] = "Bart"
	names[2] = "Homero"
	names[3] = "Lisa"
	names[4] = "Maggie"

	ages := [5] uint8 {45, 10, 47, 8, 2}

	fmt.Println("Tu nombre es: ", names[0], " y tu edad es ", ages[0])

	marge := names[0]
	marge_age := ages[0]

	fmt.Println("Tu nombre es: ", marge, " y tu edad es ", marge_age)

	// para ver el tamaño de un array usamos la funcion len()
	size := len(names)
	fmt.Println("Tamaño del array: ", size)

	// parte de un array
	daughters := names[3:5] // como el final es excluyente debo agregar un indice mas para que me incluya el último que yo quiero
	fmt.Println("Las hijas de los simpson son: ", daughters)
}