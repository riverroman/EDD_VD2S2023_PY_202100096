package main

import (
	"fmt"
	"paquete/cursos"
	"paquete/estudiante"
	"paquete/matriz"
	"paquete/tutores"
	"strconv"
)

// Lista Doble Enlazada para Estudiantes
var listaDoble *estudiante.ListaDoble = &estudiante.ListaDoble{Inicio: nil, Contador: 0}

// Lista de Prioridad y Lista Circular para Tutores
var colaPrioridad *tutores.Cola = &tutores.Cola{Inicio: nil, Longitud: 0}
var listaCircular *tutores.ListaDobleCircular = &tutores.ListaDobleCircular{Inicio: nil, Longitud: 0}

// Matriz dispersa
var matrizDispersa *matriz.Matriz = &matriz.Matriz{Raiz: &matriz.NodoMatriz{PosX: -1, PosY: -1, Dato: &matriz.Dato{Carnet_Tutor: 0, Carnet_Estudiante: 0, Curso: "Raiz"}}, Cantidad_Alumnos: 0, Cantidad_Tutores: 0}

// Arbol
var arbol *cursos.ArbolAVL = &cursos.ArbolAVL{Raiz: nil}

var loggeado_estudiante string = ""

func main() {

	opc := 0
	salir := false

	for !salir {
		fmt.Println("\n-------------------------")
		fmt.Println("\n 1. 💻 Inicio de Sesion")
		fmt.Println("\n 2. ❌ Salir")
		fmt.Println("\n-------------------------")
		fmt.Print("\nIngrese una opcion: ")

		fmt.Scanln(&opc)
		switch opc {
		case 1:
			MenuLogin()
		case 2:
			fmt.Println("\nSaliendo...")
			fmt.Println("")
			salir = true
		}
	}
}

func MenuLogin() {
	usuario := ""
	password := ""
	fmt.Println("\n--------------------------------")
	fmt.Print("\n 👷 Usuario: ")
	fmt.Scanln(&usuario)
	fmt.Print("\n 🔑 Password: ")
	fmt.Scanln(&password)
	fmt.Println("\n--------------------------------")

	if usuario == "ADMIN_202100096" && password == "Admin" {
		fmt.Println("\n ⛅︎ Bienvenido al sistema:", usuario)
		MenuAdmin()
	} else if listaDoble.Buscar(usuario, password) {
		fmt.Println("\n 👷 Bienvenido Estudiante: ", usuario)
		loggeado_estudiante = usuario
		MenuEstudiante()
	} else {
		fmt.Println("\n| Usuario y Password Incorrecto |")
	}
}

func MenuEstudiante() {
	opcion := 0
	salir := false
	for !salir {
		fmt.Println("\n--------------------------------")
		fmt.Println("  => 1. Ver Tutores Disponibles")
		fmt.Println("  => 2. Asignarse a Tutores    ")
		fmt.Println("  => 3. Regresar               ")
		fmt.Println("--------------------------------")
		fmt.Print("\nIngrese una opcion: ")
		fmt.Scanln(&opcion)

		if opcion == 1 {
			fmt.Println("\n 👷  Tutores Disponibles  ")
			fmt.Println("")
			listaCircular.Mostrar()
		} else if opcion == 2 {
			AsignacionCursos()
		} else if opcion == 3 {
			return
		} else {
			fmt.Println("\n| Ingrese una opcion Valida |")
		}
	}
}

func MenuAdmin() {
	opcion := 0
	salir := false
	for !salir {

		fmt.Println("\n------------------------------------------")
		fmt.Println("  =>	1. Carga de Estudiantes Tutores	")
		fmt.Println("  =>	2. Carga de Estudiantes			")
		fmt.Println("  =>	3. Carga Cursos al Sistema		")
		fmt.Println("  =>	4. Control de Estudiantes Tutores")
		fmt.Println("  =>	5. Reportes Estructuras			")
		fmt.Println("  =>	6. Salir						")
		fmt.Println("------------------------------------------")

		fmt.Print("\nIngrese una opcion: ")
		fmt.Scanln(&opcion)
		fmt.Println("")

		switch opcion {
		case 1:
			cargarTutores()
		case 2:
			cargarEstudiantes()
		case 3:
			cargarCursos()
		case 4:
			ControlEstudiantes()
		case 5:
			areaReportes()
		case 6:
			return
		}
	}
}

