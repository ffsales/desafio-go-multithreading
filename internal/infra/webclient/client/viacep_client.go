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

func GetViaCep(config *configs.Conf, cep string) (*dto.OutputViaCep, error) {

	ctx := context.Background()

	url := fmt.Sprintf(config.UrlViaCep, cep)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var output dto.OutputViaCep
	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, err
	}

	return &output, nil
}
