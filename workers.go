package main

import "time"


func worker(workerID int, msg chan int){
	for res := range msg {
		println("WorkerID:", workerID, " Msg:", res)
		time.Sleep(time.Second)
	}
}

func main(){
	msg := make(chan int)

	go worker(1, msg)
	go worker(2, msg)

	for i := 0; i < 10; i++ {
		msg <- i
	}
}