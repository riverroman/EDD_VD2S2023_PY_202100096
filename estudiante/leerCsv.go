package estudiante

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func (l *ListaDoble) LeerArchivo(ruta string) {
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

		valor, _ := strconv.Atoi(linea[0])
		l.AgregarOrdenado(valor, linea[1])
	}

}
