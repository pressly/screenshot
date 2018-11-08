package server

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/pressly/screenshot/lib/headless"
	pb "github.com/pressly/screenshot/rpc/screenshot"
)

type Server struct{}

type resolution struct {
	width  float64
	height float64
}

func parseResolution(s string) (*resolution, error) {
	parts := strings.Split(s, "x")
	if len(parts) < 2 {
		return nil, errors.New(fmt.Sprintf("failed to parse resolution: %s", s))
	}
	width, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse width: %s", parts[0])
	}
	height, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse height: %s", parts[1])
	}
	return &resolution{
		width:  width,
		height: height,
	}, nil
}

func (s *Server) Image(ctx context.Context, req *pb.RequestImage) (*pb.Resp, error) {
	// create context
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	c, err := headless.New(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Teardown()

	if req.Window == "" {
		req.Window = "800x600"
	}

	window, err := parseResolution(req.Window)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse window: %s", req.Window)
	}

	image, err := c.NewImage(ctx, req.Url, float64(req.X), float64(req.Y), window.width, window.height)
	if err != nil {
		return nil, err
	}

	return &pb.Resp{
		Resp: image,
	}, nil
}
