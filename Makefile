include .env
PROJECT_NAME=$(shell basename "$(PWD)")
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
SERVICE=appbackend
# gitからversion情報を取得
VERSION := $(shell git describe --always --long --dirty)
# プロジェクトで使っているpkgの一覧(vendor抜き)
PKG_LIST := $(shell go list ./... | grep -v /vendor/)
# プロジェクト内のgofileの一覧(vendor抜き)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/)

CURRENT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)

MAKEFLAGS += --silent

.DEFAULT_GOAL := help


.PHONY: setup
setup: ## Setup installs go binary
	$(GOGET) github.com/cespare/reflex
	$(GOGET) github.com/oxequa/realize
	$(GOGET) golang.org/x/lint/golint
	$(GOGET) golang.org/x/tools/cmd/goimports
	$(GOGET) github.com/pressly/goose/cmd/goose


.PHONY: dev
dev: ## Copy git dir
	cp -P git/pre-commit .git/hooks/pre-commit

.PHONY: deps
deps: ## Install vendors
	$(GOGET) github.com/golang/dep/cmd/dep
	@dep ensure

.PHONY: pretest
pretest: lint errcheck staticcheck

.PHONY: test
test: pretest ## Test tests gofiles
	@go test -race ${PKG_LIST}

.PHONY: errcheck
errcheck: ## Errcheck checks error
	$(GOGET) github.com/kisielk/errcheck
	@errcheck ${PKG_LIST}

.PHONY: staticcheck
staticcheck: ## Staticcheck checks anything
	$(GOGET) honnef.co/go/tools/cmd/staticcheck
	@staticcheck ${PKG_LIST}

.PHONY: lint
lint: ## Lint checks golint
	$(GOGET) github.com/golang/lint/golint
	@for file in ${GO_FILES} ;  do \
		golint $$file ; \
	done

.PHONY: shell
shell: ## shell to Docker appbackend
	docker exec -it ${SERVICE} sh

.PHONY: help
all: help
help:
	@awk -F ':|##' '/^[^\t].+?:.*?##/ {\
		printf "\033[36m%-30s\033[0m %s\n", $$1, $$NF \
  }' ${MAKEFILE_LIST}
