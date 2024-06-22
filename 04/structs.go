/*
Módulo 4: estructuras de datos
--------------------------------------------------------------------------
STRUCTS:
Sería el equivalente a las clases de otros lenguajes de programación

! Para declarar una estructura debemos asignarla con la palabra reservada type

! Si el nombre de la variable declarada empieza en mayúscula, significa que vamos a poder usar esa estructura en otros paquetes al importar este

Podemos poner slices dentro de una estructura
*/

package main

import "fmt"

// Persona es una estructura para darle nombre, apellido y edad a un personaje dado.
type Persona struct {
	Nombre string
	Apellido string
	Edad uint8
	Emails []string
}

func main(){
	var personaje1 Persona

	personaje1.Nombre = "Homero"
	personaje1.Apellido = "Simpson"
	personaje1.Edad = 45
	personaje1.Emails = []string {
		"homero@plantanuclear.com", 
		"bar@moe.org", 
		"vicepresidentejunior@compumundohipermegared.com",
	}

	// Otra forma de declarar

	personaje2 := Persona{"Marge", "Bouvier", 43, []string{"margesimpson@gmail.com"}}

	fmt.Println(personaje1)
	fmt.Println(personaje2)
}