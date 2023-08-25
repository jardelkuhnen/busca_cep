package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	brasilAPi := make(chan string)
	viaCep := make(chan string)
	cep := "85811050"
	go doRequest(brasilAPi, cep, "https://brasilapi.com.br/api/cep/v1/%s")
	go doRequest(viaCep, cep, "https://viacep.com.br/ws/%s/json/")

	select {
	case resp := <-brasilAPi:
		println("Response from brasilapi:")
		println(resp)
	case resp := <-viaCep:
		println("Response from viacep: ")
		println(resp)
	case <-time.After(1 * time.Second):
		println("Timeout")
	}
}

func doRequest(result chan<- string, cep, url string) {
	resp, err := http.Get(fmt.Sprintf(url, cep))
	if err != nil {
		log.Print("error getting response from api: " + err.Error())
		return
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	result <- string(responseBody)
}
