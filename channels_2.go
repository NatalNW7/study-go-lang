package main

import "time"

func main(){
	hello := make(chan string)

	go func ()  {
		time.Sleep(time.Second*2)
		hello <- "hello world"
		}()
		
		
	select{
		case x := <-hello: // caso seja atribuida algum valor de hello em x então printa x
			println(x)
		default:
			println("default")
	}
	
	println("terminou")

	queue := make(chan int)

	go func ()  {
		i := 0

		for {
			time.Sleep(time.Second)
			queue <- i // o loop só continuara se queue for esvaziado, alguem precisa ler em alguma outra thread
			i++
		}
	}()

	for x := range queue {
		println(x) // como o loop na outra thread é infinito, o valor de queue sera lido para sempre
	}
}