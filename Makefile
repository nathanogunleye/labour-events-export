.DEFAULT_GOAL := help
.PHONY: help

help:
        @grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
        | sed -n 's/^\(.*\): \(.*\)##\(.*\)/\1\3/p' \
        | column -t  -s ' '

.PHONY: build
build: ## Build binary
	go build
