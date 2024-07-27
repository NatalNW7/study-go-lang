package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Cliente struct {
	Nome  string
	Email string
	CPF   int // o lula falou q é string, fal o L hahaha
}

func (c Cliente) Exibe() {
	println("Exibindo cliente: ", c.Nome)
}

type ClienteInternacional struct {
	Cliente // composição/herança
	Pais  string `json:"pais"` // usando tag para tratar a saida json
}

func main(){
	cliente := Cliente{
		Nome: "Jão da Silva",
		Email: "jao@jao.com",
		CPF: 12345678900,
	}
	fmt.Println(cliente)

	cliente2 := Cliente{"Maria", "m@m.com", 11987654321}
	fmt.Printf("Nome: %s, Email: %s, CPF: %d\n", cliente2.Nome, cliente2.Email, cliente2.CPF)
	
	cliente3 := ClienteInternacional{
		Cliente: Cliente{
			Nome: "Enzo",
			Email: "enzopjl@enzo.com",
			CPF: 123123123123,
		},
		Pais: "EUA",
	}
	fmt.Printf("Nome: %s, Email: %s, CPF: %d, Pais: %s\n", cliente3.Nome, cliente3.Email, cliente3.CPF, cliente3.Pais)

	cliente3.Exibe()

	clienteJson, err := json.Marshal((cliente3))
	if err != nil {
		log.Fatal(err.Error())
	}

	println(string(clienteJson))

	jsonCliente4 := `{"Nome":"Enzo","Email":"enzopjl@enzo.com","CPF":123123123123,"pais":"EUA"}`
	cliente4 := ClienteInternacional{}
	json.Unmarshal([]byte(jsonCliente4), &cliente4)
	fmt.Println(cliente4)
}