package main

import "fmt"

func main() {
	slice := make([]int, 0, 4) //tipo, longitud, capacidad
	slice = append(slice, 2)   //agrego un elemento al slice y aumenta la capacidad si lo requiere
	fmt.Println(slice)         //Capacidad del slice
}

/*
	C/uno de los slices apuntan a un array interno
	longitud: cual es la longitud del arreglo interno len()
	capacidad: cuantos elementos caben en el slice cap()
*/
