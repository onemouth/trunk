package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	embeddingURL   = "https://api.openai.com/v1/embeddings"
	embeddingModel = "text-embedding-ada-002"
)

func getEmbeddingReqBody(input string) ([]byte, error) {
	//nolint: wrapcheck // internal function
	return json.Marshal(EmbeddingsReq{Model: embeddingModel, Input: input})
}

func (im *Impl) CreateEmbeddings(ctx context.Context, input string) (*EmbeddingsResp, error) {
	reqBody, err := getEmbeddingReqBody(input)
	if err != nil {
		return nil, fmt.Errorf("CreateEmbeddings: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, embeddingURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("CreateEmbeddings: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", im.apiKey))

	httpResp, err := im.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("CreateEmbeddings: %w", err)
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("CreateEmbeddings: statusCode: %d,  %w", httpResp.StatusCode, ErrHTTPNot200)
	}

	adaResp := EmbeddingsResp{}
	if err := json.NewDecoder(httpResp.Body).Decode(&adaResp); err != nil {
		return nil, fmt.Errorf("CreateEmbeddings: %w", err)
	}

	return &adaResp, nil
}
