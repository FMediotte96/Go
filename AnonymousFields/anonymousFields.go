package main

import "fmt"

/*
	Los anonymous fields nos permiten replicar el comportamiento de la herencia en lenguajes OO
*/

type Human struct {
	name string
}

func (human Human) hablar() string {
	return "bla bla bla"
}

type Dummy struct {
	name string
}

//El tutor tiene un campo Human
//anonymous fields se declaran sin nombre, solo el tipo
type Tutor struct {
	Human
	Dummy
}

func (tutor Tutor) hablar() string {
	return tutor.Human.hablar() + " Bienvenidos al curso"
}

func main() {
	tutor := Tutor{Human{"Facundo"}, Dummy{"Federico"}}

	fmt.Println(tutor.Human.name)
	//podemos acceder a los atributos y métodos de la herencia sin referenciarla tutor.name
	//ambiguous selector tutor.name en este caso tenemos que definir la estructura a la hora de acceder
	//al campo que queremos referenciar

	fmt.Println(tutor.hablar()) //podemos sobrescribir el método de la estructura Human
}
