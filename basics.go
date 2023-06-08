package main

import "fmt"

func main() {

	// 1 - Hola mundo
	/*
		message := "Hello world whit Golang"
		fmt.Print(message)
	*/

	// 2 - Variables y tipos de datos
	/*
		var edad int = 19
		var nombre string = "Agustin"
		var altura float32 = 1.85

		fmt.Println("Hola soy", nombre, "tengo ", edad, "y mido", altura)

		var pi float32 = 3.1415
		var verdad bool = true

		fmt.Println("Es el numero pi", pi, "aproximadamente?", verdad)

		var caracter rune = 'a' // los caracteres Unicode se representan internamente como números

		fmt.Println(nombre, "empieza con", string(caracter))
	*/

	// 3 - Funciones
	fmt.Println(sumar(10, 20))
}
func sumar(a, b int) int {
	return a + b
}
