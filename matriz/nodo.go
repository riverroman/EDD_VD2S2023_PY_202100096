package matriz

type Dato struct {
	Carnet_Tutor      int
	Carnet_Estudiante int
	Curso             string
}

type NodoMatriz struct {
	Siguiente *NodoMatriz
	Anterior  *NodoMatriz
	Arriba    *NodoMatriz
	Abajo     *NodoMatriz
	PosX      int
	PosY      int
	Dato      *Dato
}
