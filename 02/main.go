/*
Clase #2: variables
- tipo de dato
- nombre de variable

En go está la palabara reservada "var"
var nombre tipo_de_dato
--------------------------------------------------------------------------
! No puedo crear variables en Go y no usarlas

! Go puede detectar automáticamente el tipo de dato, y no me hace falta declarar el tipo de dato, usando nombre:=contenido

! En Go podemos crear variables juntas nombre1,nombre2 := "valor 1", 2
--------------------------------------------------------------------------
Podemos crear también constantes -> variables que no pueden alterar su valor:
const nombre = "valor constante"
--------------------------------------------------------------------------
# Tipos de datos
Tenemos que ser específicos con el tipo de dato que usamos en función del espacio que queremos usar en memoria
Ej. si vamos a almarcenar una edad, conviene almacenarla en uint8 en vez de int8 o int64

! Asignar el tipo de dato correcto nos ayuda a optimizar el uso de memoria
--------------------------------------------------------------------------
IMPORTANTE: no podemos realizar operaciones aritméticas entre diferentes tipos de datos!
Ej. var a int
	var b uint8
	c := a + b ==> ERROR

Para poder hacer esto usamos CASTING (cambio de una variable temporalmente para poder cumplir)
--------------------------------------------------------------------------
fmt.Printf() nos permite formatear la salida y concatenar de manera más simple:
$s -> significa que voy a pasar un string
%d -> significa que voy a pasar un número
%T -> me indica el tipo de dato de la variable que pase como parametro
--------------------------------------------------------------------------
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

	altura, genero := 175, "masculino"

	fmt.Println(nombre, apellido, edad, altura, genero)

	const lang string = "go"
	// lang = "hola" -> esto es ilegal porque no puedo cambiar el valor de una cte.
	fmt.Println("Estoy aprendiendo: ", lang)

	// almacenamos edad en uint8 -> unsigned int 8 bits
	var edadoptim uint8 = 22
	fmt.Println(edadoptim)

	// casting

	var a uint8
	var b int

	a = 245
	b = 7

	// c := a + b // esto da error porque tienen un tipo de dato diferente
	c := int(a) + b
	fmt.Println(c)

	// Printf
	fmt.Printf("Hola %s, como estás?", nombre)

	// tipo de variable
	fmt.Printf("\n La variable c es de tipo %T y la variable a es de tipo %T", c, a)
}
