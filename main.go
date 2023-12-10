package main

import (
	"fmt"
	"paquete/tutores"
)

var colaPrioridad *tutores.Cola = &tutores.Cola{Inicio: nil, Longitud: 0}

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

		switch opcion {
		case 1:
			CargarTutores()
		case 2:
			continue
		case 3:
			continue
		case 4:
			continue
		case 5:
			continue
		case 6:
			return
		}
	}
}

func CargarTutores() {
	ruta := ""
	fmt.Print("\nNombre de Archivo: ")
	fmt.Scanln(&ruta)
	colaPrioridad.LeerArchivo(ruta)
	fmt.Println("\nSe cargo correctamente el archivo:", ruta)
}
