# Screenshot as a service
A simple scalable server built on top of Headless Chrome. The API is built with Protocol Buffers and [Twirp](https://github.com/twitchtv/twirp).

**This is an experimental project in early phase.**

# Goals:
1. Server should be scalable Docker image with HTTP server listening on port :80
2. The Docker image is ideally built on top of some official and stable Headless Chrome image (`FROM base/image:version`)
3. Clients should be automatically generated via Protobuf file + Twirp RPC framework
4. Built-in & configurable request throttling
5. Generate functional client in Go, JS (Typescript) and maybe other languages

## Base image candidates:
- https://hub.docker.com/r/justinribeiro/chrome-headless/
- https://hub.docker.com/r/selenoid/chrome/
- https://hub.docker.com/r/browserless/chrome/
- https://github.com/GoogleChrome/puppeteer/blob/fae441cd42fc410115ce9cec40563e21260db9b6/docs/troubleshooting.md#running-puppeteer-in-docker
- anything else?

# Installation

twirp setup steps:

https://twitchtv.github.io/twirp/docs/install.html

# License

Screenshot is licensed under the [MIT license](./LICENSE).
