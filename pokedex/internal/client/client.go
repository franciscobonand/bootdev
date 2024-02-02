package client

import (
	"fmt"
	"io"
	"net/http"
)

type HTTP struct{}

func NewHTTPClient() *HTTP {
	return &HTTP{}
}

func (h HTTP) Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
