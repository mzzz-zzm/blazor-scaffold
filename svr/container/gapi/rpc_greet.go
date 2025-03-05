package gapi

import (
	"context"

	"github.com/mzzz-zzm/blazor-scaffold/svr/pb/greet"
)

func (server *Server) Greet(ctx context.Context, req *greet.GreetRequest) (*greet.GreetReply, error) {
	return &greet.GreetReply{Message: "Hello " + req.Name}, nil
}
