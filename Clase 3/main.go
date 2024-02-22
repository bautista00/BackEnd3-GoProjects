package main

import (
	"fmt"
)

/*
Registro de estudiantes
Una universidad necesita registrar a los estudiantes y generar una funcionalidad para imprimir
el detalle de los datos de cada uno de ellos, de la siguiente manera:
Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]
Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos.
Para ello es necesario generar una estructura Alumno con las variables Nombre, Apellido, DNI, Fecha
y que tenga un método detalle.
*/

type Alumno struct {
	nombre   string
	apellido string
	dni      int
	fecha    string
}

func (a Alumno) Detalle() {

	fmt.Printf("\nDetalle del alumno:")
	fmt.Printf("\nNombre: %s \nApellido: %s \nDNI: %v \nFecha: %s\n", a.nombre, a.apellido, a.dni, a.fecha)

}

func main2() {

	alumno := Alumno{
		nombre:   "Carlos",
		apellido: "Molina",
		dni:      95866074,
		fecha:    "25/10/2023",
	}

	alumno2 := Alumno{
		nombre:   "Pablo",
		apellido: "Gonzalez",
		dni:      36564453,
		fecha:    "29/01/2015",
	}

	Alumno.Detalle(alumno)
	Alumno.Detalle(alumno2)

}

// -----------------------Ejercicio 1-------------
// Crear un programa que cumpla los siguientes puntos:
// Tener una estructura llamada Product con los campos ID, Name, Price, Description y Category.
// Tener un slice global de Product llamado Products, instanciado con valores.
// Dos métodos asociados a la estructura Product: Save(), GetAll().
// El método Save() deberá tomar el slice de Products y añadir el producto desde
// el cual se llama al método. El método GetAll() deberá imprimir todos los productos
// guardados en el slice Products.
// Una función getById() a la cual se le deberá pasar un INT como parámetro y retorna el producto correspondiente al parámetro pasado.
// Ejecutar al menos una vez cada método y función definidos desde main().

type Product struct {
	ID          int
	Name        string
	Price       int
	Description string
	Category    string
}

var Products []Product = []Product{
	{ID: 1, Name: "Leche", Price: 10, Description: "Descremada", Category: "Lacteo"},
	{ID: 2, Name: "Chocolate", Price: 5, Description: "Amargo", Category: "Golosina"},
}

func (p *Product) Save() {
	Products = append(Products, *p)
}

func (p Product) GetAll() {
	for _, product := range Products {
		fmt.Printf("ID: %d, Name: %s, Price: %d, Description: %s, Category: %s\n", product.ID, product.Name, product.Price, product.Description, product.Category)
	}
}

func getById(id int) Product {
	for _, product := range Products {
		if product.ID == id {
			return product
		}
	}

	return Product{}
}

func main() {
	newProduct := Product{
		ID:          3,
		Name:        "Dulce de Leche",
		Price:       15,
		Description: "La Serenisa",
		Category:    "Lacteo",
	}
	newProduct.Save()

	newProduct.GetAll()

	idToLook := 2
	foundProduct := getById(idToLook)

	if foundProduct.ID != 0 {
		fmt.Printf("Producto encontrado con ID %d: %v\n", idToLook, foundProduct)
	} else {
		fmt.Printf("No se encontró ningún producto con ID %d\n", idToLook)
	}

}
