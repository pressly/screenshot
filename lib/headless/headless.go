package headless

import (
	"context"
	"log"

	"github.com/chromedp/cdproto/page"
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

func (c *Chrome) Close() {
	// shutdown chrome
	if err := c.Shutdown(c.ctx); err != nil {
		log.Println(err)
	}
	// wait for chrome to finish
	if err := c.Wait(); err != nil {
		log.Println(err)
	}
}

func (c *Chrome) NewImage(ctx context.Context, addr string, x, y, width, height float64) ([]byte, error) {
	// run task list

	var buf []byte
	err := c.Run(ctx,
		chromedp.Tasks{
			chromedp.Navigate(addr),
			chromedp.WaitReady("body"),
			chromedp.CaptureScreenshot(&buf),
		})
	if err != nil {
		return nil, err
	}

	currentTarget := c.GetHandlerByIndex(0)
	image, err := page.CaptureScreenshot().Do(ctx, currentTarget)
	if err != nil {
		return nil, err
	}

	return image, nil
}
