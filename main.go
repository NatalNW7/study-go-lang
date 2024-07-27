package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func soma(a int, b int) (int, bool) {
	if a > 10 {
		return a + b, true
	}
	return a + b, false
}

type Course struct {
	Name  string `json:"name"` // Tag para definir o nome do campo no JSON
	Desc  string `json:"desc"`
	Price int    `json:"price"`
}

// metodos iniciando com a letra Maiuscula são públicos, com a letra minuscula são privados
func (c Course) GetInfo() string { // Método para retornar informações do curso
	return fmt.Sprintf("Curso: %s, Descrição: %s, Preço: %d", c.Name, c.Desc, c.Price)
}

func counter() {
	for i := 0; i < 10; i++ {
		println(i)
		time.Sleep(time.Second)
	}
}

func main() {
	a := "Hello, World!"
	println(a)

	resultado, status := soma(1, 2)

	println(resultado, status)

	course := Course{
		Name:  "Go",
		Desc:  "Curso de Go",
		Price: 100,
	}

	println(course.GetInfo())
	println(course.Name, course.Desc, course.Price)

	go counter() // Executando a função em uma goroutine, cria uma thread separada
	go counter()
	counter() // Executando a função na thread principal

	http.HandleFunc("/", home)
	http.HandleFunc("/curso", curso)
	http.ListenAndServe(":8182", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	println("Acessou a home")
	w.Write([]byte("Hello, World!"))
}

func curso(w http.ResponseWriter, r *http.Request) {
	course := Course{
		Name:  "Go",
		Desc:  "Curso de Go",
		Price: 100,
	}
	// res, _ = json.Marshal(course) // Convertendo para JSON, o _ ignorará o erro
	json.NewEncoder(w).Encode(course)
}
