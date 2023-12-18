package tutores

import (
	"fmt"
	"os"
	"text/tabwriter"
)

type Cola struct {
	Inicio   *NodoCola
	Longitud int
}

func (c *Cola) Encolar(carnet int, nombre string, curso string, nota int) {
	nuevoTutor := &Tutor{Carnet: carnet, Nombre: nombre, Curso: curso, Nota: nota}
	nuevoNodo := &NodoCola{Tutor: nuevoTutor, Siguiente: nil, Prioridad: 0}

	if c.Longitud == 0 {
		c.Inicio = nuevoNodo
		c.Longitud++
	} else {
		aux := c.Inicio
		for aux.Siguiente != nil {
			aux = aux.Siguiente
		}
		aux.Siguiente = nuevoNodo
		c.Longitud++
	}
}

func (c *Cola) EncolarPrioridad(carnet int, nombre string, curso string, nota int) {
	nuevoTutor := &Tutor{Carnet: carnet, Nombre: nombre, Curso: curso, Nota: nota}
	nuevoNodo := &NodoCola{Tutor: nuevoTutor, Siguiente: nil, Prioridad: 0}

	// Asignar prioridad según las reglas proporcionadas
	if nota >= 90 && nota <= 100 {
		nuevoNodo.Prioridad = 1
	} else if nota >= 75 && nota <= 89 {
		nuevoNodo.Prioridad = 2
	} else if nota >= 65 && nota <= 74 {
		nuevoNodo.Prioridad = 3
	} else if nota >= 61 && nota <= 64 {
		nuevoNodo.Prioridad = 4
	} else {
		return
	}

	// Caso especial: nuevo inicio de la cola
	if c.Longitud == 0 || nuevoNodo.Prioridad < c.Inicio.Prioridad {
		nuevoNodo.Siguiente = c.Inicio
		c.Inicio = nuevoNodo
		c.Longitud++
		return
	}

	// Encontrar la posición correcta para el nuevo nodo
	aux := c.Inicio
	for aux.Siguiente != nil && aux.Siguiente.Prioridad <= nuevoNodo.Prioridad {
		aux = aux.Siguiente
	}

	// Insertar el nuevo nodo en la posición correcta
	nuevoNodo.Siguiente = aux.Siguiente
	aux.Siguiente = nuevoNodo
	c.Longitud++
}

func (c *Cola) Descolar() {
	if c.Longitud == 0 {
		fmt.Println("No hay tutores en la cola")
	} else {
		c.Inicio = c.Inicio.Siguiente
		c.Longitud--
	}
}

func (c *Cola) Primero() {
	if c.Longitud == 0 {
		fmt.Println("No hay mas Tutores")
	} else {
		fmt.Println("Actual: ", c.Inicio.Tutor.Carnet)
		fmt.Println("Nombre: ", c.Inicio.Tutor.Nombre)
		fmt.Println("Curso: ", c.Inicio.Tutor.Curso)
		fmt.Println("Nota: ", c.Inicio.Tutor.Nota)
		fmt.Println("Prioridad: ", c.Inicio.Prioridad)
		if c.Inicio.Siguiente != nil {
			fmt.Println("Siguiente: ", c.Inicio.Siguiente.Tutor.Carnet)
		} else {
			fmt.Print("Siguiente: No hay mas tutores por evaluar")
		}
	}
}

// Imprimir Tutores cargados
func (c *Cola) Imprimir() {
	if c.Longitud == 0 {
		fmt.Println("\nCola de Prioridad Vacia")
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "\nCarnet\tNombre\tCurso\tNota")

	aux := c.Inicio
	for aux != nil {
		fmt.Fprintf(w, "%d\t%s\t%s\t%d\t", aux.Tutor.Carnet, aux.Tutor.Nombre, aux.Tutor.Curso, aux.Tutor.Nota)
		fmt.Fprintln(w, "")
		aux = aux.Siguiente
	}

	w.Flush()
}
