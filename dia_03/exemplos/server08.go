package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"
)

type cep struct {
	Cep        string `json:"cep"`
	Cidade     string `json:"cidade,omitempty"`
	Bairro     string `json:"bairro,omitempty"`
	Logradouro string `json:"logradouro,omitempty"`
	UF         string `json:"uf,omitempty"`
}

func (c cep) exist() bool {
	return len(c.UF) != 0
}

var endpoints = map[string]string{
	"viacep":           "https://viacep.com.br/ws/%s/json/",
	"postmon":          "https://api.postmon.com.br/v1/cep/%s",
	"republicavirtual": "https://republicavirtual.com.br/web_cep.php?cep=%s&formato=json",
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Bem vindo a API de CEPs"))
}

func cepHandler(w http.ResponseWriter, r *http.Request) {
	// Restrigindo o acesso apenas pelo método GET
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	rCep := r.URL.Path[len("/cep/"):]
	// Validando o CEP
	rCep, err := sanitizeCEP(rCep)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ch := make(chan []byte, 1)
	for src, url := range endpoints {
		endpoint := fmt.Sprintf(url, rCep)
		go request(endpoint, src, rCep, ch)
	}

	w.Header().Set("Content-Type", "application/json")
	for index := 0; index < 3; index++ {
		log.Println(index)
		cepInfo, err := parseResponse(<-ch)
		if err != nil {
			continue
		}

		if cepInfo.exist() {
			cepInfo.Cep = rCep
			json.NewEncoder(w).Encode(cepInfo)
			return
		}
	}

	http.Error(w, http.StatusText(http.StatusNoContent), http.StatusNoContent)
}

func request(endpoint, src, cep string, ch chan []byte) {
	start := time.Now()

	c := http.Client{Timeout: time.Duration(time.Millisecond * 300)}
	resp, err := c.Get(endpoint)
	if err != nil {
		log.Printf("Ops! ocorreu um erro: %s", err.Error())
		ch <- nil
		return
	}
	defer resp.Body.Close()

	requestContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Ops! ocorreu um erro: %s", err.Error())
		ch <- nil
		return
	}

	if len(requestContent) != 0 && resp.StatusCode == http.StatusOK {
		log.Printf("O endpoint respondeu com sucesso - source: %s, CEP: %s, Duração: %s", src, cep, time.Since(start).String())
		ch <- requestContent
	}
}

func parseResponse(content []byte) (payload cep, err error) {
	response := make(map[string]interface{})
	_ = json.Unmarshal(content, &response)

	if err := isValidResponse(response); !err {
		return payload, errors.New("invalid response")
	}

	if _, ok := response["localidade"]; ok {
		payload.Cidade = response["localidade"].(string)
	} else {
		payload.Cidade = response["cidade"].(string)
	}

	if _, ok := response["estado"]; ok {
		payload.UF = response["estado"].(string)
	} else {
		payload.UF = response["uf"].(string)
	}

	if _, ok := response["logradouro"]; ok {
		payload.Logradouro = response["logradouro"].(string)
	}

	if _, ok := response["tipo_logradouro"]; ok {
		payload.Logradouro = response["tipo_logradouro"].(string) + " " + payload.Logradouro
	}

	payload.Bairro = response["bairro"].(string)

	return
}

func isValidResponse(requestContent map[string]interface{}) bool {
	if len(requestContent) <= 0 {
		return false
	}

	if _, ok := requestContent["erro"]; ok {
		return false
	}

	if _, ok := requestContent["fail"]; ok {
		return false
	}

	return true
}

// Função para validar o CEP
func sanitizeCEP(cep string) (string, error) {
	re := regexp.MustCompile(`[^0-9]`)
	sanitizedCEP := re.ReplaceAllString(cep, `$1`)

	if len(sanitizedCEP) < 8 {
		return "", errors.New("O CEP deve conter apenas números e no minimo 8 digitos")
	}

	return sanitizedCEP[:8], nil
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/cep/", cepHandler)

	log.Println("Iniciando o servidor na porta: 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
