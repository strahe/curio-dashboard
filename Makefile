all: go-deps web-deps web go

gen:
	go generate ./...
	 cd web && yarn graphql-codegen -c ../codegen.yml

go-deps:
	git submodule update --init --recursive
	make -C extern/filecoin-ffi

go:
	go build -o curio-dashboard ./cmd

web-deps:
	cd web && yarn install

web:
	cd web && yarn build
.PHONY: web
