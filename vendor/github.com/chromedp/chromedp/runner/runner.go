// Package runner provides a Chrome process runner.
package runner

import (
	"context"

	"github.com/chromedp/chromedp/client"
)

// Runner holds information about a running Chrome process.
type Runner struct {
}

// New creates a new Chrome process using the supplied command line options.
func New(opts ...CommandLineOption) (*Runner, error) {
	return &Runner{}, nil
}

// Start starts a Chrome process using the specified context. The Chrome
// process can be terminated by closing the passed context.
func (r *Runner) Start(ctxt context.Context, opts ...string) error {

	return nil
}

// Shutdown shuts down the Chrome process.
func (r *Runner) Shutdown(ctxt context.Context, opts ...client.Option) error {
	return nil
}

// Wait waits for the previously started Chrome process to terminate, returning
// any encountered error.
func (r *Runner) Wait() error {
	return nil
}

// Client returns a Chrome DevTools Protocol client for the running Chrome
// process.
func (r *Runner) Client(opts ...client.Option) *client.Client {
	return client.New(append(opts,
		client.URL("http://localhost:9222/json"),
	)...)
}

// Run starts a new Chrome process runner, using the provided context and
// command line options.
func Run(ctxt context.Context, opts ...CommandLineOption) (*Runner, error) {
	var err error

	// create
	r, err := New(opts...)
	if err != nil {
		return nil, err
	}

	// start
	if err = r.Start(ctxt); err != nil {
		return nil, err
	}

	return r, nil
}

// CommandLineOption is a runner command line option.
//
// see: http://peter.sh/experiments/chromium-command-line-switches/
type CommandLineOption func(map[string]interface{}) error
