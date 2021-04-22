package main

import "fmt"

func main() {
	//var nombre_variable tipo_dato declaración o := no necesitamos declarar el tipo con este simbolo
	nombre := "Coco" //version tipado dinamico de go
	//nombre := "Cocooo" //como ya existe la variable lanza un error
	//nombre = 23 //no puedo asignar un nuevo tipo de dato
	nombre = "Facu"
	fmt.Println(nombre)
}

/*
	GO ES FUERTEMENTE TIPADO
	Todas las variables deben tener un valor inicial
	int 0
	cadena vacia
	bool falso
*/
