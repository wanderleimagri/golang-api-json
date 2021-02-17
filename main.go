package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Pessoa struct {
	Id    int
	Nome  string
	Email string
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
	json.NewEncoder(w).Encode(Pessoas)
}

func configurarRotas() {
	http.HandleFunc("/", rotaPrincipal)
	http.HandleFunc("/pessoas", listarPessoas)
}

func configurarServidor() {
	configurarRotas()
	fmt.Println("Servidor rodando em localhost:8888")
	http.ListenAndServe(":8888", nil)
}

func main() {
	configurarServidor()
}
