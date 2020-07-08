/**
Programa de go que fue creado para probar que tal eran sus "Objetos en este lenguaje"
y que me sirve de practicar en este lenguaje
*/
package main

import (
	"fmt"
	"math"
)

type Puntos struct {
	x int
	y int
}

func (p Puntos) imprime() {
	fmt.Println("El punto x = ", p.x, " El punto y = ", p.y)
}

/**
Funcion que dado un Punto nos dice si esta en un cuadrante u en otro cuadrante
*/

func (p Puntos) cuadrante() int {
	if (p.x > 0) && (p.y > 0) {
		fmt.Println("su punto esta en el cuadrante 1")
		return 1
	}
	if (p.x < 0) && (p.y > 0) {
		fmt.Println("su punto esta en el cuadrante 2")
		return 2
	}
	if (p.x < 0) && (p.y < 0) {
		fmt.Println("su punto esta en el cuadrante 3")
		return 3
	}
	if (p.x > 0) && (p.y < 0) {
		fmt.Println("su punto esta en el cuadrante 4")
		return 4
	}
	if (p.x == 0) && (p.y == 0) {
		fmt.Println("su punto esta en el origen")
		return 0
	}

	if (p.x >= 0) && (p.y == 0) {
		fmt.Println("su punto esta en el eje de las y")
		return 0
	}
	if (p.x == 0) && (p.y >= 0) {
		fmt.Println("su punto esta en el eje de las y")
		return 0
	}
	return 0
}

/**
Funcion que saca la distancia de 2 puntos
*/
func (p Puntos) distancia(p2 Puntos) float64 {
	var x int = p2.x - p.x
	var y int = p2.y - p.y
	x = x * x
	y = y * y
	var suma int = x + y
	var fin float64 = float64(suma)
	return math.Sqrt(fin)
}

/**
Funcion que  saca la suma de 2 puntos y te devuelve un punto
*/
func (p Puntos) suma(p2 Puntos) Puntos {
	puntoNuevo := Puntos{p.x + p2.x, p.y + p2.y}
	return puntoNuevo
}

func main() {
	p1 := Puntos{7, 5}
	p2 := Puntos{4, 1}
	p1.imprime()
	p1.cuadrante()
	fmt.Println(p1.distancia(p2))
	p3 := p1.suma(p2)
	fmt.Println("Esta es la suma de p1 con p2")
	p3.imprime()

}
