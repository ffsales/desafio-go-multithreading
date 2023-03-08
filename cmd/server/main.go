package main

import (
	"github.com/ffsales/desafio-multithreading/configs"
	"github.com/ffsales/desafio-multithreading/internal/infra/thread"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	thread.TriggerThread(configs, "08485-310")
}
