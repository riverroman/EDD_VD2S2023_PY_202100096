package matriz

import (
	"fmt"
	"strconv"
)

type Matriz struct {
	Raiz             *NodoMatriz
	Cantidad_Alumnos int
	Cantidad_Tutores int
}

// Buscar por Columna => Tutores
func (m *Matriz) buscarColumna(carnet_tutor int, curso string) *NodoMatriz {
	aux := m.Raiz
	for aux != nil {
		if aux.Dato.Carnet_Tutor == carnet_tutor && aux.Dato.Curso == curso {
			return aux
		}
		aux = aux.Siguiente
	}
	return nil
}

// Buscar por Fila  => Estudiantes
func (m *Matriz) buscarFila(carnet_estudiante int) *NodoMatriz {
	aux := m.Raiz
	for aux != nil {
		if aux.Dato.Carnet_Estudiante == carnet_estudiante {
			return aux
		}
		aux = aux.Siguiente
	}

	return nil
}

// Inserta Dato en COlumna
func (m *Matriz) insertarColumna(nuevoNodo *NodoMatriz, nodoRaiz *NodoMatriz) *NodoMatriz {
	temp := nodoRaiz
	piv := false

	for {
		if temp.PosX == nuevoNodo.PosX {
			temp.PosY = nuevoNodo.PosY
			temp.Dato = nuevoNodo.Dato
			return temp
		} else if temp.PosX > nuevoNodo.PosX {
			piv = true
			break
		}
		if temp.Siguiente != nil {
			temp = temp.Siguiente
		} else {
			break
		}
	}

	if piv {
		nuevoNodo.Siguiente = temp
		nuevoNodo.Anterior = temp.Anterior
		temp.Anterior.Siguiente = nuevoNodo
		temp.Anterior = nuevoNodo
	} else {
		nuevoNodo.Anterior = temp
		temp.Siguiente = nuevoNodo
	}

	return nuevoNodo
}

// Funcion para crear nuevas columnas => Tutores
func (m *Matriz) nuevaColumna(x int, carnet_tutor int, curso string) *NodoMatriz {
	nuevoNodo := &NodoMatriz{PosX: x, PosY: -1, Dato: &Dato{Carnet_Tutor: carnet_tutor, Carnet_Estudiante: 0, Curso: curso}}
	columna := m.insertarColumna(nuevoNodo, m.Raiz)
	return columna
}

// Funcion para insertar filas => Estudiantes
func (m *Matriz) insertarFila(nuevoNodo *NodoMatriz, nodoRaiz *NodoMatriz) *NodoMatriz {
	temp := nodoRaiz
	piv := false

	for {
		if temp.PosY == nuevoNodo.PosY {
			temp.PosX = nuevoNodo.PosX
			temp.Dato = nuevoNodo.Dato
			return temp
		} else if temp.PosY > nuevoNodo.PosY {
			piv = true
			break
		}
		if temp.Abajo != nil {
			temp = temp.Abajo
		} else {
			break
		}
	}
	if piv {
		nuevoNodo.Abajo = temp
		nuevoNodo.Arriba = temp.Arriba
		temp.Arriba.Abajo = nuevoNodo
		temp.Arriba = nuevoNodo
	} else {
		nuevoNodo.Arriba = temp
		temp.Abajo = nuevoNodo
	}
	return nuevoNodo
}

// Funcion para crear una nueva FIla
func (m *Matriz) nuevaFila(y int, carnet_estudiante int, curso string) *NodoMatriz {
	nuevoNodo := &NodoMatriz{PosX: -1, PosY: y, Dato: &Dato{Carnet_Tutor: 0, Carnet_Estudiante: carnet_estudiante, Curso: curso}}
	fila := m.insertarFila(nuevoNodo, m.Raiz)
	return fila
}

