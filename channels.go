package main


func main() {
	forever := make(chan string)

	go func()  {
		for {

		}
	}()

	println("aguradando para sempre")
	<-forever 

	// o channel para execução até q algum valor seja passado para ele
	// numa situação normal, o go pularia o loop infinito de forma preventiva
}