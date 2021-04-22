package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	copia := make([]int, len(slice))
	//importante el size, lo mejor es en vez de pasarle un dato static
	//sería pasarle el len del slice fuente.

	//copy(destino, fuente)
	//es importante que el array destino tenga la misma longitud que la fuente ya que copia el mínimo de elementos
	//en ambos arrays
	copy(copia, slice)

	fmt.Println(slice)
	fmt.Println(copia)
}