func cargarTutores() {
	ruta := ""
	fmt.Print("\nNombre de Archivo: ")
	fmt.Scanln(&ruta)
	colaPrioridad.LeerArchivo(ruta)
	fmt.Println("\nSe cargo correctamente el archivo:", ruta)
	colaPrioridad.Imprimir()
}

func cargarEstudiantes() {
	ruta := ""
	fmt.Print("\nNombre de Archivo: ")
	fmt.Scanln(&ruta)
	listaDoble.LeerArchivo(ruta)
	fmt.Println("\nSe cargo correctamente el archivo:", ruta)
	fmt.Println("")
	listaDoble.Imprimir()
}

func cargarCursos() {
	ruta := ""
	fmt.Print("\nNombre de Archivo: ")
	fmt.Scanln(&ruta)
	arbol.LeerJson(ruta)
	fmt.Println("\nSe cargo correctamente el archivo: ", ruta)
}

func ControlEstudiantes() {
	opcion := 0
	salir := false
	for !salir {

		colaPrioridad.Primero()

		fmt.Println("\n---------------------")
		fmt.Println("    1. Aceptar       ")
		fmt.Println("    2. Rechazar      ")
		fmt.Println("    3. Salir         ")
		fmt.Println("---------------------")
		fmt.Print("\nIngrese una opcion: ")
		fmt.Scanln(&opcion)

		if opcion == 1 {
			curso := colaPrioridad.Inicio.Tutor.Curso

			if listaCircular.ExisteTutorEnCurso(curso) {
				notaNueva := colaPrioridad.Inicio.Tutor.Nota
				notaActual := listaCircular.ObtenerNotaMasAltaEnCurso(curso)

				if notaNueva > notaActual {
					fmt.Println("\nSe sustituyó tutor de curso actual")
					fmt.Println("")
					listaCircular.SustituirTutor(curso, colaPrioridad.Inicio.Tutor.Carnet, colaPrioridad.Inicio.Tutor.Nombre, notaNueva)
				} else {
					fmt.Println("\nSe registró tutor con éxito")
					fmt.Println("")
					listaCircular.Agregar(colaPrioridad.Inicio.Tutor.Carnet, colaPrioridad.Inicio.Tutor.Nombre, curso, notaNueva)
				}
			} else {
				listaCircular.Agregar(colaPrioridad.Inicio.Tutor.Carnet, colaPrioridad.Inicio.Tutor.Nombre, curso, colaPrioridad.Inicio.Tutor.Nota)
			}
			colaPrioridad.Descolar()
		} else if opcion == 2 {
			colaPrioridad.Descolar()
		} else if opcion == 3 {
			return
		} else {
			fmt.Println("\n| Ingrese una opcion valida |")
			fmt.Println("")
		}
	}
}

func AsignacionCursos() {
	opcion := ""
	salir := false
	for !salir {
		fmt.Print("\nIngrese el codigo del curso: ")
		fmt.Scanln(&opcion)

		if arbol.Busqueda(opcion) {
			if listaCircular.Buscar(opcion) {
				TutorBuscado := listaCircular.BuscarTutor(opcion)
				estudiante, err := strconv.Atoi(loggeado_estudiante)
				if err != nil {
					break
				}
				matrizDispersa.Insertar_Elemento(estudiante, TutorBuscado.Tutor.Carnet, opcion)
				fmt.Println("\nSe asigno el curso:", opcion)
				break
			} else {
				fmt.Println("\n No hay tutores para el curso:", opcion)
				break
			}

		} else {
			fmt.Println("\n El curso no existe en el sistema:", opcion)
			break
		}
	}
}

func areaReportes() {
	salir := false
	opcion := 0
	for !salir {
		fmt.Println("\n 📚 Area de Reportes")
		fmt.Println("\n1. Reporte de Alumnos")
		fmt.Println("2. Reporte de Tutores Aceptados")
		fmt.Println("3. Reporte de Asignacion")
		fmt.Println("4. Reporte de Cursos")
		fmt.Println("5. Salir")
		fmt.Print("\nIngrese una opcion: ")
		fmt.Scanln(&opcion)
		fmt.Println("")

		if opcion == 1 {
			listaDoble.Graficar()
		} else if opcion == 2 {
			listaCircular.Graficar()
		} else if opcion == 3 {
			matrizDispersa.Graficar()
		} else if opcion == 4 {
			arbol.Graficar()
		} else if opcion == 5 {
			return
		} else {
			fmt.Println("\n| Ingrese una opcion valida |")
		}
	}
}
