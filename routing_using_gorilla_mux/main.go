package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	// Cria um novo router de requisições. O router é o principal router da aplicação web e será passado como
	// parâmetro para o servidor. Ele irá receber todas as conexões HTTP e as passará para os manipuladores (handlers)
	//de request que tu registrar.
	router := mux.NewRouter()

	// Registrando um Request Handler
	// Uma vez criando um novo router, tu pode registrar manipuladores de requisições normalmente. A única diferença
	// é que, ao invés de chamar http.HandleFunc(...), tu chama HandleFunc no seu router, assim:
	// ***          router.HandleFunc()          ***

	// A maior força do Router gorilla/mux é a habilidade de extrair segmentos da URL da Request. Como um exemplo,
	// essa é uma URL da aplicação:
	// /books/go-programming-blueprint/page/10
	// Essa URL tem dois segmentos dinâmicos:
	// Um slug, que é uma versão simplificada e amigável para URL de um determinado texto, do título (go-programming)
	// Page(10)
	// Para ter um request handler correspondendo a URL mencionada acima tu precisa recolocar os segmentos dinâmicos
	// dos placeholders no seu padrão de URL:
	router.HandleFunc("/books/{title}/page/{page}", func(writer http.ResponseWriter, request *http.Request) {
		// Paga pegar os dados desses segmentos, o pacote vem com a função mux.Vars(r) que pega o http.Request como
		// parâmetro e retorna um map dos segmentos
		vars := mux.Vars(request)
		// title pega o título do livro
		title := vars["title"]
		// page pega a página do livro
		page := vars["page"]

		fmt.Fprintf(writer, "Tu solicitou o livro: %s na página %s\n", title, page)
	})

	// Configurando o router do servidor HTTP
	// O nil em http.ListenAndServe(":80", nil) significa o parâmetro para o router principal do servidor. Por padrão
	// ele é nil, que significa usar o router padrão do pacote net/http. Para fazê-lo usar o seu router, troque o nil com
	// a variável do seu router r.
	http.ListenAndServe(";80", router)
	// Funcionalidades do Router gorilla/mux
	// Métodos
	// Restrinja o handler de requisições para métodos HTTP específicos:
	router.HandleFunc("/books/{title}", CreateBook).Methods("POST")
	router.HandleFunc("/books/{title}", ReadBook).Methods("GET")
	router.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")

	// Hostnames & Subdomains
	// Restrinja o handler de requisições para hostnames específicos ou subdomínios
	router.HandleFunc("/books/{title}", BookHandler).Host("www.mybookstore.com")

	// Schemes
	// Restrinja o handler de requisições para http/https
	router.HandleFunc("/secure", SecureHandler).Schemes("https")
	router.HandleFunc("/insecure", InsecureHandler).Schemes("http")

	// Prefixos de Path & Subrouters
	// Restrinja o handler de requisições para caminhos específicos
	bookrouter := router.PathPrefix("/books").Subrouter()
	bookrouter.HandleFunc("/", AllBooks)
	bookrouter.HandleFunc("/{title}", GetBook)
}
