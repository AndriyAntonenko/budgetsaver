package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CurracyExchangeService struct {
	client *http.Client
	apiKey string
	url    string
}

type AvailableSymbolsResponse struct {
	Success bool              `json:"success"`
	Symbols map[string]string `json:"symbols"`
}

func NewCurracyExchangeService(apiKey string, url string) *CurracyExchangeService {
	return &CurracyExchangeService{
		client: &http.Client{},
		apiKey: apiKey,
		url:    url,
	}
}

func (s *CurracyExchangeService) buildUrl(path string) string {
	return fmt.Sprintf("%s%s?access_key=%s", s.url, path, s.apiKey)
}

func (s *CurracyExchangeService) GetSupportedSymbols() (map[string]string, error) {
	request, err := http.NewRequest("GET", s.buildUrl("/symbols"), nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	response, err := s.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	var symbols AvailableSymbolsResponse
	err = json.Unmarshal(body, &symbols)

	if err != nil {
		return nil, err
	}

	return symbols.Symbols, nil
}
