module github.com/pressly/screenshot

require (
	github.com/chromedp/chromedp v0.1.2
	github.com/golang/protobuf v1.2.0
	github.com/pkg/errors v0.8.0
	github.com/twitchtv/twirp v5.5.0+incompatible
	golang.org/x/sync v0.0.0-20181108010431-42b317875d0f // indirect
)

// We've forked chromedp to avoid the chromedp/runner to exec local Chrome processes.
replace github.com/chromedp/chromedp v0.1.2 => github.com/pressly/chromedp v0.1.3-0.20180717231922-bf52fed0d3e6
