package main

func main() {
	// memoria(50) <---- a <---- 50
	a := 50

	println(&a) // o & mostra o endereço da memoria

	var ponteiro *int = &a // Está apontando para o endereço de memoria da variavel a
	println(ponteiro)
	println(*ponteiro) // mostra o valor alocado no endereço da memoria q o ponteiro está apontando
	
	*ponteiro = 10
	println(ponteiro)
	println(*ponteiro)
	println(a)

	variavel := 10
	abc(&variavel)
	println(variavel)
}

func abc(a *int) { // o * significa q vamos receber um endereço de memoria como parametro
	*a = 200
}