package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	executeReadFile()
	fmt.Println("Nunca me voy a imprimir")
}

func executeReadFile() {
	ejecucion := readFile()
	fmt.Println(ejecucion)
}

/*
	La forma de evitar que haya algún error que termine abruptamente la ejecución de la fn
	o haya un return que nos impida llegar al punto que nos impida cerrar el archivo, es utilizar defer
*/
// Una forma de detener el panic es usando recover
func readFile() bool {
	archivo, err := os.Open("./holaa.txt")

	//Se agrega esta fn a un stack que luego de ejecutar nuestra fn se ejecute esta función
	defer func() {
		archivo.Close()
		fmt.Println("Defer")

		r := recover()
		fmt.Println(r)
	}()

	if err != nil {
		//Forma en que nosotros podemos imprimir un error con el stacktrace
		panic(err)
		// fmt.Println(err)
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
