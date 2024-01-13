package main

import (
	"fmt"
	"net/http"
)

func main() {
	// O pacote net/http contém todas as utilidades necessárias para aceitar solicitações e manipulá-las dinamicamente.
	// Podemos registrar um novo manipulador (handler) com a função http.HandleFunc. Seu primeiro parâmetro é o caminho
	// a ser correspondido, e o segundo é a função a ser executada. Neste exemplo: quando alguém acessa seu site
	// (http://example.com/), ele ou ela será saudado com uma mensagem agradável.
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "Bem-vindo(a) ao meu website!")
		if err != nil {
			return
		}
	})
	// Para lidar com o aspecto dinâmico, o http.Request contém todas as informações sobre a solicitação e seus
	// parâmetros. Você pode ler os parâmetros GET usando r.URL.Query().Get("token") ou os parâmetros POST
	// (campos de um formulário HTML) usando r.FormValue("email").

	// Servir ativos estáticos, como JavaScript, CSS e imagens, é feito utilizando o http.FileServer,
	// um recurso incorporado no pacote net/http. Para configurar o servidor de arquivos estáticos corretamente,
	// é necessário indicar o diretório onde os arquivos devem ser servidos. Isso pode ser feito da seguinte maneira:
	fs := http.FileServer(http.Dir("static/"))

	//	Uma vez que o servidor de arquivos estáticos está configurado, você só precisa associar um caminho de URL a ele,
	//	da mesma forma que fez com as solicitações dinâmicas. No entanto, é importante observar que, para servir
	//	arquivos corretamente, você pode precisar remover uma parte do caminho da URL. Normalmente,
	//	isso é feito com o nome do diretório onde seus arquivos estão localizados.
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Aceitar conexões é a última etapa para concluir nosso servidor HTTP básico.
	// Como você pode imaginar, Go também possui um servidor HTTP incorporado, que podemos iniciar bastante rapidamente.
	// Uma vez iniciado, você pode visualizar seu servidor HTTP em seu navegador.
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		return
	}
}
