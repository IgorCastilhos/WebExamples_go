package main

import (
	"net/http"
	"text/template"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	tmpl := template.Must(template.ParseFiles("templates/layout.html"))
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		data := TodoPageData{
			PageTitle: "Minha lista TODO",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: false},
				{Title: "Task 3", Done: false},
			},
		}
		tmpl.Execute(writer, data)
	})
	http.ListenAndServe(":80", nil)
}
