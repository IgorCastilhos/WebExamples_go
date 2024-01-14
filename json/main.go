package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Como codificar e decodificar dados em JSON usando o pacote encoding/json

type User struct {
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Age       int    `json:"age,omitempty"`
}

func main() {
	http.HandleFunc("/decode", func(writer http.ResponseWriter, request *http.Request) {
		var user User
		json.NewDecoder(request.Body).Decode(&user)

		fmt.Fprintf(writer, "%s %s tem %d anos de idade", user.Firstname, user.Lastname, user.Age)
	})

	http.HandleFunc("/encode", func(writer http.ResponseWriter, request *http.Request) {
		igor := User{
			Firstname: "Igor",
			Lastname:  "Castilhos",
			Age:       23,
		}

		json.NewEncoder(writer).Encode(igor)
	})

	http.ListenAndServe(":8080", nil)
}
