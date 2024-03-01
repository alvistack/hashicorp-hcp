ROOT_DIR = $(dir $(abspath $(firstword $(MAKEFILE_LIST))))

# MOCKERY_DIRS are directories that contain a .mockery.yaml file. When go/mocks
# is run, the existing mocks will be deleted and new ones will be generated. If
# the mock file is generated in the same package as the actual implementation,
# store the mock file in MOCKERY_OUTPUT_FILES.
MOCKERY_DIRS=./ internal/commands/auth/ internal/pkg/api/iampolicy
MOCKERY_OUTPUT_DIRS=internal/pkg/api/mocks internal/commands/auth/mocks
MOCKERY_OUTPUT_FILES=internal/pkg/api/iampolicy/mock_setter.go \
					 internal/pkg/api/iampolicy/mock_resource_updater.go

default: help

.PHONY: build
build: ## Build the HCP CLI binary
	@go build -o bin/ ./...

.PHONY: screenshot
screenshot: go/install ## Create a screenshot of the HCP CLI
	@go run github.com/homeport/termshot/cmd/termshot@v0.2.7 -c -f assets/hcp.png -- hcp

.PHONY: go/install
go/install: ## Install the HCP CLI binary
	@go install

.PHONY: go/lint
go/lint: ## Run the Go Linter
	@golangci-lint run

.PHONY: go/mocks
go/mocks: ## Generates Go mock files.
	@for dir in $(MOCKERY_OUTPUT_DIRS); do \
		rm -rf $$dir; \
    done

	@for file in $(MOCKERY_OUTPUT_FILES); do \
		rm -f $$file; \
    done

	@for dir in $(MOCKERY_DIRS); do \
		cd $(ROOT_DIR); \
		cd $$dir; \
		mockery; \
    done

.PHONY: test
test: ## Run the unit tests
	@go test -v -cover ./...

# Docker build and publish variables and targets
REGISTRY_NAME?=docker.io/hashicorp
IMAGE_NAME=hcp
IMAGE_TAG_DEV?=$(REGISTRY_NAME)/$(IMAGE_NAME):latest-$(shell git rev-parse --short HEAD)
DEV_DOCKER_GOOS ?= linux
DEV_DOCKER_GOARCH ?= amd64

.PHONY: docker-build-dev
# Builds from the locally generated binary in ./bin/
docker-build-dev: export GOOS=$(DEV_DOCKER_GOOS)
docker-build-dev: export GOARCH=$(DEV_DOCKER_GOARCH)
docker-build-dev: build
	docker buildx build \
		--load \
		--platform $(DEV_DOCKER_GOOS)/$(DEV_DOCKER_GOARCH) \
		--tag $(IMAGE_TAG_DEV) \
		--target=dev \
		.
	@echo "Successfully built $(IMAGE_TAG_DEV)"

HELP_FORMAT="    \033[36m%-25s\033[0m %s\n"
.PHONY: help
help: ## Display this usage information
	@echo "Valid targets:"
	@grep -E '^[^ ]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		sort | \
		awk 'BEGIN {FS = ":.*?## "}; \
			{printf $(HELP_FORMAT), $$1, $$2}'
	@echo ""
