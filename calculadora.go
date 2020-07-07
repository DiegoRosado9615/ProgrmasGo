package main

import "fmt"

/**
Función que suma 2 números y regresa la suma de ambos números
@param num 1 int
*/
func suma(num1 int, num2 int) int {
	return num1 + num2
}

/**
Función que multiplica 2 numero, regresa
*/
func multi(num1 int, num2 int) int {
	return num1 * num2
}

/**
Funcion que divide 2 numeros y devuelve un numero
*/
func divide(num1 float32, num2 float32) float32 {
	return num1 / num2
}

func sumaComplejos(comple1 complex64, comple2 complex64) complex64 {
	return comple1 + comple2
}

// Main
func main() {
	//var x int = 2

	var input int
	fmt.Scanln(&input)
	fmt.Println("Calculadora")
	fmt.Println("suma  ........ 1")
	fmt.Println("resta ........ 2")
	if input == 1 {
		fmt.Println("caso 1")
	} else {
		fmt.Println("Caso2")
	}

}
