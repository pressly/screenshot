package server

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/chromedp/cdproto/page"
	"github.com/pkg/errors"
	"github.com/pressly/screenshot/lib/headless"
	pb "github.com/pressly/screenshot/rpc/screenshot"
)

type Server struct{}

type resolution struct {
	x float64
	y float64
}

func parseResolution(s string) (*resolution, error) {
	parts := strings.Split(s, "x")
	if len(parts) < 3 {
		return nil, errors.New(fmt.Sprintf("failed to parse resolution: %s", s))
	}
	x, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse width: %s", parts[0])
	}
	y, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse height: %s", parts[0])
	}
	return &resolution{
		x: x,
		y: y,
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

	window, err := parseResolution(req.Window)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse window: %s", req.Window)
	}

	targetID := fmt.Sprintf("pressly/screenshot/%d", rand.Int63n(100))
	target := c.NewTarget(&targetID)
	if err := c.Run(ctx, target); err != nil {
		return nil, err
	}

	exec := c.GetHandlerByID(targetID)
	screenshotParams := page.CaptureScreenshot()
	screenshotParams.Clip = &page.Viewport{
		X:      float64(req.X),
		Y:      float64(req.Y),
		Width:  window.x,
		Height: window.y,
	}

	image, err := screenshotParams.Do(ctx, exec)

	return &pb.Resp{
		Resp: image,
	}, nil
}
