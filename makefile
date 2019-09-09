# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOGETU=$(GOGET) -u -v
BUILD_DIR=./build
BINARY_NAME=$(BUILD_DIR)/indy-build
BINARY_UNIX=$(BINARY_NAME)_unix

all: test build
deps:
		$(GOGET) github.com/stretchr/testify
		$(GOGET) github.com/manifoldco/promptui
		$(GOGET) github.com/dustin/go-humanize
		$(GOGET) github.com/spf13/cobra/cobra
		$(GOGETU) golang.org/x/net/proxy
		$(GOGET) gopkg.in/src-d/go-git.v4
build: 
		$(GOBUILD) -o $(BINARY_NAME) -v
test: 
		$(GOTEST) -v ./...
clean: 
		$(GOCLEAN)
		rm -rf $(BUILD_DIR)


# Cross compilation
build-linux:
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
