# Screenshot as a service
A simple scalable server built on top of Headless Chrome. The API is built with Protocol Buffers and [Twirp](https://github.com/twitchtv/twirp).

**This is an experimental project in early phase.**

# Goals:
1. Server should be scalable Docker image with HTTP server listening on port :80
2. The Docker image is ideally built on top of some official and stable Headless Chrome image (`FROM base/image:version`)
3. Clients should be automatically generated via Protobuf file + Twirp RPC framework
4. Built-in & configurable request throttling
5. Generate functional client in Go, JS (Typescript) and maybe other languages

# Usage

1. So far, this project works with [Chrome Canary](https://www.google.com/chrome/canary/) `/Applications/Google\ Chrome\ Canary.app/Contents/MacOS/Google\ Chrome\ Canary --headless --remote-debugging-port=9222 --no-gpu` 
2. However, it doesn't work with `docker run --name=chrome -p 9222:9222 --cap-add=SYS_ADMIN justinribeiro/chrome-headless --headless --remote-debugging-port=9222 --disable-gpu`, which is strange.. trying to figure that out

# Development

1. Install tools to work with [Protocol Buffers and Twirp](https://twitchtv.github.io/twirp/docs/install.html).

# License

Screenshot is licensed under the [MIT license](./LICENSE).
