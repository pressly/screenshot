define run
	rerun -watch ./ -ignore bin vendor rpc -run sh -c 'GOGC=off go build -i -o ./bin/$(1) ./cmd/$(1)/main.go && make proto && ./bin/$(1)'
endef

run:
	@$(MAKE) proto
	$(call run,screenshot)

run-example:
	rerun -watch ./ -ignore bin vendor rpc -run sh -c 'sleep 1; go run ./example/go/main.go http://google.com google.png && open google.png'

proto:
	protoc --twirp_out=. --go_out=. ./rpc/screenshot/service.proto