package main

import "fmt"

func main() {
	// slice := []int{1, 2, 3, 4}
	slice := []int{1, 2}
	if slice == nil {
		fmt.Println("Esta vació")
	} else {
		fmt.Println(len(slice))
	}

	array := [3]int{1, 2, 3}
	slice2 := array[0:] //Slicing (desde:hasta) no inclusivos
	fmt.Println(slice2)

}

/*
	Slices: estructura que tiene como base los arrays, pero en este caso tenemos una dimensión dinamica
	Una diferencia con los arrays es que cuando tenemos un slice no inicializado es que su valor es null

	Estructura de un slice:
	-Longitud del array
	-Puntero al array
	-Capacidad

	Podemos obtener un slice a partir de un arreglo utilizando el concepto llamdo Slicing
*/
