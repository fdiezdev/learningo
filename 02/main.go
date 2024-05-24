/*
Clase #2: variables
- tipo de dato
- nombre de variable

En go está la palabara reservada "var"
var nombre tipo_de_dato

! No puedo crear variables en Go y no usarlas

! Go puede detectar automáticamente el tipo de dato, y no me hace falta declarar el tipo de dato, usando nombre:=contenido

! En Go podemos crear variables juntas nombre1,nombre2 := "valor 1", 2

*/

package main

import (
	"fmt"
)

func main() {
	var nombre, apellido string
	nombre = "Francisco"
	apellido = "Diez"

	edad := 22

	altura,genero := 175,"masculino"

	fmt.Println(nombre, apellido, edad, altura, genero)
}