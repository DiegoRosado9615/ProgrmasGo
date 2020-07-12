/**
Progrma que me va permitir inscribir alumnos a un grupo, me va a permitir elimniarlos o
modificar su nombre esto esta pensado para poder probrar estructura de datos en Go XD
**/

package main

import (
	"fmt"
)

/**
Objeto alumnos contendra un ID , un Nombre y una edad XD
*/

type Alumno struct {
	id     int
	nombre string
	edad   string
}

func (a Alumno) imprime() {
	fmt.Println("El nombre del alumno ", a.nombre, " con el id ", a.id, "con una edad de ", a.edad)
}

/**
Funcion que nos dice que si la lista que nos pasa esta vacia o contiene algun elemento
*/
func hayAlumnos(lista []Alumno) bool {
	return len(lista) == 0
}

/**
Funcion que nos imprime  el número de alumnos que hay en total en nuestro grupo XD
*/

func imprimeAlumnos(lista []Alumno) {
	prueba := Alumno{0, "", "0"}
	for i := 0; i < len(lista); i++ {
		prueba = lista[i]
		prueba.imprime()
	}
}

/**
Funcion que me permite agregar un alumno a la lista de alumnos
*/

func agregaAlumno(lista []Alumno, a Alumno) []Alumno {
	lista = append(lista, a)
	return lista

}

/**
Funcion que nos dice si 2 alumnos son iguales :V
*/
func (a Alumno) sonIguales(a2 Alumno) bool {
	return a.id == a2.id && a.nombre == a2.nombre && a.edad == a2.edad
}

/**
Funcion que nos devuelve el indice donde esta nuestro alumno, en caso de no encotrarlo
devuelve -1
*/
func buscaAlumnos(lista []Alumno, a Alumno) int {
	centinela := Alumno{0, "", "0"}
	for i := 0; i < len(lista); i++ {
		centinela = lista[i]
		if a.sonIguales(centinela) {
			return i
		}
	}
	return -1
}

func eliminaLista(lista []Alumno, i int) []Alumno {
	if hayAlumnos(lista) {
		fmt.Println("No hay alumnos")
		return lista
	} else if len(lista) == 1 {
		listaAlumnos := make([]Alumno, 0)
		return listaAlumnos
	}
	lista[i] = lista[len(lista)-1]
	return lista[:len(lista)-1]
}

func main() {
	alumno1 := Alumno{1, "Paco", "15"}
	alumno2 := Alumno{1, "Pancho", "14"}
	alumno3 := Alumno{1, "julia", "16"}
	listaAlumnos := []Alumno{alumno1, alumno2}
	listaAlumnos = agregaAlumno(listaAlumnos, alumno3)
	//listaAlumnos2 := make([]Alumno, 0)
	imprimeAlumnos(listaAlumnos)
	print(alumno1.sonIguales(alumno2))
}
