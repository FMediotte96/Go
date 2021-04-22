package main

import "fmt"

//type nombre interface (firma)
type User interface {
	permisos() int
	nombre() string
}

type Admin struct {
	name string
}

//Implementación de la interface
func (admin Admin) permisos() int {
	return 5
}

func (admin Admin) nombre() string {
	return admin.name
}

type Editor struct {
	name string
}

//Implementación de la interface
func (editor Editor) permisos() int {
	return 3
}

func (editor Editor) nombre() string {
	return editor.name
}

func auth(user User) string {
	permisos := user.permisos()
	if permisos > 4 {
		return user.nombre() + " tiene permisos de administrador"
	} else if permisos == 3 {
		return user.nombre() + " tiene permisos de editor"
	}
	return ""
}

func main() {
	admin := Admin{"Facundo"}
	editor := Editor{"Federico"}

	usuarios := []User{admin, editor}

	//iteración de array
	for _, usuario := range usuarios {
		fmt.Println(auth(usuario))
	}

}

/*
	Firma de métodos sin implementar
*/
