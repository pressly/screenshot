define run
	rerun -watch ./ -ignore bin rpc -run sh -c 'GOGC=off go build -i -o ./bin/$(1) ./cmd/$(1)/main.go && make proto && ./bin/$(1)'
endef

run:
	@$(MAKE) proto
	$(call run,screenshot)

run-example:
	rerun -watch ./ -ignore bin rpc -run sh -c 'sleep 1; go run ./example/go/main.go http://google.com /tmp/google.png && open /tmp/google.png'

proto:
	protoc --twirp_out=. --go_out=. ./rpc/screenshot/service.proto

.PHONY: vendor
vendor:
	GO111MODULE=on go mod vendor && GO111MODULE=on go mod tidy

docs:
	protoc --doc_out=./doc --doc_opt=html,index.html ./rpc/screenshot/*.proto && open ./doc/index.html