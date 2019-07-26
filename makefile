.PHONY: test
test: ## Run test cases
	go test -short -cover -race -timeout=60s ./...

.PHONY: fmt
fmt: ## Run go fmt on go files
	go fmt ./...

EXCECUTABLE=frontier-cg

.PHONY: build
build: clean test 
	go build -v -o ${EXCECUTABLE}

.PHONY: run
run: clean test build 
	./${EXCECUTABLE}

.PHONY: docker-image
docker-image: clean 
	docker build --tag=frontiercg .

.PHONY: docker-run
docker-run: docker-image 
	docker run -d -p 8090:8080 frontiercg:latest

.PHONY: clean
clean: 
	if [ -f ${EXCECUTABLE} ] ; then rm ${EXCECUTABLE} ; fi

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := build