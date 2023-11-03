.PHONY: clean test bins lint tools
GITHUB_REPOSITORY = temporalio/temporal-demo-infra

# default target
default: clean test bins

TAG_COMMIT := $(shell git rev-list --abbrev-commit --tags --max-count=1)
TAG := $(shell git describe --abbrev=0 --tags ${TAG_COMMIT} 2>/dev/null || true)
COMMIT := $(shell git rev-parse --short HEAD )
SHORT_SHA := $(echo $COMMIT | cut -c 1-8)
DATE := $(shell git log -1 --format=%cd --date=format:"%Y%m%d")
VERSION := $(TAG:v%=%)
DOMAIN_BUILD_PKG := github.com/$(GITHUB_REPOSITORY)/build
LINKER_FLAGS := "-X '${DOMAIN_BUILD_PKG}.Commit=${COMMIT}' -X '${DOMAIN_BUILD_PKG}.Version=${VERSION}' -X '${DOMAIN_BUILD_PKG}.BuildDate=${DATE}'"

out:
	mkdir -p out

domain: out
	@go build -ldflags ${LINKER_FLAGS} -o ./out/domain ./main.go

bins: domain

test:
	go test -race -timeout=5m -cover -count=1  ./...

clean:
	@rm -rf out

lint:
	golangci-lint run .