package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Pessoa struct {
	Id    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

var Pessoas []Pessoa = []Pessoa{
	Pessoa{
		Id: 1, Nome: "Wanderlei", Email: "wanderleimagri@gmail.com",
	},
	Pessoa{
		Id: 2, Nome: "Jose", Email: "jose@gmail.com",
	},
	Pessoa{
		Id: 3, Nome: "Joaozinho", Email: "joao@gmail.com",
	},
}

func rotaPrincipal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem vindo")
}

func listarPessoas(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Pessoas)
}

func cadastrarPessoa(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var novaPessoa Pessoa
	json.Unmarshal(body, &novaPessoa)
	novaPessoa.Id = len(Pessoas) + 1
	Pessoas = append(Pessoas, novaPessoa)
	json.NewEncoder(w).Encode(novaPessoa)

}

func rotearPessoas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	partes := strings.Split(r.URL.Path, "/")
	if len(partes) == 2 || len(partes) == 3 && partes[2] == "" {
		if r.Method == "GET" {
			listarPessoas(w, r)
		} else if r.Method == "POST" {
			cadastrarPessoa(w, r)
		}
	} else if len(partes) == 3 || len(partes) == 4 && partes[3] == "" {
		if r.Method == "GET" {
			buscarPessoa(w, r)
		} else if r.Method == "DELETE" {
			excluirPessoa(w, r)
		} else if r.Method == "PUT" {
			alterarPessoa(w, r)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func buscarPessoa(w http.ResponseWriter, r *http.Request) {

	partes := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(partes[2])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, pessoa := range Pessoas {
		if pessoa.Id == id {
			json.NewEncoder(w).Encode(pessoa)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func alterarPessoa(w http.ResponseWriter, r *http.Request) {

	partes := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(partes[2])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	indicePessoa := -1
	for indice, pessoa := range Pessoas {
		if pessoa.Id == id {
			indicePessoa = indice
			break
		}
	}
	if indicePessoa < 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var pessoaModificada Pessoa
	erroJson := json.Unmarshal(body, &pessoaModificada)
	if erroJson != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	pessoaModificada.Id = id
	Pessoas[indicePessoa] = pessoaModificada
	json.NewEncoder(w).Encode(pessoaModificada)
}

func excluirPessoa(w http.ResponseWriter, r *http.Request) {

	partes := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(partes[2])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	indicePessoa := -1
	for indice, pessoa := range Pessoas {
		if pessoa.Id == id {
			indicePessoa = indice
			break
		}
	}
	if indicePessoa < 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	Pessoas = append(Pessoas[0:indicePessoa], Pessoas[indicePessoa+1:len(Pessoas)]...)
	w.WriteHeader(http.StatusNoContent)
}

func configurarRotas() {
	http.HandleFunc("/", rotaPrincipal)
	http.HandleFunc("/pessoas", rotearPessoas)
	http.HandleFunc("/pessoas/", rotearPessoas)
}

func configurarServidor() {
	configurarRotas()

	fmt.Println("Servidor rodando em localhost:8888")
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func main() {
	configurarServidor()
}
