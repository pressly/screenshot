package main

import (
	"log"
	"net/http"

	"github.com/pressly/screenshot/rpc/screenshot"
	"github.com/pressly/screenshot/server"
)

func main() {
	srv := server.New()

	twirpHandler := screenshot.NewScreenshotServer(srv, nil)

	log.Fatal(http.ListenAndServe(":6666", twirpHandler))
}
