package tutores

import (
	"fmt"
	"strconv"
)

type ListaDobleCircular struct {
	Inicio   *NodoListaCircular
	Longitud int
}

func (l *ListaDobleCircular) Agregar(carnet int, nombre string, curso string, nota int) {
	nuevoTutor := &Tutor{Carnet: carnet, Nombre: nombre, Curso: curso, Nota: nota}
	nuevoNodo := &NodoListaCircular{Tutor: nuevoTutor, Siguiente: nil, Anterior: nil}

	if l.Longitud == 0 {
		l.Inicio = nuevoNodo
		l.Inicio.Anterior = nuevoNodo
		l.Inicio.Siguiente = nuevoNodo
		l.Longitud++
	} else {
		aux := l.Inicio
		for i := 0; i < l.Longitud; i++ {
			if aux.Tutor.Curso == curso {
				if aux.Tutor.Nota < nota {
					aux.Tutor = nuevoTutor
				}
				return
			}
			aux = aux.Siguiente
		}
		nuevoNodo.Siguiente = l.Inicio
		nuevoNodo.Anterior = l.Inicio.Anterior
		l.Inicio.Anterior.Siguiente = nuevoNodo
		l.Inicio.Anterior = nuevoNodo
		l.Longitud++
	}
}

// Funcion para poder mostrar los tutores a los estudiantes
func (l *ListaDobleCircular) Mostrar() {
	aux := l.Inicio
	contador := 1
	for contador <= l.Longitud {
		fmt.Println("Curso:", aux.Tutor.Curso, " Nombre:", aux.Tutor.Nombre)
		aux = aux.Siguiente
		contador += 1
	}
}

// Funcion para poder graficar la lista de los tutores aceptados
func (l *ListaDobleCircular) Graficar() {
	nombreArchivo := "./Reportes/listadoblecircular.dot"
	nombreImagen := "./Reportes/listadoblecircular.jpg"
	texto := "digraph lista{\n"
	texto += "rankdir=LR;\n"
	texto += "node[shape = record];\n"
	aux := l.Inicio
	contador := 0
	for i := 0; i < l.Longitud; i++ {
		texto += "nodo" + strconv.Itoa(i) + "[label=\"" + "Nombre: " + aux.Tutor.Nombre + "\\n" + "Carnet: " + strconv.Itoa(aux.Tutor.Carnet) + "\"];\n"
		aux = aux.Siguiente
	}
	for i := 0; i < l.Longitud-1; i++ {
		c := i + 1
		texto += "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(c) + ";\n"
		texto += "nodo" + strconv.Itoa(c) + "->nodo" + strconv.Itoa(i) + ";\n"
		contador = c
	}
	texto += "nodo" + strconv.Itoa(contador) + "->nodo0 \n"
	texto += "nodo0 -> " + "nodo" + strconv.Itoa(contador) + "\n"
	texto += "}"
	crearArchivo(nombreArchivo)
	escribirArchivo(texto, nombreArchivo)
	ejecutar(nombreImagen, nombreArchivo)
}

// Funcion para buscar
func (l *ListaDobleCircular) Buscar(curso string) bool {
	if l.Longitud == 0 {
		return false
	} else {
		aux := l.Inicio
		contador := 1
		for l.Longitud >= contador {
			if aux.Tutor.Curso == curso {
				return true
			}
			aux = aux.Siguiente
			contador++
		}
	}
	return false
}

// Funcion para buscar Tutor
func (l *ListaDobleCircular) BuscarTutor(curso string) *NodoListaCircular {
	aux := l.Inicio
	contador := 1
	for l.Longitud >= contador {
		if aux.Tutor.Curso == curso {
			return aux
		}
		aux = aux.Siguiente
		contador++
	}
	return nil
}

// Función para verificar si ya existe un tutor para el curso dado en la lista circular
func (l *ListaDobleCircular) ExisteTutorEnCurso(codigoCurso string) bool {
	actual := l.Inicio
	for i := 0; i < l.Longitud; i++ {
		if actual.Tutor.Curso == codigoCurso {
			return true
		}
		actual = actual.Siguiente
	}
	return false
}

// Función para obtener la nota del tutor actual en el curso dado en la lista circular
func (l *ListaDobleCircular) ObtenerNotaTutorEnCurso(codigoCurso string) int {
	actual := l.Inicio
	for i := 0; i < l.Longitud; i++ {
		if actual.Tutor.Curso == codigoCurso {
			return actual.Tutor.Nota
		}
		actual = actual.Siguiente
	}
	// Puedes manejar el caso cuando no se encuentra el tutor según tu lógica
	return 0
}

// Función para sustituir el tutor actual en el curso dado en la lista circular
func (l *ListaDobleCircular) SustituirTutor(codigoCurso string, carnet int, nombre string, nota int) {
	actual := l.Inicio
	for i := 0; i < l.Longitud; i++ {
		if actual.Tutor.Curso == codigoCurso {
			actual.Tutor.Carnet = carnet
			actual.Tutor.Nombre = nombre
			actual.Tutor.Nota = nota
			return
		}
		actual = actual.Siguiente
	}
	// Puedes manejar el caso cuando no se encuentra el tutor según tu lógica
	fmt.Println("Error: Tutor no encontrado para sustitución")
}

// Función para eliminar el tutor de un curso específico en la lista circular
func (l *ListaDobleCircular) EliminarTutor(curso string) {
	if l.Longitud == 0 {
		return
	}

	// Buscar el tutor en la lista
	aux := l.Inicio
	for i := 0; i < l.Longitud; i++ {
		if aux.Tutor.Curso == curso {
			// Enlace para saltar el nodo actual
			aux.Anterior.Siguiente = aux.Siguiente
			aux.Siguiente.Anterior = aux.Anterior

			// Ajustar el inicio si es necesario
			if l.Inicio == aux {
				l.Inicio = aux.Siguiente
			}

			// Reducir la longitud de la lista
			l.Longitud--

			// Liberar memoria del nodo
			aux = nil
			return
		}
		aux = aux.Siguiente
	}
}

// Función para obtener la nota más alta en el curso dado en la lista circular
func (l *ListaDobleCircular) ObtenerNotaMasAltaEnCurso(codigoCurso string) int {
	if l.Longitud == 0 {
		return 0
	}

	maxNota := 0
	actual := l.Inicio
	for i := 0; i < l.Longitud; i++ {
		if actual.Tutor.Curso == codigoCurso && actual.Tutor.Nota > maxNota {
			maxNota = actual.Tutor.Nota
		}
		actual = actual.Siguiente
	}
	return maxNota
}
