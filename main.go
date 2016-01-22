package main

import (
	"github.com/lansanacamara/davillex/controllers"
	"net/http"
)

func main() {
	cssHandler := http.FileServer(http.Dir("./public/css/"))
	jsHandler := http.FileServer(http.Dir("./public/js/"))
	imgHandler := http.FileServer(http.Dir("./public/img/"))
	fontHandler := http.FileServer(http.Dir("./public/font-awesome/"))

	http.Handle("/css/", http.StripPrefix("/css/", cssHandler))
	http.Handle("/js/", http.StripPrefix("/js/", jsHandler))
	http.Handle("/img/", http.StripPrefix("/img/", imgHandler))
	http.Handle("/font-awesome/", http.StripPrefix("/font-awesome/", fontHandler))

	http.HandleFunc("/", controllers.Index)
	http.ListenAndServe(":80", nil)
}