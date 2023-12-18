package tutores

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func (c *Cola) LeerArchivo(ruta string) {
	file, err := os.Open(ruta)
	if err != nil {
		fmt.Println("No se pudo abrir el archivo")
	}

	defer file.Close()

	lectura := csv.NewReader(file)
	lectura.Comma = ','
	encabezado := true

	for {
		linea, err := lectura.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("No se pudo leer la linea del csv")
		}
		if encabezado {
			encabezado = false
			continue
		}

		carnet, _ := strconv.Atoi(linea[0])
		nota, _ := strconv.Atoi(linea[3])
		c.EncolarPrioridad(carnet, linea[1], "0"+linea[2], nota)

	}
}
