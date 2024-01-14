package main

import (
	"fmt"
	"log"
	"net/http"
)

// Esse exemplo mostra como criar um middleware básico de Login em Go.
// Um middleware simplesmente pega um http.HandlerFunc como um dos seus parâmetros, engloba ele e retorna um novo
// http.HandlerFunc para o servidor chamar.

// loggin é um middleware que faz um wrap com o HandlerFunc recebido por parâmetro e retorna um novo
func loggin(f http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request.URL.Path)
		f(writer, request)
	}
}

func foo(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "foo")
}

func bar(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "bar")
}

func main() {
	http.HandleFunc("/foo", loggin(foo))
	http.HandleFunc("/bar", loggin(bar))

	http.ListenAndServe(":8080", nil)
}
