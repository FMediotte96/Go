package main

import "net/http"

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		// http.ServeFile(w, r, "index.html") //relativo al path de donde lo ejecuto
// 		http.ServeFile(w, r, r.URL.Path[1:])
// 		//r.URL.Path nos disponibiliza todo lo que esta después de la diagonal via el objeto request
// 		//toma del caracter 1 en adelante lo cual es peligroso porque expone nuestro código y expone a algún hackeo
// 	})
// 	http.ListenAndServe(":8000", nil)
// }

/*
	Para evitar este tipo de hacks lo que hacemos es crear un folder público y levantar o servir
	los mismos se utiliza
*/
func main() {
	fileServer := http.FileServer(http.Dir("public"))

	//A handle le decimos el prefijo al cual tiene que acceder y un handler
	http.Handle("/public/", http.StripPrefix("/public/", fileServer))
	// http.Handle("/", http.StripPrefix("/", fileServer))

	http.ListenAndServe(":8000", nil)
}
