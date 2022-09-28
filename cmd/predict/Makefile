TARGET := recsys-predict
MAKEFILE_PATH := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
git_tag=$(shell git describe --long --tags --dirty --tags --always)
git_commit=$(shell git rev-parse HEAD)
git_build_time=$(shell TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ')

GO := go
GO_DOWNLOAD_FLAGS := --insecure
GO_BUILD_FLAGS := -gcflags "-N -l"

.PHONY: $(TARGET)
$(TARGET):
	@echo "  >  Building binary $(TARGET) ..."
	@${GO} build ${GO_BUILD_FLAGS} -o $@ ./main.go
