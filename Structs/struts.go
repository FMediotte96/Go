package main

import "fmt"

/*
	type -> indica que vamos a definir un nuevo tipo
	User -> nombre del tipo
	struct -> a su vez es una estructura, listando atributos que va a tener el nuevo tipo

	Las estructuras son mutables
*/
type User struct {
	edad     int
	nombre   string
	apellido string
}

func main() {
	//se puede indicar el nombre del atributo o inicializarlo con el orden declarado
	facundo := User{edad: 25, nombre: "Facundo", apellido: "Mediotte"}
	diego := User{27, "Diego", "Mediotte"}

	fer := new(User) //new nos crea un puntero

	//Dos formas de asignar, como es un puntero go nos ofrece la posibilidad de obviar el *
	(*fer).nombre = "Fernando"
	fer.edad = 25

	fmt.Println(facundo)
	fmt.Println(diego)
	fmt.Println((*fer))
	fmt.Println(fer.nombre)
}
