package main

import (
	"fmt"

	"github.com/ffsales/desafio-multithreading/configs"
	"github.com/ffsales/desafio-multithreading/internal/infra/thread"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	if response, err := thread.TriggerThread(configs, "08485-310"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.Api)
		fmt.Println(response.CepResponse)
	}
}
