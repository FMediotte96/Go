package main

import "fmt"

func main() {

	/*
		Punteros:

		1. Es una dirección de memoria
		2. En lugar del valor, tenemos la dirección en la que está el valor
		3. X, Y => as123d => 5
		4. X => as123d => 6 se actualiza el valor
		5. Y ¿? => 6 por lo que Y vale 6 tambien

		*T => *int, *string, *Struct (T tipo de dato)
		Valor zero de un puntero es => nil

		& -> accede a la dirección de memoria
		* -> nos da acceso al valor de la dirección de memoria
	*/

	var x, y *int
	entero := 5

	x = &entero //& accede a la dirección de memoria
	y = &entero

	*x = 6 //aca alteramos el valor de la dirección de memoria

	fmt.Println(*x)
	fmt.Println(*y)
}
