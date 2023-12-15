package cursos

import (
	"fmt"
	"strconv"
)

type ArbolBB struct {
	Raiz *NodoArbol
}

// Funcion para poder insertar Nodos al arbol
func (a *ArbolBB) InsertarNodo(raiz *NodoArbol, nuevoNodo *NodoArbol) *NodoArbol {
	if raiz == nil {
		raiz = nuevoNodo
	} else {
		if raiz.Valor > nuevoNodo.Valor {
			raiz.Izquierdo = a.InsertarNodo(raiz.Izquierdo, nuevoNodo)
		} else {
			raiz.Derecho = a.InsertarNodo(raiz.Derecho, nuevoNodo)
		}
	}
	return raiz
}

// Funcion para la insercion de Elementos
func (a *ArbolBB) InsertarElemento(valor string) {
	nuevoNodo := &NodoArbol{Valor: valor}
	a.Raiz = a.InsertarNodo(a.Raiz, nuevoNodo)
}

// Funcion para reccorer
func (a *ArbolBB) recorridoInorder(raiz *NodoArbol) {
	if raiz == nil {
		if raiz.Izquierdo != nil {
			a.recorridoInorder(raiz.Izquierdo)
			fmt.Print("->")
		}
		fmt.Print(raiz.Valor, "")
		if raiz.Derecho != nil {
			fmt.Print("->")
			a.recorridoInorder(raiz.Derecho)
		}
	}
}

// Funcion para recorrer en Preorden
func (a *ArbolBB) recorridoPreorden(raiz *NodoArbol) {
	if raiz != nil {
		fmt.Print(raiz.Valor, "")
		if raiz.Izquierdo != nil {
			fmt.Print("->")
			a.recorridoPreorden(raiz.Izquierdo)
		}
		if raiz.Derecho != nil {
			fmt.Print("->")
			a.recorridoPreorden(raiz.Derecho)
		}
	}
}

// Funcion para hacer un recorrido en PostOrden
func (a *ArbolBB) recorridoPostOrden(raiz *NodoArbol) {
	if raiz != nil {
		if raiz.Izquierdo != nil {
			a.recorridoPostOrden(raiz.Izquierdo)
			fmt.Print("->")
		}
		if raiz.Derecho != nil {
			a.recorridoPostOrden(raiz.Derecho)
			fmt.Print("->")
		}
		fmt.Print(raiz.Valor, " ")
	}
}

//Funcion para recorrer los arboles

func (a *ArbolBB) Recorridos() {
	a.recorridoInorder(a.Raiz)
	fmt.Println("")
	a.recorridoPostOrden(a.Raiz)
	fmt.Println()
	a.recorridoPreorden(a.Raiz)
}

// Funcion para realizar la busqueda del Valor

func (a *ArbolBB) busqueda_arbol(valor string, raiz *NodoArbol) *NodoArbol {
	var valorEncontrado *NodoArbol
	if raiz != nil {
		if raiz.Valor == valor {
			valorEncontrado = raiz
		} else {
			if raiz.Valor > valor {
				valorEncontrado = a.busqueda_arbol(valor, raiz.Izquierdo)
			} else {
				valorEncontrado = a.busqueda_arbol(valor, raiz.Derecho)
			}
		}
	}

	return valorEncontrado
}

// Funcion para realizar una busqueda
func (a *ArbolBB) Busqueda(valor string) bool {
	buscarElemento := a.busqueda_arbol(valor, a.Raiz)
	if buscarElemento != nil {
		return true
	}
	return false
}

// Funcion para graficar
// Reporte Grafico
func (a *ArbolBB) Graficar() {
	cadena := ""
	nombre_archivo := "./Reportes/ArbolCursos.dot"
	nombre_imagen := "./Reportes/ArbolCursos.jpg"
	if a.Raiz != nil {
		cadena += "digraph arbol{ "
		cadena += a.retornarValoresArbol(a.Raiz, 0)
		cadena += "}"
	}
	crearArchivo(nombre_archivo)
	escribirArchivo(cadena, nombre_archivo)
	ejecutar(nombre_imagen, nombre_archivo)
}

func (a *ArbolBB) retornarValoresArbol(raiz *NodoArbol, indice int) string {
	cadena := ""
	numero := indice + 1
	if raiz != nil {
		cadena += "\""
		cadena += raiz.Valor
		cadena += "\" ;"
		if raiz.Izquierdo != nil && raiz.Derecho != nil {
			cadena += " x" + strconv.Itoa(numero) + " [label=\"\",width=.1,style=invis];"
			cadena += "\""
			cadena += raiz.Valor
			cadena += "\" -> "
			cadena += a.retornarValoresArbol(raiz.Izquierdo, numero)
			cadena += "\""
			cadena += raiz.Valor
			cadena += "\" -> "
			cadena += a.retornarValoresArbol(raiz.Derecho, numero)
			cadena += "{rank=same" + "\"" + (raiz.Izquierdo.Valor) + "\"" + " -> " + "\"" + (raiz.Derecho.Valor) + "\"" + " [style=invis]}; "
		} else if raiz.Izquierdo != nil && raiz.Derecho == nil {
			cadena += " x" + strconv.Itoa(numero) + " [label=\"\",width=.1,style=invis];"
			cadena += "\""
			cadena += raiz.Valor
			cadena += "\" -> "
			cadena += a.retornarValoresArbol(raiz.Izquierdo, numero)
			cadena += "\""
			cadena += raiz.Valor
			cadena += "\" -> "
			cadena += "x" + strconv.Itoa(numero) + "[style=invis]"
			cadena += "{rank=same" + "\"" + (raiz.Izquierdo.Valor) + "\"" + " -> " + "x" + strconv.Itoa(numero) + " [style=invis]}; "
		} else if raiz.Izquierdo == nil && raiz.Derecho != nil {
			cadena += " x" + strconv.Itoa(numero) + " [label=\"\",width=.1,style=invis];"
			cadena += "\""
			cadena += raiz.Valor
			cadena += "\" -> "
			cadena += "x" + strconv.Itoa(numero) + "[style=invis]"
			cadena += "; \""
			cadena += raiz.Valor
			cadena += "\" -> "
			cadena += a.retornarValoresArbol(raiz.Derecho, numero)
			cadena += "{rank=same" + " x" + strconv.Itoa(numero) + " -> \"" + (raiz.Derecho.Valor) + "\"" + " [style=invis]}; "
		}
	}
	return cadena
}
