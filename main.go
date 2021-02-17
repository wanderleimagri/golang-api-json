package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
		Id: 3, Nome: "Joao", Email: "joao@gmail.com",
	},
}

func rotaPrincipal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem vindo")
}

func listarPessoas(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Pessoas)
}

func cadastrarPessoa(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}
}

func configurarRotas() {
	http.HandleFunc("/", rotaPrincipal)
	http.HandleFunc("/pessoas", listarPessoas)
	http.HandleFunc("/pessoas", cadastrarPessoa)
}

func configurarServidor() {
	configurarRotas()

	fmt.Println("Servidor rodando em localhost:8888")
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func main() {
	configurarServidor()
}
