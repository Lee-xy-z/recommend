RECOMMEND_PATH=github.com/Lee-xy-z/recommend
RECOMMEND_BINARY_NAME=recommend

GOBUILD=go build

GIT_SHA=$(shell git rev-parse HEAD)
GIT_CLOSEST_TAG=$(shell git describe --abbrev=0 --tags)
DATE=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
BUILD_INFO_IMPORT_PATH=$(RECOMMEND_PATH)/pkg/version
BUILD_INFO=-ldflags "-X $(BUILD_INFO_IMPORT_PATH).commitSHA=$(GIT_SHA) -X $(BUILD_INFO_IMPORT_PATH).latestVersion=$(GIT_CLOSEST_TAG) -X $(BUILD_INFO_IMPORT_PATH).date=$(DATE)"

.PYNOY: build-recommend
build-recommend:
	$(GOBUILD) -o $(RECOMMEND_BINARY_NAME) $(BUILD_INFO) ./cmd/recommend/main.go