package main

import (
	"fmt"
)

// Consigna
// Una empresa de redes sociales requiere implementar una estructura usuarios con funciones
// que vayan agregando información a la misma. Para optimizar y ahorrar memoria requieren
// que la estructura usuarios ocupe el mismo lugar en memoria para el main del programa y
// para las funciones. La estructura debe tener los campos: nombre, apellido, edad, correo
// y contraseña. Y deben implementarse las funciones:
// cambiarNombre: permite cambiar el nombre y apellido.
// cambiarEdad: permite cambiar la edad.
// cambiarCorreo: permite cambiar el correo.
// cambiarContraseña: permite cambiar la contraseña.
// *

type usuario struct {
	nombre     string
	apellido   string
	edad       int
	correo     string
	contraseña string
}

func main2() {
	fmt.Println("Agregar usuario: ")

	user := usuario{
		nombre:     "Juan",
		apellido:   "Peres",
		edad:       27,
		correo:     "juanperez@gmail.com",
		contraseña: "123facil",
	}
	fmt.Println(user)

	var usPuntero *usuario

	usPuntero = &user

	fmt.Println("\nCambiar nombre:")
	fmt.Println(usPuntero.cambiarNombre("Juan Martin", "Perez"))

	fmt.Println("Cambiar edad:")
	fmt.Println(usPuntero.cambiarEdad(28))

	fmt.Println("Cambiar correo:")
	fmt.Println(usPuntero.cambiarCorreo("jperez28@gmail.com"))

	fmt.Println("Cambiar contraseña:")
	fmt.Println(usPuntero.cambiarContraseña("456dificil"))

	fmt.Println("\nUser cambiado totalmente:")
	fmt.Println(user)

}

func (c *usuario) cambiarNombre(nombre, apellido string) string {
	c.nombre = nombre
	c.apellido = apellido
	return fmt.Sprintln(c)
}

func (c *usuario) cambiarEdad(edad int) string {
	c.edad = edad
	return fmt.Sprintln(c)
}

func (c *usuario) cambiarCorreo(correo string) string {
	c.correo = correo
	return fmt.Sprintln(c)
}

func (c *usuario) cambiarContraseña(contraseña string) string {
	c.contraseña = contraseña
	return fmt.Sprintln(c)
}

// -----------------Ejercicio 1---------------------------------
// Productos
// Algunas tiendas de e-commerce necesitan realizar una funcionalidad en Go para administrar
// productos y retornar el valor del precio total. La empresa tiene tres tipos de productos:
// Pequeño, Mediano y Grande. Pero se espera que sean muchos más.
// Los costos adicionales para cada uno son:
// Pequeño: solo tiene el costo del producto.
// Mediano: el precio del producto + un 3% por mantenerlo en la tienda.
// Grande: el precio del producto + un 6% por mantenerlo en la tienda y un adicional de $2500 de costo de envío.
// El porcentaje de mantenerlo en la tienda se calcula sobre el precio del producto.
// Se requiere una función factory que reciba el tipo de producto y el precio, y retorne
// una interfaz Producto que tenga el método Precio.
// Se debe poder ejecutar el método Precio y que el método devuelva el precio total
// basándose en el costo del producto y los adicionales, en caso de que los tenga.

type Product interface {
	Precio() float64
}

type Pequeño struct {
	precioProducto float64
}

type Mediano struct {
	precioProducto float64
}

type Grande struct {
	precioProducto float64
}

func (p Pequeño) Precio() float64 {
	return p.precioProducto
}

func (m Mediano) Precio() float64 {
	return m.precioProducto * 1.03
}

func (g Grande) Precio() float64 {
	return g.precioProducto * 1.06 + 2500
}

func CreateProduct(tipo string, precio float64) Product {
	switch tipo{
	case "Pequeño":
		return Pequeño{precioProducto: precio}
    case "Mediano":
        return Mediano{precioProducto: precio}
    case "Grande":
        return Grande{precioProducto: precio}
    default:
        return nil
    }
}

func main() {
    
    productoPequeño := CreateProduct("Pequeño", 100)
    productoMediano := CreateProduct("Mediano", 200)
    productoGrande := CreateProduct("Grande", 300)

    fmt.Println("Precio Pequeño:", productoPequeño.Precio())
    fmt.Println("Precio Mediano:", productoMediano.Precio())
    fmt.Println("Precio Grande:", productoGrande.Precio())
}

