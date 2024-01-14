package main

import (
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
)

// Como armazenar dados nos cookies de sessão usando o famoso pacote gorilla/sessions do Go.
// Cookies são pequenos pedaços de dado armazenados no navegador de um usuário, sendo enviados para o nosso servidor em
// cada requisição. Nós podemos armazená-los, estando o usuário logado ou não no nosso website, descobrindo assim quem
// ele realmente é (no nosso sistema).

// Nesse teste, será permitido somente aos usuários autenticados visualizarem a mensagem secreta na página /secret.
// Para acessá-la, primeiro tem que acessar /login para pegar um cookie de sessão, que loga o user. Adicionalmente ele
// pode visitar /logout para reivindicar o acesso dele à mensagem secreta.

var (
	// A chave deve conter 16, 24 ou 32 bytes (AES-128, AES-192 ou AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func secret(writer http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "cookie-name")

	// Verifica se o usuário está autenticado
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(writer, "Forbidden", http.StatusForbidden)
		return
	}

	fmt.Fprintln(writer, "Mensagem super secreta!!!")
}

func login(writer http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "cookie-name")
	// autenticação vai aqui
	//	 ...

	// Configura o usuário como autenticado
	session.Values["authenticated"] = true
	session.Save(request, writer)
}

func logout(writer http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "cookie-name")

	// Reivindica a autenticação do usuário
	session.Values["authenticated"] = false
	session.Save(request, writer)
}

func main() {
	http.HandleFunc("/secret", secret)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)

	http.ListenAndServe(":8080", nil)
}
