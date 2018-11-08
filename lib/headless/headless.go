package headless

import (
	"bytes"
	"context"
	"io"
	"log"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/pkg/errors"
)

type Chrome struct {
	*chromedp.CDP
	ctx context.Context
}

func New(ctx context.Context) (*Chrome, error) {
	c, err := chromedp.New(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to init chromedp")
	}

	return &Chrome{c, ctx}, nil
}

func (c *Chrome) Teardown() {
	// shutdown chrome
	if err := c.Shutdown(c.ctx); err != nil {
		log.Println(err)
	}
	// wait for chrome to finish
	if err := c.Wait(); err != nil {
		log.Println(err)
	}
}

func (c *Chrome) NewImage(ctx context.Context, addr string) (io.Reader, error) {
	// run task list
	var buf []byte
	err := c.Run(ctx,
		chromedp.Tasks{
			chromedp.Navigate(addr),
			chromedp.Sleep(3 * time.Second),
			chromedp.WaitReady("body"),
			chromedp.CaptureScreenshot(&buf)},
	)

	if err != nil {
		return nil, err
	}

	return bytes.NewReader(buf), nil
}
