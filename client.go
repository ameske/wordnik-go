package wordnik

import "net/http"

const (
	BASE = "http://api.wordnik.com/v4"
)

// API is a wrapper around an http.Client that will append a developer's
// API key to every request
type APIClient struct {
	client *http.Client
	apiKey string
}

// NewClient creates an http.Client that will use the given API key
// for acesss to the Wordnik API
func NewAPIClient(key string) *APIClient {
	return &APIClient{
		client: &http.Client{},
		apiKey: key,
	}
}

func (c *APIClient) do(r *http.Request) (*http.Response, error) {
	r.Header["api_key"] = []string{c.apiKey}
	return c.client.Do(r)
}
