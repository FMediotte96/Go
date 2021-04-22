package main

//Podemos agregar métodos a las estructuras declaradas en el mismo paquete no de otro paquete

import "fmt"

type User struct {
	edad     int
	nombre   string
	apellido string
}

// func (nombreDato tipo) nombre_fn(params) retorno {} firma de una función
func (user User) nombre_completo() string {
	return user.nombre + " " + user.apellido
}

//Cada vez que pasamos un argumento en Go,
//el mismo se pasa como una copia, es decir se duplica el objeto
//excepto que pasemos como tipo de dato el puntero
//Otro ventaja es que se evita copiar toda la estructura que la ref a memoria
func (user *User) setName(n string) {
	user.nombre = n
}

func main() {
	user := new(User)

	user.nombre = "Facundo"
	user.apellido = "Mediotte"
	user.edad = 25

	user.setName("Facu")

	fmt.Println((*user))
	fmt.Println(user.nombre_completo())
}
