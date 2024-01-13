package main

import (
	"fmt"
	"net/http"
)

// Handler que recebe todas as conexões HTTP dos navegadores, clientes HTTP ou requisições de API.
// A função recebe dois parâmetros:
// http.ResponseWriter é onde tu escreve a resposta do tipo text/html
// http.Request contém todas as informações sobre essa solicitação HTTP, incluindo coisas como a URL ou campos de cabeçalho
func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello World! Tu requisitou : %s\n", request.URL.Path)
	})

	// http.ListenAndServe vai ficar escutando na porta 80
	http.ListenAndServe(":80", nil)
}
