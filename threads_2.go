package main

import "time"


func contador(tipo string){
	for i:= 0; i < 5; i++{
		println(tipo, i)
		time.Sleep(time.Second)
	}
}

func main(){
	// contador("sem go routine")
	// go contador("com go routine")
	// println("hello 1")
	// println("hello 2")
	// time.Sleep(time.Second)
	
	go contador("a")
	go contador("b")
	time.Sleep(time.Second*10)
}