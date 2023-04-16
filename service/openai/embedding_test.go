package openai_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/onemouth/trunk/service/openai"
)

func TestImpl_CreateEmbeddings(t *testing.T) {
	t.Parallel()

	mockCtx := context.Background()

	type fields struct {
		apiKey     string
		httpClient http.Client
	}

	tests := []struct {
		name    string
		fields  fields
		input   string
		wantErr bool
	}{
		{
			name:    "401 case",
			fields:  fields{apiKey: "", httpClient: http.Client{}},
			input:   "This is an apple",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			im := openai.New(tt.fields.apiKey, tt.fields.httpClient)
			_, err := im.CreateEmbeddings(mockCtx, tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("%s error = %v, wantErr %v", tt.name, err, tt.wantErr)

				return
			}
		})
	}
}
