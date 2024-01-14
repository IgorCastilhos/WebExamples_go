package main

import "net/http"

// Como servir arquivos estáticos como CSS, JavaScript ou imagens de um diretório específico

func main() {
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}
