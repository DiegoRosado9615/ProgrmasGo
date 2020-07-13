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

func multComplejos(comple1 complex64, comple2 complex64) complex64 {
	return comple1 * comple2
}

// Main
func main() {
	fmt.Println("Calculadora")
	fmt.Println("suma  ........ 1")
	fmt.Println("resta ........ 2")
	fmt.Println("multiplicacion ..... 3")
	fmt.Println("division ........ 4")
	fmt.Println("Suma de complejo .... 5")
	fmt.Println("multiplicacion del complejo .... 6")
	var input int
	fmt.Scanln(&input)
	if input == 1 {
		fmt.Println("Escribe el  valor del numero 1")
		var num1 int
		fmt.Scanln(&num1)
		fmt.Println("Escribe el valor del  num 2")
		var num2 int
		fmt.Scanln(&num2)
		fmt.Println("numero1 + numero 2 =")
		fmt.Println(suma(num1, num2))
	} else if input == 2 {
		fmt.Println("Resta")
		fmt.Println("Escribe el  valor del numero 1")
		var num1 int
		fmt.Scanln(&num1)
		fmt.Println("Escribe el valor del  num 2")
		var num2 int
		fmt.Scanln(&num2)
		num2 = num2 * -1
		fmt.Println("numero1 - numero 2 = ")
		fmt.Println(suma(num1, num2))
	} else if input == 3 {
		fmt.Println("Escribe el  valor del numero 1")
		var num1 int
		fmt.Scanln(&num1)
		fmt.Println("Escribe el valor del  num 2")
		var num2 int
		fmt.Scanln(&num2)
		fmt.Println("numero1 * numero 2 =")
		fmt.Println(multi(num1, num2))
	} else if input == 4 {
		fmt.Println("Escribe el  valor del numero 1")
		var num1 float32
		fmt.Scanln(&num1)
		fmt.Println("Escribe el valor del  num 2")
		var num2 float32
		fmt.Scanln(&num2)
		fmt.Println("numero1 / numero 2 =")
		fmt.Println(divide(num1, num2))
	} else if input == 5 {
		fmt.Println("Escribe el  valor del numero 1 la parte real")
		var num1 float32
		fmt.Scanln(&num1)
		fmt.Println("Escribe el valor del  numero 1 la parte imaginaria")
		var num1i float32
		fmt.Scanln(&num1i)
		var complejo1 = complex(num1, num1i)
		fmt.Println("Escribe el  valor del numero 2 la parte real")
		var num2 float32
		fmt.Scanln(&num2)
		fmt.Println("Escribe el valor del  numero 2 la parte imaginaria")
		var num2i float32
		fmt.Scanln(&num2i)
		var complejo2 = complex(num2, num2i)
		fmt.Println("La suma de los 2 numero complejos son ")
		fmt.Println(sumaComplejos(complejo1, complejo2))
	} else if input == 6 {
		fmt.Println("Escribe el  valor del numero 1 la parte real")
		var num1 float32
		fmt.Scanln(&num1)
		fmt.Println("Escribe el valor del  numero 1 la parte imaginaria")
		var num1i float32
		fmt.Scanln(&num1i)
		var complejo1 = complex(num1, num1i)
		fmt.Println("Escribe el  valor del numero 2 la parte real")
		var num2 float32
		fmt.Scanln(&num2)
		fmt.Println("Escribe el valor del  numero 2 la parte imaginaria")
		var num2i float32
		fmt.Scanln(&num2i)
		var complejo2 = complex(num2, num2i)
		fmt.Println("La multiplicacion de los 2 numero complejos son ")
		fmt.Println(multComplejos(complejo1, complejo2))
	}

}
