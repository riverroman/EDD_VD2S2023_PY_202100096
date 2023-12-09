package estudiante

type NodolistaDoble struct {
	Estudiante *Estudiante
	Siguiente  *NodolistaDoble
	Anterior   *NodolistaDoble
}
