package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/holamundo", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Holaaaa")
	})
	http.HandleFunc("/", handler)     //Lo toma como un handler global
	http.ListenAndServe(":8000", nil) //el segundo parametro es para routing
}

//Siempre estos dos parametros para poder responder a una petición web
//ResponseWriter nos permite definir como vamos a responder una petición
//*http.Request es un puntero a la información de la petición que nos envio el navegador
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hay una nueva petición")
	io.WriteString(w, "Hola Mundo")
}
