package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	ejecucion := readFile()
	fmt.Println(ejecucion)
}

/*
	La forma de evitar que haya algún error que termine abruptamente la ejecución de la fn
	o haya un return que nos impida llegar al punto que nos impida cerrar el archivo, es utilizar defer
*/
func readFile() bool {
	archivo, err := os.Open("./hola.txt")

	if err != nil {
		fmt.Println("Hubo error")
	}

	//Se agrega esta fn a un stack que luego de ejecutar nuestra fn se ejecute esta función
	defer func() {
		archivo.Close()
		fmt.Println("Defer")
	}()

	scanner := bufio.NewScanner(archivo)

	var i int
	for scanner.Scan() { //lee linea a linea
		i++
		linea := scanner.Text()
		fmt.Println(i)
		fmt.Println(linea)
	}

	fmt.Println("Nunca me ejecuto")

	archivo.Close() //Debemos siempre cerrar el archivo, pero puede pasar que no llegue nunca a esta linea
	return true
}
