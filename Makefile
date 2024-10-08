all: go-deps ui-deps ui go
all-in-one: go-deps ui-deps ui go-dist
aio: all-in-one

go-deps:
	git submodule update --init --recursive
	make -C extern/filecoin-ffi

go:
	go build -ldflags="-s -w" -o curio-dashboard ./cmd

go-dist:
	go build -ldflags="-s -w" -o curio-dashboard -tags dist ./cmd

ui-deps:
	cd ui && yarn install

ui:
	cd ui && yarn build
.PHONY: ui

lint:
	golangci-lint run -v ./...
	cd ui && yarn lint

curio-rpc-gen:
	go run ./graph/curiorpc/gen
	goimports -w graph/curiorpc
.PHONY: api-gen

type-gen:
	 cd ui && yarn graphql-codegen -c ../codegen.yml

go-gen:
	go generate ./...

gen: go-gen curio-rpc-gen type-gen
