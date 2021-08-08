package main

import "fmt"

func main() {

	//Declaración de constantes

	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	//Declaración de variables enteras

	base := 12          //los dos puntos sirven para indicar que la variable no ha sido declarada antes
	var altura int = 14 //Declaración de tipo de dato y se le asigna un valor
	var area int        //Declaración tipo de dato

	fmt.Println(base, altura, area)

	//Zero values
	var a int     //0
	var b float64 //0
	var c string  //""
	var d bool    //false

	fmt.Println(a, b, c, d)

	//Area de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area Cuadrado:", areaCuadrado)

}
