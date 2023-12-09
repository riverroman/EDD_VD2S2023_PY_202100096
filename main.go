package main

import (
	"fmt"
	"paquete/tutores"
)

func main() {

	listatutos := tutores.Cola{Inicio: nil, Longitud: 0}
	listatutos.LeerArchivo("tutores.csv")

	listatutos.Primero()
	listatutos.Descolar()
	fmt.Println("-------------")

	listatutos.Primero()
	listatutos.Descolar()
	fmt.Println("-------------")

	listatutos.Primero()
	listatutos.Descolar()
	fmt.Println("-------------")

}
