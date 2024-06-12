/*
Módulo 4: estructuras de datos
--------------------------------------------------------------------------
SLICES:
- Apuntan a un array
- Tienen tamaño
- Tienen una capacidad
> Nos permiten aumentar o disminuir la cantidad de elementos
> Por defecto tienen el doble de capacidad que les damos
Ejemplo: 
Creo un slice de capacidad 3 -> go le da capacidad 6
Cuando lo lleno a los 6 lugares -> go aumenta a 7 y le da el doble de capacidad => 14 lugares

! Esto le permite al programa evitar agregar memoria en tiempo de ejecución

La función cap(nombre del slice) me dice la capacidad de ese slice
--------------------------------------------------------------------------
! La forma más convencional de crear slices es usando la función make()
nombreSlice := make(tipo de dato del slice, tamaño inicial, capacidad inicial opcional)
*/

package main

import "fmt"

func main(){
	var names []string
	// función append(que slice le quiero agregar info, que info le quiero agregar)
	names = append(names, "Homero")
	names = append(names, "Bart")
	names = append(names, "Marge")
	names = append(names, "Lisa")
	names = append(names, "Maggie")
	fmt.Println(names)
	fmt.Printf("Tamaño del slice: %d y su capacidad es: %d \n", len(names), cap(names))

	// ! ¿Cuál es la diferencia con el array tradicional? 
	// Yo NO definí el slice con un tamaño fijo
	// Le puedo seguir agregando nombres sin aumentar la capacidad
	// Go lo hace solo:
	names = append(names, "Moe")
	names = append(names, "Barney")
	names = append(names, "Lenny")
	names = append(names, "Carl")
	
	fmt.Println(names)
	fmt.Printf("Tamaño del slice: %d y su capacidad es: %d \n", len(names), cap(names))

	// Podemos editar los valores de un slice como si fueran un array normal:
	names[6] = "Nelson"

	fmt.Println(names)

	// Usando la función make:
	pets := make([]string, 0)
	pets = append(pets, "Ayudante de santa")
	pets = append(pets, "Bola de nieve")

	fmt.Println(pets)
	fmt.Printf("Tamaño del slice: %d y su capacidad es: %d \n", len(pets), cap(pets))

	// También podemos crear slices de la siguiente manera
	drunks := []string {
		"Barney",
		"Homero",
		"Lenny",
		"Carl",
	} // si lo hacemos en una sola línea no hacen falta las comillas

	fmt.Println(drunks)
	fmt.Printf("Tamaño del slice: %d y su capacidad es: %d \n", len(drunks), cap(drunks))
}