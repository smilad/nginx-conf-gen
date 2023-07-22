GOPATH:=$(shell go env GOPATH)

.PHONY: init
init:
	make submodule
	make proto
	make config
	make gitlab

.PHONY: submodule
submodule:
	git submodule init
	git submodule update

.PHONY: proto
proto:
	mkdir -p api/pb

	protoc -I pb/proto -I pb/lib \
	--go_out api/pb \
	--go_temporal_opt paths=source_relative \
	--go_opt paths=source_relative \
	--go-grpc_out api/pb \
	--go-grpc_opt paths=source_relative \
	--go_temporal_out api/pb \
	--go_temporal_opt paths=source_relative pb/proto/base/temporal_service.proto

.PHONY: build
build:
	go build -o base *.go

.PHONY: test
test:
	go test -v ./... -cover -race

.PHONY: vendor
vendor:
	go get ./...
	go mod vendor
	go mod verify

.PHONY: config
config:
	cp -rf ./config.example.yaml ./config.yaml
	cp -rf ./config.example.yaml ./config.test.yaml

.PHONY: gitlab
gitlab:
	-cp -rf ./-gitlab-ci.yml ./.gitlab-ci.yml
	-rm -rf ./-gitlab-ci.yml

# ==============================================================================
# Swagger

swagger:
	@echo Starting swagger generating
	swag init -g **/**/*.go