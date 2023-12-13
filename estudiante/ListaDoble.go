package estudiante

import (
	"fmt"
	"strconv"
)

type ListaDoble struct {
	Inicio   *NodolistaDoble
	Contador int
}

// Agregar Estudiante a la Lista
func (l *ListaDoble) Agregar(carnet int, nombre string) {
	nuevoEstudiante := &Estudiante{Carnet: carnet, Nombre: nombre}
	nuevoNodo := &NodolistaDoble{Estudiante: nuevoEstudiante, Siguiente: nil, Anterior: nil}

	if l.Contador == 0 {
		l.Inicio = nuevoNodo
		l.Contador += 1
	} else {
		aux := l.Inicio
		for aux.Siguiente != nil {
			aux = aux.Siguiente
		}
		nuevoNodo.Anterior = aux
		aux.Siguiente = nuevoNodo
		l.Contador++
	}
}

// Funcion para encontrar los datos de inicio de sesion del estudiante
func (l *ListaDoble) Buscar(carnet string, password string) bool {
	if l.Contador == 0 {
		return false
	} else {
		aux := l.Inicio
		for aux != nil {
			if strconv.Itoa(aux.Estudiante.Carnet) == carnet && carnet == password {
				return true
			}
			aux = aux.Siguiente
		}
	}
	return false
}

// Funcion para poder Agregar Ordenadamente los estudiantes por medio de su carnet
func (l *ListaDoble) AgregarOrdenado(carnet int, nombre string) {
	nuevoEstudiante := &Estudiante{Carnet: carnet, Nombre: nombre}
	nuevoNodo := &NodolistaDoble{Estudiante: nuevoEstudiante, Siguiente: nil, Anterior: nil}

	if l.Contador == 0 {
		l.Inicio = nuevoNodo
		l.Inicio.Anterior = nuevoNodo
		l.Inicio.Siguiente = nuevoNodo
		l.Contador += 1
	} else {
		aux := l.Inicio
		contador := 1
		for contador < l.Contador {
			if l.Inicio.Estudiante.Carnet > carnet {
				nuevoNodo.Siguiente = l.Inicio
				nuevoNodo.Anterior = l.Inicio.Anterior
				l.Inicio.Anterior = nuevoNodo
				l.Inicio = nuevoNodo
				l.Contador++
				return
			}

			if aux.Estudiante.Carnet < carnet {
				aux = aux.Siguiente
			} else {
				nuevoNodo.Anterior = aux.Anterior
				aux.Anterior.Siguiente = nuevoNodo
				nuevoNodo.Siguiente = aux
				aux.Anterior = nuevoNodo
				l.Contador += 1
				return
			}
			contador += 1
		}
		if aux.Estudiante.Carnet > carnet {
			nuevoNodo.Siguiente = aux
			nuevoNodo.Anterior = aux.Anterior
			aux.Anterior.Siguiente = nuevoNodo
			aux.Anterior = nuevoNodo
			l.Contador += 1
			return
		}
		nuevoNodo.Anterior = aux
		nuevoNodo.Siguiente = l.Inicio
		aux.Siguiente = nuevoNodo
		l.Contador += 1
	}
}

// Funcion para poder imprimir la informacion de los alumnos
func (l *ListaDoble) Imprimir() {
	aux := l.Inicio
	contador := 1
	fmt.Println("Total Estudiantes: ", l.Contador)
	fmt.Println("-----------------------------------------------")
	for contador <= l.Contador {
		fmt.Println(aux.Estudiante.Carnet, "=>", aux.Estudiante.Nombre)
		aux = aux.Siguiente
		contador += 1
	}
	fmt.Println("-----------------------------------------------")
	fmt.Println("")
}

// Funcion para poder graficar
func (l *ListaDoble) Graficar() {
	nombreArchivo := "./Reportes/listadoble.dot"
	nombreImagen := "./Reportes/listadoble.jpg"
	texto := "digraph lista{\n"
	texto += "rankdir=LR;\n"
	texto += "node[shape = record];\n"
	texto += "nodonull1[label=\"null\"];\n"
	texto += "nodonull2[label=\"null\"];\n"
	aux := l.Inicio
	contador := 0
	texto += "nodonull1->nodo0 [dir=back];\n"
	for i := 0; i < l.Contador; i++ {
		texto += "nodo" + strconv.Itoa(i) + "[label=\"" + "Nombre: " + aux.Estudiante.Nombre + "\\n" + "Carnet: " + strconv.Itoa(aux.Estudiante.Carnet) + "\"];\n"
		aux = aux.Siguiente
	}
	for i := 0; i < l.Contador-1; i++ {
		c := i + 1
		texto += "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(c) + ";\n"
		texto += "nodo" + strconv.Itoa(c) + "->nodo" + strconv.Itoa(i) + ";\n"
		contador = c
	}
	texto += "nodo" + strconv.Itoa(contador) + "->nodonull2;\n"
	texto += "}"
	crearArchivo(nombreArchivo)
	escribirArchivo(texto, nombreArchivo)
	ejecutar(nombreImagen, nombreArchivo)
}
