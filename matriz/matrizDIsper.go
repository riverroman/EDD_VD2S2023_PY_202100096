package matriz

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
	nuevoNOdo := &NodoMatriz{PosX: -1, PosY: y, Dato: &Dato{Carnet_Tutor: 0, Carnet_Estudiante: carnet_estudiante, Curso: curso}}
	fila := m.insertarFila(nuevoNOdo, m.Raiz)
	return fila
}

func (m *Matriz) Insertar_Elemento(carnet_estudiante int, carnet_tutor int, curso string) {
	nodoColumna := m.buscarColumna(carnet_tutor, curso)
	nodoFila := m.buscarFila(carnet_estudiante)

}
