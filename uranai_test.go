package uranai

import (
	"context"
	"fmt"
	m "github.com/anaregdesign/msproto/go/msp/azure/openai/chat/v1"
	"net/http"
	"os"
	"testing"
)

func TestClient_Get(t *testing.T) {
	request := &m.CompletionRequest{}
	request.Messages = []*m.CompletionRequest_Message{
		{
			Role:    "user",
			Content: "hello world",
		},
	}

	type fields struct {
		httpClient     *http.Client
		resourceName   string
		deploymentName string
		apiVersion     string
		accessToken    string
	}
	type args struct {
		ctx     context.Context
		request *m.CompletionRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *m.CompletionResponse
		wantErr bool
	}{
		{
			name: "valid case ",
			fields: fields{
				httpClient:     &http.Client{},
				resourceName:   "example-aoai-02",
				deploymentName: "gpt-4o",
				apiVersion:     "2024-02-01",
				accessToken:    os.Getenv("AOAI_API_KEY"),
			},
			args: args{
				ctx:     context.Background(),
				request: request,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				httpClient:     tt.fields.httpClient,
				resourceName:   tt.fields.resourceName,
				deploymentName: tt.fields.deploymentName,
				apiVersion:     tt.fields.apiVersion,
				accessToken:    tt.fields.accessToken,
			}
			got, err := c.Get(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}

func TestTeller_Listen(t1 *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ResultSet
		wantErr bool
	}{
		{
			name: "valid case",
			fields: fields{
				client: &Client{
					httpClient:     &http.Client{},
					resourceName:   "example-aoai-02",
					deploymentName: "gpt-4o",
					apiVersion:     "2024-02-01",
					accessToken:    os.Getenv("AOAI_API_KEY"),
				},
			},
			args: args{
				context.Background(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Teller{
				client: tt.fields.client,
			}
			got, err := t.Listen(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t1.Errorf("Listen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}
