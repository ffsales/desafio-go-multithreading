package thread

import (
	"errors"
	"time"

	"github.com/ffsales/desafio-multithreading/configs"
	"github.com/ffsales/desafio-multithreading/internal/infra/dto"
	"github.com/ffsales/desafio-multithreading/internal/infra/webclient/client"
)

func TriggerThread(config *configs.Conf, cep string) (*dto.OutputResponse, error) {

	apiCepChain := make(chan dto.OutputApiCep)
	viaCepChain := make(chan dto.OutputViaCep)
	var responseCep dto.OutputResponse = dto.OutputResponse{}

	go func() {
		responseApiCep, err := client.GetApiCep(config, cep)
		if err != nil {
			panic(err)
		}
		apiCepChain <- *responseApiCep
	}()
	go func() {
		responseViaCep, err := client.GetViaCep(config, cep)
		if err != nil {
			panic(err)
		}
		viaCepChain <- *responseViaCep
	}()

	select {
	case responseCep.ApiCep = <-apiCepChain:
		responseCep.Api = "ApiCep"
	case responseCep.ViaCep = <-viaCepChain:
		responseCep.Api = "ViaCep"
	case <-time.After(time.Second):
		return nil, errors.New("timeout")
	}

	return &responseCep, nil
}
