package tutores

type Tutor struct {
	Carnet int
	Nombre string
	Curso  string
	Nota   int
}

type NodoCola struct {
	Tutor     *Tutor
	Siguiente *NodoCola
	Prioridad int
}

type NodoListaCircular struct {
	Tutor     *Tutor
	Siguiente *NodoListaCircular
	Anterior  *NodoListaCircular
}
