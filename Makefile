all: go-deps ui-deps ui go

gen:
	go generate ./...
	 cd web && yarn graphql-codegen -c ../codegen.yml

go-deps:
	git submodule update --init --recursive
	make -C extern/filecoin-ffi

go:
	go build -o curio-dashboard ./cmd

ui-deps:
	cd ui && yarn install

ui:
	cd ui && yarn build
.PHONY: ui
