package main

import (
	"log"
	"net/http"

	"github.com/pressly/screenshot/internal/server"
	"github.com/pressly/screenshot/rpc/screenshot"
)

func main() {

	srv := &server.Server{}
	twirpHandler := screenshot.NewScreenshotServer(srv, nil)

	log.Fatal(http.ListenAndServe(":6666", twirpHandler))
}
