package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ffsales/desafio-multithreading/configs"
	"github.com/ffsales/desafio-multithreading/internal/infra/dto"
)

func GetViaCep(config *configs.Conf, cep string) dto.OutputViaCep{

	ctx := context.Background()

	url := fmt.Sprintf(config.UrlViaCep, cep)
	println(url)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		panic(err)
		// return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
		// return nil, err
	}
	defer resp.Body.Close()

	body, error := io.ReadAll(resp.Body)
	if error != nil {
		panic(err)
		// return nil, err
	}

	var output dto.OutputViaCep
	error = json.Unmarshal(body, &output)
	if error != nil {
		panic(err)
		// return nil, err
	}

	return output
}