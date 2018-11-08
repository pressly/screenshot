package main

import (
	"log"
	"net/http"

	"github.com/pressly/screenshot/internal/screenshotserver"
	"github.com/pressly/screenshot/rpc/screenshot"
)

func main() {

	srv := &screenshotserver.Server{}
	twirpHandler := screenshot.NewScreenshoterServer(srv, nil)

	log.Fatal(http.ListenAndServe(":6666", twirpHandler))
}
