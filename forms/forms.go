package main

import (
	"html/template"
	"net/http"
)

// Como simular um formulário de contato e analisar ou interpretar esses dados

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func main() {
	tmpl := template.Must(template.ParseFiles("forms/forms.html"))

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPost {
			tmpl.Execute(writer, nil)
			return
		}

		details := ContactDetails{
			Email:   request.FormValue("email"),
			Subject: request.FormValue("subject"),
			Message: request.FormValue("message"),
		}

		//	 Faz alguma coisa com os detalhes
		_ = details

		tmpl.Execute(writer, struct {
			Success bool
		}{true})
	})
	http.ListenAndServe(":8080", nil)
}
