package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type cep struct {
	Cep        string `json:"cep"`
	Cidade     string `json:"cidade"`
	Bairro     string `json:"bairro"`
	Logradouro string `json:"logradouro"`
	UF         string `json:"uf"`
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Bem vindo a API de CEPs"))
}

func cepHandler(w http.ResponseWriter, r *http.Request) {
	rCep := r.URL.Path[len("/cep/"):]
	c := cep{Cep: rCep}
	ret, err := json.Marshal(c)
	if err != nil {
		log.Printf("Ops! ocorreu um erro: %s", err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	w.Write([]byte(ret))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/cep/", cepHandler)

	log.Println("Iniciando o servidor na porta: 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
