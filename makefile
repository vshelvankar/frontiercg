.PHONY: test
test: ## Run test cases
	go test -short -cover -race -timeout=60s ./...

.PHONY: fmt
fmt: ## Run go fmt on go files
	go fmt ./...

BINARY=frontier-cg

.PHONY: build
build: clean test ## Build binary. Can have machine specific seperate build commands
	go build -v -o ${BINARY}

.PHONY: clean
clean: ## Remove binary and temp files
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := build