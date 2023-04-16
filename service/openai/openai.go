package openai

import (
	"context"
	"errors"
)

var ErrHTTPNot200 = errors.New("http status code not 200")

type Service interface {
	CreateEmbeddings(ctx context.Context, input string) (*EmbeddingsResp, error)
}

type EmbeddingsReq struct {
	Model string `json:"model"`
	Input string `json:"input"`
}

type Embedding struct {
	Embedding []float64 `json:"embedding"`
	Index     int       `json:"index"`
	Object    string    `json:"object"`
}

type Usage struct {
	PromptTokens int `json:"prompt_tokens"`
	TotalTokens  int `json:"total_tokens"`
}

type EmbeddingsResp struct {
	Object     string      `json:"object"`
	Embeddings []Embedding `json:"data"`
	Model      string      `json:"model"`
	Usage      Usage       `json:"usage"`
}
