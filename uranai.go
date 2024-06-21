package uranai

import (
	"bytes"
	"context"
	"fmt"
	m "github.com/anaregdesign/msproto/go/msp/azure/openai/chat/v1"
	"google.golang.org/protobuf/encoding/protojson"
	"io"
	"net/http"
)

type Teller struct {
	client *Client
}

func (t *Teller) Listen(ctx context.Context) (*ResultSet, error) {
	return nil, nil
}

type Client struct {
	httpClient     *http.Client
	resourceName   string
	deploymentName string
	apiVersion     string
	accessToken    string
}

func (c *Client) endpoint() string {
	return fmt.Sprintf("https://%s.openai.azure.com/openai/deployments/%s/chat/completions?api-version=%s", c.resourceName, c.deploymentName, c.apiVersion)
}

func (c *Client) header() http.Header {
	header := http.Header{}
	header.Add("api-key", c.accessToken)
	header.Add("Content-Type", "application/json")
	return header
}

func (c *Client) Get(ctx context.Context, request *m.CompletionRequest) (*m.CompletionResponse, error) {
	body, err := protojson.Marshal(request)
	if err != nil {
		return nil, err
	}

	httpRequest, err := http.NewRequestWithContext(ctx, "POST", c.endpoint(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	httpRequest.Header = c.header()

	httpResponse, err := c.httpClient.Do(httpRequest)
	defer httpResponse.Body.Close()
	if err != nil {
		return nil, err
	}

	responseBody, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}

	if httpResponse.StatusCode == http.StatusOK {
		response := &m.CompletionResponse{}
		err := protojson.Unmarshal(responseBody, response)
		if err != nil {
			return nil, err
		}
		return response, nil
	} else {
		fmt.Printf(string(responseBody))
		return nil, fmt.Errorf("unexpected status code: %d", httpResponse.StatusCode)
	}
}
