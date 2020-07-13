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
Función que me permite
*/

func creaAlumno(x int) Alumno {
	nombre := ""
	fmt.Println("Escriba su nombre")
	fmt.Scanln(&nombre)
	edad := ""
	fmt.Println("Escriba su edad")
	fmt.Scanln(&edad)
	alumno := Alumno{x, nombre, edad}
	return alumno
}

/**
Funcion auxiliar que me permite agregar un alumno a la lista de alumnos
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

/**
Funcion auxiliar que me permite assegurar que los id no se van a repetir
si se borra un alumno de la lista.
Recibe una lista de alumnos
Devuelve una lista de alumnos
*/

func validarID(lista []Alumno) []Alumno {
	listaAuxiliar := make([]Alumno, 0)
	centinela := Alumno{0, "1", "0"}
	x := 0
	for j := 0; j < len(lista); j++ {
		centinela = lista[j]
		x = j + 1
		centinela.id = x
		listaAuxiliar = append(listaAuxiliar, centinela)
	}
	return listaAuxiliar

}

/**
Funcion que me permite eliminar a los alumnos de la lista, recibe la lista
con los alumnos en el grupo y el indice  que me dice que alumno quiere
borrar
*/
func eliminaLista(lista []Alumno, i int) []Alumno {
	if hayAlumnos(lista) {
		fmt.Println("No hay alumnos")
		return lista
	} else if len(lista) == 1 {
		listaAlumnos := make([]Alumno, 0)
		return listaAlumnos
	}
	lista[i] = lista[len(lista)-1]
	lista = lista[:len(lista)-1]
	return validarID(lista)
}

func main() {
	var salir int = 0
	listaAlumnos := make([]Alumno, 0)
	centinela := Alumno{0, "", "0"}
	index := 0
	for ok := 0; salir != 5; ok++ {
		fmt.Println("1 ingresa alumo")

		fmt.Println("2 ver lista de alumnos")
		fmt.Println("3 elimina alumnos")
		fmt.Println("4 Busca alumnos")
		fmt.Println("5 Salir")
		fmt.Scanln(&salir)

		if salir == 1 {
			if hayAlumnos(listaAlumnos) {
				centinela = creaAlumno(1)
				listaAlumnos = append(listaAlumnos, centinela)
			} else {
				centinela = creaAlumno(len(listaAlumnos) + 1)
				listaAlumnos = append(listaAlumnos, centinela)

			}

		}

		if salir == 2 {
			imprimeAlumnos(listaAlumnos)
		}

		if salir == 3 {
			fmt.Println("Escriba el numero de el alumno que quiere eliminar")
			fmt.Scanln(&index)
			if len(listaAlumnos) < index || index < 1 {
				fmt.Println("Ese Id no existe")

			} else {
				fmt.Println("Se elimino con exito")
				eliminaLista(listaAlumnos, index-1)
			}

		}
		if salir == 4 {
			fmt.Println("Escriba el numero del alumno que quiere buscar")
			fmt.Scanln(&index)
			if len(listaAlumnos) < index || index < 1 {
				fmt.Println("Ese Id no existe")

			} else {
				fmt.Println("El alumno encontrado es")
				listaAlumnos[index-1].imprime()
			}

		}

	}
}
