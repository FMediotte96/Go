package main

import (
	"encoding/json"
	"net/http"
)

//si el primer caracter esta en minuscula significa private y el mayus que es publico
//por lo que go si son privados no los va a mostrar en el json
type Curso struct {
	Title        string `json:"titulo"`
	NumeroVideos int    `json:"numero_videos"`
}

type Cursos []Curso

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cursos := Cursos{
			Curso{"Curso de Go", 30},
			Curso{"Curso de Ruby", 20},
			Curso{"Curso de NodeJS", 40},
		}
		json.NewEncoder(w).Encode(cursos)
	})
	http.ListenAndServe(":8000", nil)
}

/*
	Go viene con una libreria para poder hacer render de json, a partir de estructuras
*/
