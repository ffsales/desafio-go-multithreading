package thread

import (
	"fmt"
	"time"
	"github.com/ffsales/desafio-multithreading/configs"
	"github.com/ffsales/desafio-multithreading/internal/infra/dto"
	"github.com/ffsales/desafio-multithreading/internal/infra/webclient/client"
)

func TriggerThread(config *configs.Conf, cep string) {

	apiCepChain := make(chan dto.OutputApiCep)
	viaCepChain := make(chan dto.OutputViaCep)

	go func() {
		apiCepChain <-client.GetApiCep(config, cep)
	}()
	go func() {
		viaCepChain <-client.GetViaCep(config, cep)
	}()
	
	select {
	case jsonCep := <-apiCepChain:
		fmt.Printf("received (ApiCep): %s\n", jsonCep.Address)
		break
	case jsonCep := <-viaCepChain:
		fmt.Printf("received (ViaCep): %s\n", jsonCep.Logradouro)
		break
	case <-time.After(time.Second):
		fmt.Println("timeout")
	}	
}