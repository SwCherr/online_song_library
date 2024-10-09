package base

import (
	"errors"
	"net/http"
	"os"
	"time"
)

type Client struct {
	Host   string
	Client *http.Client
}

func NewClient() (*Client, error) {
	host_api := os.Getenv("API_HOST")
	if host_api == "" {
		return &Client{}, errors.New("empty hostname API client")
	}

	return &Client{
		Host: host_api,
		Client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}, nil
}
