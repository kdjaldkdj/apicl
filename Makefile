NAME=apicl
BUILD_DIR=_build

GIT_TAG=$(shell git tag -l 'v*' --points-at HEAD | tail -1)
GIT_HASH=$(shell git rev-parse --short HEAD)
GIT_HASH_FLAG=-X main.GitHash=$(GIT_HASH)
GIT_TAG_FLAG=-X main.GitTag=$(GIT_TAG)



linux: cmd/apicl/main.go
	GOOS=linux GOARCH=amd64 go build -ldflags "$(GIT_HASH_FLAG) $(GIT_TAG_FLAG)" -o $(BUILD_DIR)/$@/$(NAME) $^

mac: cmd/apicl/main.go
	GOOS=darwin GOARCH=amd64 go build -ldflags "$(GIT_HASH_FLAG) $(GIT_TAG_FLAG)" -o $(BUILD_DIR)/$@/$(NAME) $^

clean: 
	@rm -rf _build

.PHONY: linux mac

