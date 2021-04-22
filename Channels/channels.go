package main

import "fmt"

func main() {
	channel := make(chan string)

	//la go routine no va a parar de solicitar información porque esta en un ciclo infinito
	go func(channel chan string) {
		for {
			var name string
			fmt.Scanln(&name)
			channel <- name
		}
	}(channel)

	for {
		msg := <-channel //Se detiene la ejecución debido a que se esta esperando información del canal

		fmt.Println("Este imprimiendo lo que recibi del canal " + msg)
	}

}

/*
	Los go channels nos permiten comunicar goroutines, unas con otras
	para crear un canal hacemos uso de la fn make( chan tipo_dato )
	tipo_dato es los datos que voy a transferir entre canales

	La forma en que enviamos información a un canal es utilizando la fecha <-
	canal <- información a enviar

	Si estamos recibiendo el canal aparece del lado derecho
	variable <- canal
*/
