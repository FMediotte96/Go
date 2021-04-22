package main

import (
	"fmt"
	"strings"
	"time"
)

/*
	Go no utiliza los threads del sistema sino que utiliza algo más ligero llamado goroutines,
	no teniendo limites puedo tener las que yo necesite.
	Cuando multiples goroutines se bloquean, Go crea una nueva si lo necesita y luego las elimina
	No hace falta que sea una función definida, podemos crear una fn anonima
*/
func main() {
	go mi_nombre_lento("Facundo")
	fmt.Println("Que aburridoo") //esto se ejecuta dsp porque es sincronico

	var wait string
	fmt.Scanln(&wait) //Espera a que yo tipie algo para finalizar

	//podemos hacerlo concurrente anteponiendo la palabra go, la palabra go movio la ejecución a
	//una go routine, es decir a una ejecución separada

}

/*
	GoRoutines -> concurrencia

	librerias que voy a necesitar:
	time
	strings
*/

func mi_nombre_lento(name string) {
	letras := strings.Split(name, "")

	for _, letra := range letras { //_ es el index
		time.Sleep(1000 * time.Millisecond) //Duerme
		fmt.Println(letra)
	}

}
