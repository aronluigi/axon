.PHONY: help

help: ## Show this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m\033[0m\n"} /^[0-9a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

install: ## install utils
	@go install github.com/air-verse/air@latest
	@go install golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment@latest
	@asdf reshimer golang

format: ## Format source code
	fieldalignment -fix ./pkg/...

watch: ## Start dev watch
	@air

codegen: ## Run code generation for entire project
	@make -C ./pkg/gql codegen
