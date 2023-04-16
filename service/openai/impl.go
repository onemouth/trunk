package openai

import "net/http"

type Impl struct {
	apiKey     string
	httpClient http.Client
}

func New(apiKey string, httpClient http.Client) *Impl {
	return &Impl{apiKey: apiKey, httpClient: httpClient}
}
