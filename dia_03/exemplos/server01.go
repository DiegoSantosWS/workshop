package main

import (
	"log"
	"net/http"
)

// Defina uma função manipuladora (handler) chamada home que escreve
// um slice de bytes contendo "Bem vindo a API de CEPs" no o corpo da resposta.
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bem vindo a API de CEPs"))
}

func main() {
	// Use a função http.NewServeMux() para inicializar um novo servermux,
	// depois registre a função home como manipulador do padrão de URL "/".
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// Use a função http.ListenAndServe() para iniciar um novo servidor web.
	// Passamos dois parâmetros: o endereço de rede TCP que será escutado (neste caso ": 4000")
	// e os servermux que acabamos de criar.
	// Se http.ListenAndServe() retornar um erro, usamos a função log.Fatal()
	// para registrar a mensagem de erro e sair.
	log.Println("Iniciando o servidor na porta: 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
