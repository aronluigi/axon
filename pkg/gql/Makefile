.PHONY: help

help: ## Show this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m\033[0m\n"} /^[0-9a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

install: ## Install dependencies
	@go get github.com/99designs/gqlgen

codegen: ## Generate GQL code
	@rm -rf ./graph/generated/ && rm -rf ./graph/model/
	@go run github.com/99designs/gqlgen generate
