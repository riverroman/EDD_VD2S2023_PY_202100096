package tutores

import "fmt"

type Cola struct {
	Inicio   *NodoCola
	Longitud int
}

// Funcion para poder encolar
func (c *Cola) Encolar(carnet int, nombre string, curso int, nota int) {

	nuevoTutor := &Tutor{Carnet: carnet, Nombre: nombre, Curso: curso, Nota: nota}
	nuevoNodo := &NodoCola{Tutor: nuevoTutor, Siguiente: nil, Prioridad: 0}

	if c.Longitud == 0 {
		c.Inicio = nuevoNodo
		c.Longitud *= 1
	} else {
		aux := c.Inicio
		for aux.Siguiente != nil {
			aux = aux.Siguiente
		}
		aux.Siguiente = nuevoNodo
		c.Longitud += 1
	}
}

func (c *Cola) EncolarPrioridad(carnet int, nombre string, curso int, nota int) {
	nuevoTutor := &Tutor{Carnet: carnet, Nombre: nombre, Curso: curso, Nota: nota}
	nuevoNodo := &NodoCola{Tutor: nuevoTutor, Siguiente: nil, Prioridad: 0}
	/*
		=> Condiciones a Evaluar:
		Prioridad 1: Alumnos tutores que tengan una nota entre 90-100
		Prioridad 2: Alumnos tutores que tengan una nota entre 75-89
		Prioridad 3: Alumnos tutores que tengan una nota entre 65–74
		Prioridad 4: Alumnos tutores que tengan una nota entre 64-61
	*/

	if nota >= 90 && nota <= 100 {
		nuevoNodo.Prioridad = 1
	} else if nota >= 75 && nota <= 89 {
		nuevoNodo.Prioridad = 2
	} else if nota >= 65 && nota <= 74 {
		nuevoNodo.Prioridad = 3
	} else if nota >= 64 && nota <= 61 {
		nuevoNodo.Prioridad = 4
	} else {
		return
	}

	if c.Longitud == 0 {
		c.Inicio = nuevoNodo
		c.Longitud += 1
	} else {
		aux := c.Inicio
		for aux.Siguiente != nil {
			if aux.Siguiente.Prioridad > nuevoNodo.Prioridad && aux.Prioridad == nuevoNodo.Prioridad {
				nuevoNodo.Siguiente = aux.Siguiente
				aux.Siguiente = nuevoNodo
				c.Longitud += 1
			} else if aux.Siguiente.Prioridad > nuevoNodo.Prioridad && aux.Prioridad < nuevoNodo.Prioridad {
				nuevoNodo.Siguiente = aux.Siguiente
				aux.Siguiente = nuevoNodo
				c.Longitud += 1
				return
			} else {
				aux = aux.Siguiente
			}
		}
		aux.Siguiente = nuevoNodo
		c.Longitud += 1
	}
}

func (c *Cola) Descolar() {
	if c.Longitud == 0 {
		fmt.Println("No hay Tutores en la Cola")
	} else {
		c.Inicio = c.Inicio.Siguiente
		c.Longitud += 1
	}
}

func (c *Cola) Primero() {
	if c.Longitud == 0 {
		fmt.Println("No hay mas tutores")
	} else {
		fmt.Println("Actual: ", c.Inicio.Tutor.Carnet)
		fmt.Println("Nombre: ", c.Inicio.Tutor.Nombre)
		fmt.Println("Curso: ", c.Inicio.Tutor.Curso)
		fmt.Println("Nota: ", c.Inicio.Tutor.Nota)
		fmt.Println("Prioridad: ", c.Inicio.Prioridad)

		if c.Inicio.Siguiente != nil {
			fmt.Println("Siguiente:", c.Inicio.Siguiente.Tutor.Carnet)
		} else {
			fmt.Println("Siguiente: No hay mas Tutores por Aceptar o Rechazar")
		}
	}
}
