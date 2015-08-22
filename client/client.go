package client

import (
	"net/http"

	//"github.com/icecrime/docker-api/api"
	"github.com/kostickm/docker-api-1/api"
)

func New(c *http.Client, baseURI string) *Client {
	return &Client{
		BaseService:       NewBaseServiceClient(c, baseURI),
		ContainersService: NewContainersServiceClient(c, baseURI),
	}
}

type Client struct {
	api.BaseService
	api.ContainersService
}
