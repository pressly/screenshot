package server

import (
	"context"
	"io/ioutil"
	"log"
	"time"

	"github.com/chromedp/chromedp"
	pb "github.com/pressly/screenshot/rpc/screenshot"
)

type Server struct{}

func (s *Server) Image(ctx context.Context, req *pb.RequestImage) (*pb.Resp, error) {
	// create context
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()


	chrome, headless.NewImage


	cdp, teardown := newChromeInstance(ctx)
	defer teardown()

	// run task list
	var buf []byte
	if err := cdp.Run(ctx,
		chromedp.Tasks{
			chromedp.Navigate(req.Url),
			chromedp.Sleep(3 * time.Second),
			chromedp.WaitReady("body"),
			chromedp.CaptureScreenshot(&buf)}); err != nil {
		log.Fatalln(err)
	}

	if err := ioutil.WriteFile("img/sample.png", buf, 0644); err != nil {
		log.Fatalln(err)
	}

	return &pb.Resp{}, nil
}

func newChromeInstance(ctx context.Context) (*chromedp.CDP, func()) {
	// create chrome instance
	c, err := chromedp.New(ctx, chromedp.WithLog(log.Printf))
	if err != nil {
		log.Fatalln(err)
	}

	teardown := func() {
		// shutdown chrome
		if err := c.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
		// wait for chrome to finish
		err = c.Wait()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("teardown complete. no errors.\n")
	}

	return c, teardown
}
