package main

import "fmt"

// Ejercicio 1 - Letras de una palabra
// La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada una de las letras por separado para deletrearla.
// 1.  Crear una aplicación que tenga una variable con la palabra e imprimir la cantidad de letras que tiene la misma.
// 2.  Luego, imprimir cada una de las letras.

func main() {
	palabra := "DigitalHouse2023"
	var cant int = len(palabra)
	fmt.Printf(" La palabra %s tiene %d letras \n", palabra, cant)

	for i := 0; i < cant; i++ {
		fmt.Println(string(palabra[i]))
	}
}

// Ejercicio 2 - A qué mes corresponde
// Realizar una aplicación que contenga una variable con el número del mes.
// 1. Según el número, imprimir el mes que corresponda en texto.
// 2. ¿Se te ocurre si se puede resolver de más de una manera? ¿Cuál elegirías y por qué?

func main1() {
	meses := []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio",
		"Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}

	mes := 8

	if mes >= 1 && mes <= 12 {
		for i := 0; i <= len(meses); i++ {
			if i == mes {
				fmt.Println(meses[i-1])
			}
		}
	}
}

// Ejercicio 3- Listado de nombres
// Una profesora de la universidad quiere tener un listado con todos sus estudiantes.
// Es necesario crear una aplicación que contenga dicha lista.
// Sus estudiantes son: Benjamin, Nahuel, Brenda, Marcos, Pedro,
// Axel, Alez, Dolores, Federico, Hernán, Leandro, Eduardo, Duvraschka.
// ----------------------------------------------------------
// Luego de dos clases, se sumó un estudiante nuevo.
// Es necesario agregarlo al listado, sin modificar el código que escribiste
// inicialmente. El nombre de la nueva estudiante es Gabriela.

func main2() {
	
	estudiantes := []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores",
		"Federico", "Hernan", "Leandro", "Eduardo", "Duvraschka"}

	
	listaEstudiantes := append(estudiantes, "Gabriela")

	fmt.Println(listaEstudiantes)

}


// Ejercicio 4 - Qué edad tiene...
// Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. Según el siguiente map, debemos imprimir la edad de Benjamin

// -----------------------------------------------------------------------
// Por otro lado, también es necesario:
// Saber cuántos de sus empleados son mayores de 21 años.
// Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
// Eliminar a Pedro del map.


func main3(){
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Println(employees["Benjamin"])

	olderEmployees := 0
	
	for _, edad := range employees {
		if edad > 21 {
			olderEmployees++
		}
	}
	fmt.Println(olderEmployees)

	employees["Federico"] = 25
	fmt.Println(employees)

	delete(employees,"Pedro")
	fmt.Println(employees)
}
