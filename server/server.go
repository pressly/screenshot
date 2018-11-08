package server

import (
	"context"
	"io/ioutil"

	"github.com/pkg/errors"
	"github.com/pressly/screenshot/lib/headless"
	pb "github.com/pressly/screenshot/rpc/screenshot"
)

type server struct {
	chrome *headless.Chrome
}

func New() (*server, error) {
	chrome, err := headless.New(context.TODO())
	if err != nil {
		return nil, errors.Wrap(err, "failed to init headless")
	}

	return &server{chrome: chrome}, nil
}

func (s *server) Close() {
	s.chrome.Close()
}

func (s *server) Image(ctx context.Context, req *pb.RequestImage) (*pb.Resp, error) {
	// create context
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	r, err := s.chrome.NewImage(ctx, req.Url)
	if err != nil {
		return nil, err
	}
	by, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return &pb.Resp{Resp: by}, nil
}
