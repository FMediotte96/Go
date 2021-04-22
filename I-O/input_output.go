package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	precio := 14.3
	fmt.Printf("El precio es %f\n", precio)

	// var nombre string
	// fmt.Println("Coloca tu nombre: ")
	// fmt.Scanf("%s\n", &nombre)
	// fmt.Printf("Mi nombre es %s\n", nombre)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingresa tu nombre: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Hola " + text)
	}

}

/*
	%d entero
	%v valor por default del dato
	%f flotante
	%e y %b flotante pero para números cientificos
*/
