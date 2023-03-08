package main

import (
	"github.com/ffsales/desafio-multithreading/configs"
	"github.com/ffsales/desafio-multithreading/internal/infra/webclient/client"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	println(configs.UrlApiCep)
	println(configs.UrlViaCep)

	client.GetApiCep(configs, "08485-310")
	client.GetViaCep(configs, "08485-310")
}
