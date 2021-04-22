package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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
}
