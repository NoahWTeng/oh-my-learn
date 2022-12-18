package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  //analizar argumentos , tiene que llamar a esto por su cuenta
	fmt.Println(r.Form)  //imprime información en el form en el lado del servidor
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") // enviar datos al lado del cliente
}

func main() {
	http.HandleFunc("/", sayhelloName) // define la ruta
	err := http.ListenAndServe(":9090", nil) //  establece el puerto de escucha
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}