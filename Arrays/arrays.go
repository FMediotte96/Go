package main

import "fmt"

func main() {
	array := [3]int{1, 2}

	array[2] = 20

	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}

	var matriz [3][2]int

	matriz[0][1] = 1
	fmt.Println(matriz)
}

/*
	Los arreglos tienen un único tipo en go y se inicializan por defecto con null, vacio o null según el tipo de dato
	para instanciar un arreglo ponemos {} luego de tipo de datos
*/
