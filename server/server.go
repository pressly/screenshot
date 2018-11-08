package server

import (
	"context"
	"io/ioutil"

	"github.com/pressly/screenshot/lib/headless"
	pb "github.com/pressly/screenshot/rpc/screenshot"
)

type Server struct{}

func (s *Server) Image(ctx context.Context, req *pb.RequestImage) (*pb.Resp, error) {
	// create context
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	c, err := headless.New(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Teardown()

	r, err := c.NewImage(ctx, req.Url)
	if err != nil {
		return nil, err
	}
	by, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return &pb.Resp{Resp: by}, nil
}
