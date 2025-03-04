package gapi

import (
	"github.com/mzzz-zzm/blazor-scaffold/svr/pb/greet"
	// "github.com/mzzz-zzm/blazor-scaffold/svr/utils"
)

type Server struct {
	// Config *Config
	greet.UnimplementedGreeterServer
	// TODO: add config, db store, token generator
}

func NewServer() (*Server, error) {

	// TODO: create a token

	server := &Server{}

	return server, nil
}
