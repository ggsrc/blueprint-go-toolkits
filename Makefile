NAME=blueprint-go-toolkits
TEST_DIRS := $(shell go list -f '{{.Dir}}' -m | xargs -I {} go list {}/...)

lint:
	@echo "--> Running linter"
	@go list -f '{{.Dir}}' -m | xargs -I {} golangci-lint run {}/...

lint-fix:
	@echo "--> Running linter auto fix"
	@go list -f '{{.Dir}}' -m | xargs -I {} golangci-lint run {}/... --fix

build:
	@go list -f '{{.Dir}}' -m | xargs -I {} go build -v {}/...

test:
	@export ENV=test && \
		go run gotest.tools/gotestsum@latest --format github-actions -- -tags=integration github.com/ggsrc/blueprint-go-toolkits/... -race -shuffle=on

codecov:
	export ENV=test && \
		go run gotest.tools/gotestsum@latest --format github-actions -- -tags=integration github.com/ggsrc/blueprint-go-toolkits/... -race -shuffle=on -coverprofile=coverage.txt -covermode=atomic