func (m *Matriz) Insertar_Elemento(carnet_estudiante int, carnet_tutor int, curso string) {
	nodoColumna := m.buscarColumna(carnet_tutor, curso)
	nodoFila := m.buscarFila(carnet_estudiante)

	if nodoColumna == nil && nodoFila == nil {
		nodoColumna = m.nuevaColumna(m.Cantidad_Tutores, carnet_tutor, curso)
		nodoFila = m.nuevaFila(m.Cantidad_Alumnos, carnet_estudiante, curso)
		m.Cantidad_Alumnos += 1
		m.Cantidad_Tutores += 1

		nuevoNodo := &NodoMatriz{PosX: nodoColumna.PosX, PosY: nodoFila.PosY, Dato: &Dato{Carnet_Tutor: carnet_tutor, Carnet_Estudiante: carnet_estudiante, Curso: curso}}
		nuevoNodo = m.insertarColumna(nuevoNodo, nodoFila)
		nuevoNodo = m.insertarFila(nuevoNodo, nodoColumna)

	} else if nodoColumna != nil && nodoFila == nil {
		nodoFila = m.nuevaFila(m.Cantidad_Alumnos, carnet_estudiante, curso)
		m.Cantidad_Alumnos += 1
		nuevoNodo := &NodoMatriz{PosX: nodoColumna.PosX, PosY: nodoFila.PosY, Dato: &Dato{Carnet_Tutor: carnet_tutor, Carnet_Estudiante: carnet_estudiante, Curso: curso}}
		nuevoNodo = m.insertarColumna(nuevoNodo, nodoFila)
		nuevoNodo = m.insertarFila(nuevoNodo, nodoColumna)

	} else if nodoColumna == nil && nodoFila != nil {
		nodoColumna = m.nuevaColumna(m.Cantidad_Tutores, carnet_tutor, curso)
		m.Cantidad_Tutores += 1
		nuevoNodo := &NodoMatriz{PosX: nodoColumna.PosX, PosY: nodoFila.PosY, Dato: &Dato{Carnet_Tutor: carnet_tutor, Carnet_Estudiante: carnet_estudiante, Curso: curso}}
		nuevoNodo = m.insertarColumna(nuevoNodo, nodoFila)
		nuevoNodo = m.insertarFila(nuevoNodo, nodoColumna)
	} else if nodoColumna != nil && nodoFila != nil {
		nuevoNodo := &NodoMatriz{PosX: nodoColumna.PosX, PosY: nodoFila.PosY, Dato: &Dato{Carnet_Tutor: carnet_tutor, Carnet_Estudiante: carnet_estudiante, Curso: curso}}
		nuevoNodo = m.insertarColumna(nuevoNodo, nodoFila)
		nuevoNodo = m.insertarColumna(nuevoNodo, nodoColumna)
	} else {
		fmt.Println("Err !")
	}

}

// Funcion para poder generar la grafica de la Matriz
func (m *Matriz) Graficar() {
	texto := ""
	nombre_archivo := "./Reportes/matriz.dot"
	nombre_imagen := "./Reportes/matrizDispersa.jpg"
	aux1 := m.Raiz
	aux2 := m.Raiz
	aux3 := m.Raiz
	if aux1 != nil {
		texto = "digraph MatrizCapa{ \n node[shape=box] \n rankdir=UD; \n {rank=min; \n"

		for aux1 != nil {
			texto += "nodo" + strconv.Itoa(aux1.PosX+1) + strconv.Itoa(aux1.PosY+1) + "[label=\"" + strconv.Itoa(aux1.Dato.Carnet_Tutor) + "\" ,rankdir=LR,group=" + strconv.Itoa(aux1.PosX+1) + "]; \n"
			aux1 = aux1.Siguiente
		}
		texto += "}"
		aux2 = aux2.Abajo
		for aux2 != nil {
			aux1 = aux2
			texto += "{rank=same; \n"
			flag_raiz := true
			for aux1 != nil {
				if flag_raiz {
					texto += "nodo" + strconv.Itoa(aux1.PosX+1) + strconv.Itoa(aux1.PosY+1) + "[label=\"" + strconv.Itoa(aux1.Dato.Carnet_Estudiante) + "\" ,group=" + strconv.Itoa(aux1.PosX+1) + "]; \n"
					flag_raiz = false
				} else {
					texto += "nodo" + strconv.Itoa(aux1.PosX+1) + strconv.Itoa(aux1.PosY+1) + "[label=\"" + aux1.Dato.Curso + "\" ,group=" + strconv.Itoa(aux1.PosX+1) + "]; \n"
				}

				aux1 = aux1.Siguiente
			}
			texto += "}"
			aux2 = aux2.Abajo
		}
		/** Conexiones entre los nodos de la matriz */
		aux2 = aux3
		for aux2 != nil {
			aux1 = aux2
			for aux1.Siguiente != nil {
				texto += "nodo" + strconv.Itoa(aux1.PosX+1) + strconv.Itoa(aux1.PosY+1) + " -> " + "nodo" + strconv.Itoa(aux1.Siguiente.PosX+1) + strconv.Itoa(aux1.Siguiente.PosY+1) + " [dir=both];\n"
				aux1 = aux1.Siguiente
			}
			aux2 = aux2.Abajo
		}
		aux2 = aux3
		for aux2 != nil {
			aux1 = aux2
			for aux1.Abajo != nil {
				texto += "nodo" + strconv.Itoa(aux1.PosX+1) + strconv.Itoa(aux1.PosY+1) + " -> " + "nodo" + strconv.Itoa(aux1.Abajo.PosX+1) + strconv.Itoa(aux1.Abajo.PosY+1) + " [dir=both];\n"
				aux1 = aux1.Abajo
			}
			aux2 = aux2.Siguiente
		}
		texto += "}"
	} else {
		texto = "No hay elementos en la matriz"
	}

	CrearArchivo(nombre_archivo)
	EscribirArchivo(texto, nombre_archivo)
	Ejecutar(nombre_imagen, nombre_archivo)
}
