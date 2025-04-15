MODULE = $(shell go list -m)
PROJECT = $(shell basename $(MODULE))
VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || echo "1.0.0")
PACKAGES := $(shell go list ./... | grep -v /vendor/)
LDFLAGS := -ldflags "-X main.Version=${VERSION}"
FORTIFYFLAGS := "-Dcom.fortify.sca.ProjectRoot=.fortify" -verbose -debug
FORTIFYSCANFLAGS := $(FORTIFYFLAGS) -scan -rules etc/sast-custom-rules/example-custom-rules.xml -filter etc/sast-filters/example-filter.txt
# Which Go Router to use  - update to one of the following:
# - chi
# - echo
# - gin
# - gorilla
# - servemux
# - servemux-pre1.22
GOROUTER := servemux
CMDSERVER := cmd/$(GOROUTER)

.PHONY: default
default: help

# generate help info from comments: thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help: ## help information about make commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: test
test: ## run unit tests
	@echo "Not yet implemented"

.PHONY: test-cover
test-cover: test ## run unit tests and show test coverage information
	go tool cover -html=coverage-all.out

.PHONY: run
run: ## run the API server
	go run ${LDFLAGS} $(CMDSERVER)/main.go

.PHONY: build
build:  ## build the API server binary
	CGO_ENABLED=0 go build ${LDFLAGS} -a -o server $(MODULE)/$(CMDSERVER)

.PHONY: clean
clean: ## remove temporary files
	rm -rf server coverage.out coverage-all.out .fortify "$(PROJECT).fpr"

.PHONY: version
version: ## display the version of the API server
	@echo $(VERSION)

.PHONY: lint
lint: ## run golint on all Go package
	@golint $(PACKAGES)

.PHONY: fmt
fmt: ## run "go fmt" on all Go packages
	@go fmt $(PACKAGES)

.PHONY: sast-scan
sast-scan: ## run static application security testing
##	gosec -exclude=G104 ./...
	@sourceanalyzer $(FORTIFYFLAGS) -b "$(MODULE)" -clean
	@sourceanalyzer $(FORTIFYFLAGS) -b "$(MODULE)" -exclude vendor "cmd/$(GOROUTER)/*.go" "internal/$(GOROUTER)/*.go" "internal/server/*.go" 
	@sourceanalyzer $(FORTIFYFLAGS) -b "$(MODULE)" $(FORTIFYSCANFLAGS) -f "$(PROJECT).fpr"
	@FPRUtility -information -categoryIssueCounts -project "$(PROJECT).fpr" 
