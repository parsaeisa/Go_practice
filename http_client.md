# golang_http_client

```go
package package_name

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
)

type HttpHandler struct {
	client  *http.Client
	baseURL string
}

// FailResponse is failure HTTP response
type FailResponse struct {
	Error string `json:"error"`
}

func NewHttpHandler(baseUrl string, timeout time.Duration) *HttpHandler {
	return &HttpHandler{
		client:  &http.Client{Timeout: timeout},
		baseURL: baseUrl,
	}
}

func (h HttpHandler) GetDriverRateHistory(rideIds []uint64) (*handlers.RateHistoryResponse, error) {
	request := handlers.RateHistoryRequest{
		RideIds: rideIds,
	}
	v, err := query.Values(request)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/smaple?%s", h.baseURL, strings.TrimSpace(v.Encode()))
	print(url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		var response handlers.RateHistoryResponse
		body, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()

		err = json.Unmarshal(body, &response)
		if err != nil {
			return nil, err
		}
		print(response.RidesData[0].Rate)

		return &response, nil
	}

	var failResponse FailResponse
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	err = json.Unmarshal(body, &failResponse)
	if err != nil {
		return nil, err
	}

	return nil, errors.New(failResponse.Error)
}

```