package main

import "time"

func worker(workerId int, data chan int) {
	for i := range data {
		println("Worker: ", workerId, "Processando: ", i)
		time.Sleep(time.Second)
	}
}

// func main é a thread 1
func main() {
	channel := make(chan string) // Criando um canal para comunicação entre goroutines
	newChannel := make(chan int)

	go func() { // Criando uma goroutine, thread 2
		channel <- "Hello, World!" // Enviando uma mensagem para a thread 1
	}()

	println(<-channel) // Recebendo a mensagem da thread 2

	go worker(1, newChannel) // Criando uma goroutine, thread 3
	go worker(2, newChannel) // Criando uma goroutine, thread 4

	for i := 0; i < 10; i++ {
		newChannel <- i
	}
}
