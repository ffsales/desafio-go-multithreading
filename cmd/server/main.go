package main

import "github.com/ffsales/desafio-multithreading/configs"

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	println(configs.UrlApiCep)
	println(configs.UrlViaCep)
}
