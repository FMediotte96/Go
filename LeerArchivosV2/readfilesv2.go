package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readFile()
}

func readFile() bool {
	archivo, err := os.Open("./hola.txt")

	if err != nil {
		fmt.Println("Hubo error")
	}

	scanner := bufio.NewScanner(archivo)

	var i int
	for scanner.Scan() { //lee linea a linea
		i++
		linea := scanner.Text()
		fmt.Println(i)
		fmt.Println(linea)
	}

	archivo.Close() //Debemos siempre cerrar el archivo, pero puede pasar que no llegue nunca a esta linea
	return true
}
