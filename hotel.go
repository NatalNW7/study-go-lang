package main

import (
	"context"
	"time"
)

func main(){
	ctxPai := context.Background() // contexto pai/raiz
	ctx, cancel := context.WithCancel(ctxPai)
	defer cancel() // defer espera tudo ser executado para depois executar, ou seja, é executado por ultimo
	
	go func ()  {
		time.Sleep(time.Second * 3) // se a tarefa q for executada aqui terminar antes do bookHotel(), o bookHotel, para imediatamente
		// aqui poderia ser uma outra tarefa para agendar quarto, então teriamos duas tarefas simultaneas para agendamento de quarto
		cancel()
	}()
	
	bookHotel(ctx)
	
	ctx, newCancel := context.WithTimeout(ctxPai, time.Second*10)
	defer newCancel()
	bookHotel(ctx)

}

func bookHotel(ctx context.Context){
	select { // select aguarda algum condição dar certo
	case <- ctx.Done():
		println("Tempo execido para agendar quarto")
	case <- time.After(time.Second * 5):
		println("Quarto reservado!")
	}
}