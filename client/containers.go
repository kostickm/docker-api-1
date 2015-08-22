package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	//"github.com/icecrime/docker-api/api"
	"github.com/kostickm/docker-api-1/api"
)

func NewContainersServiceClient(client *http.Client, baseURI string) *containersClient {
	return &containersClient{
		baseURI: baseURI,
		client:  client,
	}
}

// containersClient provides client-side implementation of the BaseService interface.
type containersClient struct {
	baseURI string
	client  *http.Client
}

func (b *containersClient) List(_ *api.ListContainersParams) ([]*api.Container, error) {
	r, err := b.client.Get(fmt.Sprintf("%s/containers/ps", b.baseURI))
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var out []*api.Container
	if err := json.NewDecoder(r.Body).Decode(&out); err != nil {
		return nil, err
	}

	return out, nil
}

func (b *containersClient) Create(w interface{}) (*api.ListContainerID, error) {
	jsonData := bytes.NewBuffer(nil)
	if err := json.NewEncoder(jsonData).Encode(w); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/containers/create", b.baseURI), jsonData)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	r, err := b.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var containerID *api.ListContainerID
	if err := json.NewDecoder(r.Body).Decode(&containerID); err != nil {
		return nil, err
	}

	// TODO - Not sure how this command should look
	//containerID, warnings, err := b.daemon.ContainerCreate(r.Form, Get("name"), w.config, w.getHostConfig)
	//if err != nil {
	//	return nil, err
	//}

	// TODO - Need to return correctly
	return containerID, nil
}

//TODO - Want to return status and meaning, right now just returns status
func (b *containersClient) Start(id string) (int, error) {
	r, err := b.client.Post(fmt.Sprintf("%s/containers/%s/start", b.baseURI, id), "application/json", nil)
	if err != nil {
		return -1, err
	}
	defer r.Body.Close()

	// Is this necessary when we can get the status from the response using r.StatusCode ???
	var status int
	if err := json.NewDecoder(r.Body).Decode(&status); err != nil {
		return -1, err
	}

	return status, err
}
