package server

import (
	"context"
	"strconv"
	"strings"

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

type resolution struct {
	width  float64
	height float64
}

// parseResolution parses a string of format 800x600, and converts it to
// resolution. If s is empty, an initilized zero value resolution is returned.
func parseResolution(s string) (*resolution, error) {
	if s == "" {
		return &resolution{}, nil
	}

	parts := strings.Split(s, "x")
	if len(parts) < 2 {
		return nil, errors.Errorf("failed to parse resolution: %s", s)
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

func (s *server) Close() {
	s.chrome.Close()
}

func (s *server) Image(ctx context.Context, req *pb.RequestImage) (*pb.Resp, error) {
	// create context
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	window, err := parseResolution(req.Window)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse window resolution")
	}

	image, err := s.chrome.NewImage(ctx, req.Url, float64(req.X), float64(req.Y), window.width, window.height)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse window: %s", req.Window)
	}

	return &pb.Resp{
		Resp: image,
	}, nil
}
