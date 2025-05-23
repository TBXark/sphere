.PHONY: dash
dash: ## Build dash
ifneq ($(wildcard $(DASH_DIR)),)
	# You can `git clone https://github.com/pure-admin/vue-pure-admin.git $(DASH_DIR)` to get the dash project
	mkdir -p $(DASH_DIST)
	cd $(DASH_DIR) && pnpm build
	rm -rf $(DASH_DIST)/*
	cp -r $(DASH_DIR)/dist/* $(DASH_DIST)
else
	@echo "Skipping dash build - DASH_DIR does not exist"
endif

.PHONY: build
build: ## Build binary
	$(GO_BUILD) -o ./build/current_arch/ ./...

.PHONY: run
run:
	$(GO_RUN) $(MODULE)/cmd/app

.PHONY: build-linux-amd64
build-linux-amd64: ## Build linux amd64 binary
	GOOS=linux GOARCH=amd64 $(GO_BUILD) -o ./build/linux_amd64/ ./...

.PHONY: build-linux-arm64
build-linux-arm64: ## Build linux arm64 binary
	GOOS=linux GOARCH=arm64 $(GO_BUILD) -o ./build/linux_arm64/ ./...

.PHONY: build-all
build-all: build build-linux-amd64 ## Build all arch binary

.PHONY: delpoy
deploy: build-linux-amd64 ## Deploy binary
	ansible-playbook -i devops/ansible/hosts/inventory.ini devops/ansible/deploy.yml

.PHONY: lint
lint: ## Run linter
	go tool golangci-lint run
	go tool buf lint

.PHONY: fmt
fmt: ## Run formatter
	go tool golangci-lint fmt
	go tool golangci-lint run --fix
	go mod tidy
	go fmt ./...
	go tool buf format