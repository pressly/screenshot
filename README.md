# Screenshot as a service
A simple scalable server built on top of Headless Chrome. The API is built with Protocol Buffers and [Twirp](https://github.com/twitchtv/twirp).

This is an experimental project in early phase.

Goals:
1. Server should be scalable Docker image with HTTP server listening on port :80
2. The Docker image is ideally built on top of some official and stable Headless Chrome image (`FROM base/image:version`)
3. Clients should be automatically generated via Protobuf file + Twirp RPC framework
4. Built-in & configurable request throttling
5. Generate functional client in Go, JS (Typescript) and maybe other languages

# License

Screenshot is licensed under the [MIT license](./LICENSE).
