package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"time"

	"github.com/pressly/screenshot/rpc/screenshot"
)

func main() {
	flag.Parse()

	ss := flag.Args()
	if len(ss) == 0 {
		log.Fatalln("must supply one or more URLs")
	}

	client := screenshot.NewScreenshotJSONClient("http://localhost:6666", &http.Client{})

	for _, s := range ss {

		resp, err := client.Image(context.Background(), &screenshot.RequestImage{Url: s})
		if err != nil {
			log.Fatal(err)
		}

		u, err := url.Parse(s)
		if err != nil {
			log.Fatalln(err)
		}

		name := fmt.Sprintf("%s_%s.png", u.Hostname(), time.Now().Format(time.RFC3339))
		if err := ioutil.WriteFile(filepath.Join("example/go/img", name), resp.Resp, 0644); err != nil {
			log.Fatalln(err)
		}
	}
}
