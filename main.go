package main

import (
	"fmt"
	"paquete/estudiante"
	"paquete/tutores"
)

// Lista Doble Enlazada para Estudiantes
var listaDoble *estudiante.ListaDoble = &estudiante.ListaDoble{Inicio: nil, Contador: 0}

// Lista de Prioridad y Lista Circular para Tutores
var colaPrioridad *tutores.Cola = &tutores.Cola{Inicio: nil, Longitud: 0}
var listaCircular *tutores.ListaDobleCircular = &tutores.ListaDobleCircular{Inicio: nil, Longitud: 0}

func main() {

	opc := 0
	salir := false

	for !salir {
		fmt.Println("\n-------------------------")
		fmt.Println(" 1. 💻 Inicio de Sesion")
		fmt.Println(" 2. ❌ Salir")
		fmt.Println("-------------------------")
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
	fmt.Print(" 👷 Usuario: ")
	fmt.Scanln(&usuario)
	fmt.Print("\n 🔑 Password: ")
	fmt.Scanln(&password)
	fmt.Println("--------------------------------")

	if usuario == "ADMIN_202100096" && password == "admin" {
		MenuAdmin()
	} else if listaDoble.Buscar(usuario, password) {
		fmt.Println("Holaaa")
	} else {
		fmt.Println("\n| Usuario y Password Incorrecto |")
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
			continue
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
			listaCircular.Agregar(colaPrioridad.Inicio.Tutor.Carnet, colaPrioridad.Inicio.Tutor.Nombre, colaPrioridad.Inicio.Tutor.Curso, colaPrioridad.Inicio.Tutor.Nota)
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
			continue
		} else if opcion == 4 {
			continue
		} else if opcion == 5 {
			return
		} else {
			fmt.Println("\n| Ingrese una opcion valida |")
		}

	}
}
