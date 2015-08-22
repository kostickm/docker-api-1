package main

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	//"github.com/icecrime/docker-api/api"
	//"github.com/icecrime/docker-api/server"
	"github.com/kostickm/docker-api-1/api"
	"github.com/kostickm/docker-api-1/server"
)

type testServer struct{}

func (*testServer) List(p *api.ListContainersParams) ([]*api.Container, error) {
	log.Infof("testServer.List(%v)", p)
	return []*api.Container{}, nil
}

func (*testServer) Ping() (string, error) {
	log.Info("testServer.Ping()")
	return "OK", nil
}

func (*testServer) Version() (*api.Version, error) {
	log.Info("testServer.Version()")
	return &api.Version{
		APIVersion:    "APIVersion",
		Arch:          "Arch",
		GitCommit:     "GitCommit",
		GoVersion:     "GoVersion",
		KernelVersion: "KernelVersion",
		OS:            "OS",
		Version:       "Version",
	}, nil
}

func (*testServer) Create(p interface{}) (*api.ListContainerID, error) {
	log.Infof("testServer.Create(%v)", p)
	var containerID *api.ListContainerID
	return containerID, nil
}

func (*testServer) Start(p string) (int, error) {
	log.Infof("testServer.Start(%v)", p)
	var statusCode int
	return statusCode, nil
}

func main() {
	srv := server.New(&testServer{}, "swagger-ui/dist/")
	if err := http.ListenAndServe("127.0.0.1:8080", srv); err != nil {
		panic(err)
	}
}
