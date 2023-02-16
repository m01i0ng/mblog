COMMON_SELF_DIR:=$(dir $(lastword $(MAKEFILE_LIST)))
ROOT_DIR:=$(abspath $(shell cd $(COMMON_SELF_DIR)/ && pwd -P))
OUTPUT_DIR:=$(ROOT_DIR)/_output

.PHONY: all
all: add-copyright format build

.PHONY: build
build: tidy
	@go build -v -o $(OUTPUT_DIR)/mblog $(ROOT_DIR)/cmd/mblog/main.go

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
