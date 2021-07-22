package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Bem vindo a API de CEPs"))
}

func cepHandler(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Path[len("/cep/"):]
	w.Write([]byte(cep))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/cep/", cepHandler)

	log.Println("Iniciando o servidor na porta: 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
