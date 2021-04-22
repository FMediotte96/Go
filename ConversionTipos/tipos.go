package main

import (
	"fmt"
	"strconv"
)

func main() {
	edad := 25

	fmt.Println(edad)

	edad_str := strconv.Itoa(edad)

	fmt.Println("Mi edad es " + edad_str)

	//fmt.Println(edad + 10)

	//Atoi me devuelve un par int, error, si queremos obviar esta última utilizamos el operador "_" que desecha la variable recibida
	edad_int, _ := strconv.Atoi(edad_str)

	fmt.Println(edad_int + 10)

}

//El compilador no nos permite importar una libreria y no utilizarla,
//como así tampoco tener variables declaradas sin utilizar
