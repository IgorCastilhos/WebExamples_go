package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Um middleware pega um http.HandlerFunc como um dos seus parâmetros, engloba ele e retorna um novo.
// Aqui definimos um novo type Middleware que faz ser mais fácil de juntar múltiplos middlewares juntos
// Essa ideia foi inspirada por Mat Ryers em uma conversa sobre Construindo Apis.

type Middleware func(http.HandlerFunc) http.HandlerFunc

func createNewMiddleware() Middleware {
	// Cria um novo Middleware
	middleware := func(next http.HandlerFunc) http.HandlerFunc {

		// Define o http.handlerFunc que será chamado pelo servidor
		handler := func(writer http.ResponseWriter, request *http.Request) {

			// ... middleware faz alguma coisa

			// Chama o próximo middleware/handler na linha
			next(writer, request)
		}

		// Retorna o novo handler criado
		return handler
	}

	// Retorna o novo middleware criado
	return middleware
}

// Logging - Faz o log de todas as requests com o caminho e o tempo levado para processar
func Logging() Middleware {

	// Cria um novo middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define o http.HandlerFunc
		return func(writer http.ResponseWriter, request *http.Request) {

			// Faz coisas de middleware
			start := time.Now()
			defer func() { log.Println(request.URL.Path, time.Since(start)) }()

			// Chama o próximo middleware/handler na linha
			f(writer, request)
		}
	}
}

// Method garante que a URL só pode ser requisitada com um método específico, caso contrário retorna 400 Bad Request
func Method(m string) Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			if request.Method != m {
				http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			f(writer, request)
		}
	}
}

func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func Hello(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(writer, "Hello World")
}
func main() {
	http.HandleFunc("/", Chain(Hello, Method("GET"), Logging()))
	http.ListenAndServe(":8080", nil)
}
