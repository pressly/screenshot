package main

import (
	"context"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/pressly/screenshot/rpc/screenshot"
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) != 2 {
		log.Fatalln("./example URL OUTPUT-FILE")
	}
	urlLink, filename := args[0], args[1]

	u, err := url.Parse(urlLink)
	if err != nil {
		log.Fatalln(err)
	}

	client := screenshot.NewScreenshotJSONClient("http://localhost:6666", &http.Client{})

	resp, err := client.Image(context.Background(), &screenshot.RequestImage{Url: u.String()})
	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile(filename, resp.Resp, 0644); err != nil {
		log.Fatalln(err)
	}
}
