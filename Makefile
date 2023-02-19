COMMON_SELF_DIR:=$(dir $(lastword $(MAKEFILE_LIST)))
ROOT_DIR:=$(abspath $(shell cd $(COMMON_SELF_DIR)/ && pwd -P))
OUTPUT_DIR:=$(ROOT_DIR)/_output

VERSION_PACKAGE=github.com/m01i0ng/mblog/pkg/version

ifeq ($(origin VERSION), undefined)
	VERSION:=$(shell git describe --tags --always --match='v*')
endif

GIT_TREE_STATE:="dirty"
ifeq (, $(shell git status --pocelain 2>/dev/null))
	GIT_TREE_STATE="clean"
endif

GIT_COMMIT:=$(shell git rev-parse HEAD)

GO_LDFLAGS += \
  -X $(VERSION_PACKAGE).GitVersion=$(VERSION) \
  -X $(VERSION_PACKAGE).GitCommit=$(GIT_COMMIT) \
  -X $(VERSION_PACKAGE).GitTreeState=$(GIT_TREE_STATE) \
  -X $(VERSION_PACKAGE).BuildDate=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')

.PHONY: all
all: add-copyright format build

.PHONY: build
build: tidy
	@go build -v -ldflags "$(GO_LDFLAGS)" -o $(OUTPUT_DIR)/mblog $(ROOT_DIR)/cmd/mblog/main.go

.PHONY: format
format:
	@gofmt -s -w ./

.PHONY: add-copyright
add-copyright:
	@addlicense -v -f $(ROOT_DIR)/scripts/boilerplate.txt $(ROOT_DIR) --skip-dirs=third_party,vendor,.idea,$(OUTPUT_DIR)

.PHONY: swagger
swagger:
	@swagger serve -F=swagger --no-open --port 65534 $(ROOT_DIR)/api/openapi/openapi.yaml

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: clean
clean:
	@-rm -vrf $(OUTPUT_DIR)
