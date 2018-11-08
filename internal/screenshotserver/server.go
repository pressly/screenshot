package screenshotserver

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/chromedp/chromedp"
	pb "github.com/pressly/screenshot/rpc/screenshot"
)

type Server struct{}

func (s *Server) Screenshot(ctx context.Context, pageURL *pb.URL) (*pb.Empty, error) {
	// create context
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	cdp, teardown := newChromeInstance(ctx)
	defer teardown()

	fmt.Println(pageURL.String())
	// run task list
	var buf []byte
	if err := cdp.Run(ctx, chromedp.CaptureScreenshot(&buf)); err != nil {
		log.Fatalln(err)
	}

	if err := ioutil.WriteFile("sample.png", buf, 0644); err != nil {
		log.Fatalln(err)
	}

	return &pb.Empty{}, nil
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
