package main

import (
	"context"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/pressly/screenshot/rpc/screenshot"
)

var (
	flags = flag.NewFlagSet("screenshot", flag.ExitOnError)

	window = flags.String("window", "", "{width}x{height}")
)

func main() {
	flags.Parse(os.Args[1:])

	args := flags.Args()
	if len(args) < 2 {
		log.Fatalln("./example -window={width}x{height} URL OUTPUT-FILE")
	}
	urlLink, filename := args[0], args[1]

	u, err := url.Parse(urlLink)
	if err != nil {
		log.Fatalln(err)
	}

	client := screenshot.NewScreenshotProtobufClient("http://localhost:6666", &http.Client{})

	resp, err := client.Image(context.Background(), &screenshot.RequestImage{Url: u.String(), Window: *window})
	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile(filename, resp.Resp, 0644); err != nil {
		log.Fatalln(err)
	}
}
