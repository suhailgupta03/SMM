SHELL := /bin/bash

ASSETS := assets
PLUGIN_CODE_ROOT := plugins
PLUGIN_BUILD_ROOT := $(ASSETS)/plugins

# Plugin directory names
DJANGO_EOL := djangoeol
ECR_VUL := ecrvuln
HAS_JSON_LOGGING := hasjsonlogging
HAS_LOGGING := haslogging
LATEST_PATCH_DJANGO := latestpatchdjango
LATEST_PATCH_NODE := latestpatchnode
LATEST_PATCH_PYTHON := latestpatchpython
NODE_EOL := nodeeol
PYTHON_EOL := pythoneol
REACT_EOL := reacteol
GO_EOL := goeol
README := readme
REPO_VUL := repovuln
MIN_COV := mincov

# Plugin build paths
PLUGIN_BUILD_PATHS := $(PLUGIN_BUILD_ROOT)/$(DJANGO_EOL).so \
				$(PLUGIN_BUILD_ROOT)/$(ECR_VUL).so \
				$(PLUGIN_BUILD_ROOT)/$(HAS_JSON_LOGGING).so \
				$(PLUGIN_BUILD_ROOT)/$(HAS_LOGGING).so \
				$(PLUGIN_BUILD_ROOT)/$(LATEST_PATCH_DJANGO).so \
				$(PLUGIN_BUILD_ROOT)/$(LATEST_PATCH_NODE).so \
				$(PLUGIN_BUILD_ROOT)/$(LATEST_PATCH_PYTHON).so \
				$(PLUGIN_BUILD_ROOT)/$(NODE_EOL).so \
				$(PLUGIN_BUILD_ROOT)/$(PYTHON_EOL).so \
				$(PLUGIN_BUILD_ROOT)/$(REACT_EOL).so \
				$(PLUGIN_BUILD_ROOT)/$(README).so \
				$(PLUGIN_BUILD_ROOT)/$(REPO_VUL).so \
				$(PLUGIN_BUILD_ROOT)/$(MIN_COV).so \
				$(PLUGIN_BUILD_ROOT)/$(GO_EOL).so


BIN := smm
PLUGIN_GO_FILES := $(shell find . -type f  ! -name "runner.go" -name "*.go")
RUNNER_FILES := $(shell find . -path ./plugins -prune -o -name '*.go')

plugins: $(PLUGIN_BUILD_PATHS)

$(PLUGIN_BUILD_ROOT)/$(DJANGO_EOL).so: $(PLUGIN_GO_FILES)
	go build -buildmode=plugin -o $@ $(PLUGIN_CODE_ROOT)/$(DJANGO_EOL)/$(DJANGO_EOL).go

$(PLUGIN_BUILD_ROOT)/$(ECR_VUL).so: $(PLUGIN_GO_FILES)
	go build -buildmode=plugin -o $@ $(PLUGIN_CODE_ROOT)/$(ECR_VUL)/$(ECR_VUL).go

$(PLUGIN_BUILD_ROOT)/$(HAS_JSON_LOGGING).so: $(PLUGIN_GO_FILES)
	go build -buildmode=plugin -o $@ $(PLUGIN_CODE_ROOT)/$(HAS_JSON_LOGGING)/$(HAS_JSON_LOGGING).go

$(PLUGIN_BUILD_ROOT)/$(HAS_LOGGING).so: $(PLUGIN_GO_FILES)
	go build -buildmode=plugin -o $@ $(PLUGIN_CODE_ROOT)/$(HAS_LOGGING)/$(HAS_LOGGING).go

$(PLUGIN_BUILD_ROOT)/$(LATEST_PATCH_DJANGO).so: $(PLUGIN_GO_FILES)
	go build -buildmode=plugin -o $@ $(PLUGIN_CODE_ROOT)/$(LATEST_PATCH_DJANGO)/$(LATEST_PATCH_DJANGO).go

$(PLUGIN_BUILD_ROOT)/$(LATEST_PATCH_NODE).so: $(PLUGIN_GO_FILES)
	go build -buildmode=plugin -o $@ $(PLUGIN_CODE_ROOT)/$(LATEST_PATCH_NODE)/$(LATEST_PATCH_NODE).go

$(PLUGIN_BUILD_ROOT)/$(LATEST_PATCH_PYTHON).so: $(PLUGIN_GO_FILES)
	go build -buildmode=plugin -o $@ $(PLUGIN_CODE_ROOT)/$(LATEST_PATCH_PYTHON)/$(LATEST_PATCH_PYTHON).go

$(PLUGIN_BUILD_ROOT)/$(NODE_EOL).so: $(PLUGIN_GO_FILES)
	go build -buildmode=plugin -o $@ $(PLUGIN_CODE_ROOT)/$(NODE_EOL)/$(NODE_EOL).go

$(PLUGIN_BUILD_ROOT)/$(PYTHON_EOL).so: $(PLUGIN_GO_FILES)
	go build -buildmode=plugin -o $@ $(PLUGIN_CODE_ROOT)/$(PYTHON_EOL)/$(PYTHON_EOL).go

$(PLUGIN_BUILD_ROOT)/$(REACT_EOL).so: $(PLUGIN_GO_FILES)
	go build -buildmode=plugin -o $@ $(PLUGIN_CODE_ROOT)/$(REACT_EOL)/$(REACT_EOL).go

$(PLUGIN_BUILD_ROOT)/$(README).so: $(PLUGIN_GO_FILES)
	go build -buildmode=plugin -o $@ $(PLUGIN_CODE_ROOT)/$(README)/$(README).go

$(PLUGIN_BUILD_ROOT)/$(REPO_VUL).so: $(PLUGIN_GO_FILES)
	go build -buildmode=plugin -o $@ $(PLUGIN_CODE_ROOT)/$(REPO_VUL)/$(REPO_VUL).go

$(PLUGIN_BUILD_ROOT)/$(MIN_COV).so: $(PLUGIN_GO_FILES)
	go build -buildmode=plugin -o $@ $(PLUGIN_CODE_ROOT)/$(MIN_COV)/$(MIN_COV).go

$(PLUGIN_BUILD_ROOT)/$(GO_EOL).so: $(PLUGIN_GO_FILES)
	go build -buildmode=plugin -o $@ $(PLUGIN_CODE_ROOT)/$(GO_EOL)/$(GO_EOL).go

.PHONY: build
build: $(BIN)

$(BIN): $(RUNNER_FILES)
	go build -o $(BIN)

.PHONY: run
run:
	go run runner.go

.PHONY: run-local
run-local: $(plugins)
	source test.env && make run

.PHONY: dist
dist: build $(plugins)
	rm -rf dist && mkdir dist && cp -R assets dist/assets && cp ${MATURITY_REPO_YAML} dist/${MATURITY_REPO_YAML} && cp ./smm dist/

# Run Go tests
.PHONY: test
test:
	go test ./... -coverpkg=./... -coverprofile coverage.txt
	go tool cover -func=coverage.txt
	go tool cover -html=coverage.txt