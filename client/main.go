package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main(){

	// se a chamda demorar mais q 3 segundos a requisição é cancelada
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	// preparando requisição
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8182", nil)
	if err != nil {
		log.Fatalf("erro ao criar requisição %v", err)
	}

	// executando requisição
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Ocorreu um erro %v", err)
	}

	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)
}