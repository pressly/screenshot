proto:
	protoc --twirp_out=. --go_out=. ./rpc/screenshot/service.proto