/*
Módulo 4: estructuras de datos
--------------------------------------------------------------------------
MAPS:
Sería como los diccionarios en otros lenguajes o arrays asociativos
Tenemos llaves personalizadas en vez de índices secuenciales
Ejemplo

idiomas["es"] = "Español"
idiomas["en"] = "Inglés"

! La función make tmb nos permite crear mapas
nombreMapa := make(map[tipo de dato del indice]tipo de dato del valor del mapa)

! La función delete nos permite eliminar posiciones de un mapa
*/

package main

import "fmt"

func main(){
	lang := make(map[string]string)

	lang["es"] = "Spanish"
	lang["en"] = "English"
	lang["fr"] = "French"
	lang["pt"] = "Portuguese"
	lang["ru"] = "Russian"
	
	fmt.Println(lang)

	// Tambien puedo definir los mapas asi
	countries := map[string]string{
		"ar": "Argentina",
		"us": "United States", 
		"uk": "United Kingdom", 
		"br": "Brazil",
	}

	fmt.Println(countries)

	// Podemos mostrar un valor específico

	fmt.Println(countries["ar"])

	// Función delete
	delete(countries, "uk")
	fmt.Println(countries)

	// Queremos saber si dentro de un mapa hay un valor
	// con una clave
	// ejemplo quiero saber si existe el ruso en el mapa de idiomas
	// > fmt.Println(lang["ru"])
	// Me devuelve un valor vacío
	// Vamos a chequear con un if
	if lang["ru"] != "" {
		fmt.Println("Existe el ruso")
	} else {
		fmt.Println("No existe el ruso")
	}

	fmt.Println(lang["ru"])
}