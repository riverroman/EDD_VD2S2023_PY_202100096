package tutores


type Tutor struct{
	Carnet int
	Nombre string
	Curso int
	Nota int
}


type NodoCola struct{
	Tutor *Tutor
	Siguiente *NodoCola
	Prioridad int
}

