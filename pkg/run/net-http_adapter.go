package run

import (
	"fmt"
	"github.com/Mona-bele/action-integration-terraform/pkg/utils"
	"io"
	"net/http"
)

type NetHttpAdapter struct {
	config *utils.ConfigEnv
}

func (n *NetHttpAdapter) NewRequest(method string, url string, bodyData io.Reader) ([]byte, error) {
	client := &http.Client{}

	uri := fmt.Sprintf("https://app.terraform.io%s", url)

	req, err := http.NewRequest(method, uri, bodyData)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+n.config.TFAPITOKEN)
	req.Header.Add("Content-Type", "application/vnd.api+json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func NewNetHttpAdapter(config *utils.ConfigEnv) HttpClient {
	return &NetHttpAdapter{
		config: config,
	}
}
