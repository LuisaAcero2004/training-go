package main

import "fmt"

func main() {

	//Operadores aritméticos

	x := 10
	y := 50

	//Suma
	result := x + y
	fmt.Println("Suma:", result)

	//Resta
	result = x - y
	fmt.Println("Resta:", result)

	//Multiplicación
	result = x * y
	fmt.Println("Multiplicación:", result)

	//División
	result = x / y
	fmt.Println("División:", result)

	//Módulo
	result = x % y
	fmt.Println("Módulo:", result)

	//Incremental
	x++
	fmt.Println("Incremental:", result)

	//Decremental
	x--
	fmt.Println("Decremental:", result)

	//Retos
	//Área rectángulo, trapecio y círculo

	//Área rectángulo
	b := 8
	h := 4
	areaRectangulo := b * h
	fmt.Println("Área Rectángulo:", areaRectangulo)

	//Área trapecio
	b1 := 3
	b2 := 6
	h = 4
	areaTrapecio := (b1 * b2) * h / 2
	fmt.Println("Área Trapecio:", areaTrapecio)

	//Área círculo
	const pi float64 = 3.1416
	r := 5.0
	areaCirculo := pi * r * r
	fmt.Println("Área Circulo:", areaCirculo)

}
