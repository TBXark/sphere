MODULE := $(shell go list -m)
MODULE_NAME := $(lastword $(subst /, ,$(MODULE)))
BUILD := $(shell git rev-parse --short HEAD)@$(shell date +%s)
CURRENT_OS := $(shell uname -s | tr '[:upper:]' '[:lower:]')
CURRENT_ARCH := $(shell uname -m | tr '[:upper:]' '[:lower:]')

DOCKER_IMAGE=ghcr.io/tbxark/$(MODULE_NAME)
DOCKER_FILE=cmd/app/Dockerfile


LD_FLAGS="-X $(MODULE)/configs.BuildVersion=$(BUILD)"
GO_BUILD=CGO_ENABLED=0 go build -ldflags $(LD_FLAGS)

.PHONY: init
init:
	go get entgo.io/ent/cmd/ent@latest
	go get github.com/google/wire/cmd/wire@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go mod download
	make generate
	make docs
	#make config

.PHONY: generate
generate:
	go generate ./...

.PHONY: config
config:
	go run ./cmd/config gen

.PHONY: docs
docs:
	rm -rf ./docs/dash
	rm -rf ./docs/api
	swag init --output ./docs/api  --exclude internal/biz/dash --instanceName API  -g internal/biz/api/web.go
	swag init --output ./docs/dash --exclude internal/biz/api  --instanceName Dash -g internal/biz/dash/web.go

.PHONY: typescript
typescript: docs
	npx swagger-typescript-api -p ./docs/api/API_swagger.json   -o ./docs/api/typescript  --modular
	npx swagger-typescript-api -p ./docs/dash/Dash_swagger.json -o ./docs/dash/typescript --modular

.PHONY: tmpl
tmpl:
	go run -tags=tmplgen ./assets/tmpl/gen/generate.go ./assets/tmpl

.PHONY: dash
dash:
	sh ./assets/dash/build.sh

.PHONY: build
build:
	$(GO_BUILD) -o ./build/$(CURRENT_OS)_$(CURRENT_ARCH)/ ./...

.PHONY: buildLinuxX86
buildLinuxX86:
	GOOS=linux GOARCH=amd64 $(GO_BUILD) -o ./build/linux_x86/ ./...

.PHONY: buildLinuxARM64
buildLinuxARM64:
	GOOS=linux GOARCH=arm64 $(GO_BUILD) -o ./build/linux_arm64/ ./...

.PHONY: buildDockerImage
buildDockerImage:
	docker buildx build --platform=linux/amd64,linux/arm64 -t $(DOCKER_IMAGE) . -f  $(DOCKER_FILE) --push --provenance=false

.PHONY: delpoy
deploy:
	ansible-playbook -i devops/hosts/inventory.ini devops/delpoy-binary.yaml

.PHONY: lint
lint:
	golangci-lint run