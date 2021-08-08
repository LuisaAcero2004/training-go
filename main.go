package main

import "fmt"

func main() {

	helloVar := "Hola"
	worldVar := "Mundo"

	//Println imprime en consola y agrega salto de línea al final
	fmt.Println(helloVar, worldVar)
	fmt.Println(helloVar, worldVar)

	//Printf
	fmt.Printf("Mi primer programa fue un %s %s\n", helloVar, worldVar) // Usar %v cuando no se sabe el tipo de dato

	//Sprintf genera un string pero no lo imprime en consola
	message := fmt.Sprintf("Mi primer programa fue un %s %s\n", helloVar, worldVar)
	fmt.Printf(message)

	//Tipo de datos
	fmt.Printf("helloVar es %T", helloVar) //Imprime el tipo de dato que le corresponde

}
