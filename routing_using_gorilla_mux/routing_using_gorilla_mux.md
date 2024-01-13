### Routing using Gorilla Mux

* O pacote net/http do Go fornece muitas funcionalidades para o protocolo HTTP.
No entanto, uma coisa que ele não faz muito bem é roteamento de solicitações
complexas, como segmentar uma URL de solicitação em parâmetros individuais.
* Felizmente, há um pacote muito popular para isso, conhecido por sua boa
qualidade de código na comunidade Go. Neste exemplo, você verá como usar o
pacote gorilla/mux para criar rotas com parâmetros nomeados, manipuladores
GET/POST e restrições de domínio.
* Para instalar o pacote que adapta o roteador HTTP padrão do GO



* `go get -u github.com/gorilla/mux`