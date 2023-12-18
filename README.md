## ESTRUCTURA DE DATOS VD2S2023 PY 202100096 - FASE 1

## UNIVERSIDAD DE SAN CARLOS DE GUATEMALA


<p align="center">

|**CARNET**  |      **NOMBRE COMPLETO**          |  
|----------|:-----------------------------------:|
|202100096 |  RIVER ANDERSON - ISMALEJ ROMAN     |    
| AUXILIAR |            CRISTIAN SUY             |   
| SECCION  |                "A"                  |  

</p>

#### 📌 MANUAL USUARIO

### **Objetivos Generales**
* Aplicar los conocimientos del curso Estructuras de Datos en el desarrollo de las diferentes estructuras de datos y los diferentes algoritmos de manipulación de la información en ellas.

---
>### **Login**
* Al iniciar el programa, contaremos con un login para iniciar sesión como Administrador o como Estudiante.

<p align="center">
    <img src="./Reportes/imagenes/login.png">
</p>

---
>### **Menu Administrador**
* Al iniciar sesión como administrador tendremos diferentes acciones por hacer cargar archivos de entradas, generar reportes gráficos y tener control sobre aceptar o rechazar Tutores Académicos.

<p align="center">
    <img src="./Reportes/imagenes/menu_administrador.png">
</p>

---
#### **Carga de Tutores**

<p align="center">
    <img src="./Reportes/imagenes/tutores.png">
</p>

---
### **Carga de Estudiantes**

<p align="center">
    <img src="./Reportes/imagenes/estudiantes.jpg">
</p>

---

### **Carga de Cursos**

<p align="center">
    <img src="./Reportes/imagenes/cursos.png">
</p>

---
---
* Nota: Para cargar los archivos de entrada se debe colocar el nombre del archivo con extensión .csv.

---

### **Control de Estudiantes Tutores**

* Luego de cargar los tutores, contaremos con una opción para elegir qué tutores aceptar.

<p align="center">
    <img src="./Reportes/imagenes/control_tutores.png">
</p>

---

### **Reportes**

* En el Área de Reportes podremos generar gráficos, para poder visualizar las estructuras que se implementaron en el proyecto.

<p align="center">
    <img src="./Reportes/imagenes/reportes.png">
</p>

>### **Menu Estudiante**
* Al iniciar sesión como estudiante podremos ver los tutores disponibles y poder asignarnos a un tutor en específico ingresando el número de curso.

<p align="center">
    <img src="./Reportes/imagenes/menu_estudiante.jpg">
</p>

---

>#### **Tutores Disponibles**
* En ver Tutores Disponibles tendremos los tutores que fueron aceptados por el Administrador, con el curso que imparten mas su nombre.

<p align="center">
    <img src="./Reportes/imagenes/tutores_disponible.png">
</p>

---
>#### **Asignacion de Cursos**
* Para poder asignarnos un curso debemos de seleccionar la opcion de Asignarse a Tutores, luego de eso ingresar el codigo del curso, si el curso existe en sistema se asigna satisfactoriamente de lo contrario no.

<p align="center">
    <img src="./Reportes/imagenes/asignacion_cursos.png">
</p>

---

#### 📌 MANUAL TECNICO

>**Lista Enlazada Doble**

>##### Se utilizo una estructura para el manejo de ordenado de estudiante por medio de su carnet.

    func (l *ListaDoble) AgregarOrdenado(carnet int, nombre string) {
	    nuevoEstudiante := &Estudiante{Carnet: carnet, Nombre: nombre}
	    nuevoNodo := &NodolistaDoble{Estudiante: nuevoEstudiante, Siguiente: nil, Anterior: nil}

	    if l.Contador == 0 {
		    l.Inicio = nuevoNodo
		    l.Inicio.Anterior = nuevoNodo
		    l.Inicio.Siguiente = nuevoNodo
		    l.Contador += 1
	    } else {
		    aux := l.Inicio
		    contador := 1
		    for contador < l.Contador {
			    if l.Inicio.Estudiante.Carnet > carnet {
				    nuevoNodo.Siguiente = l.Inicio
				    nuevoNodo.Anterior = l.Inicio.Anterior
				    l.Inicio.Anterior = nuevoNodo
				    l.Inicio = nuevoNodo
				    l.Contador++
				    return
			    }

			    if aux.Estudiante.Carnet < carnet {
				    aux = aux.Siguiente
			    } else {
				    nuevoNodo.Anterior = aux.Anterior
				    aux.Anterior.Siguiente = nuevoNodo
				    nuevoNodo.Siguiente = aux
				    aux.Anterior = nuevoNodo
				    l.Contador += 1
				    return
			    }
			    contador += 1
		    }
		    if aux.Estudiante.Carnet > carnet {
			    nuevoNodo.Siguiente = aux
			    nuevoNodo.Anterior = aux.Anterior
			    aux.Anterior.Siguiente = nuevoNodo
			    aux.Anterior = nuevoNodo
			    l.Contador += 1
			    return
		}
		    nuevoNodo.Anterior = aux
		    nuevoNodo.Siguiente = l.Inicio
		    aux.Siguiente = nuevoNodo
		    l.Contador += 1
	    }
    }









