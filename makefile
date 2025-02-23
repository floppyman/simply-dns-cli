# --- Makefile ----

# This how we want to name the binary output
BINARY=./bin/simply-dns-cli
PROJECT=github.com/umbrella-sh/simply-dns-cli

# These are the values we want to pass for VERSION and BUILD
# git tag 1.0.1
# git commit -am "One more change after the tags"
VERSION := $(shell git describe --tags --match "v*" | sed s/awt-//g)
BUILD_DATE := $(shell date -u +%Y-%m-%d)
DEBUG := false

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-X main.Version=$(VERSION) -X main.BuildDate=$(BUILD_DATE) -X main.Debug=$(DEBUG)"

# Builds the project
linux:
	$(info )
	$(info build)
	$(info --------------------)
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BINARY)

windows:
	$(info )
	$(info build)
	$(info --------------------)
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(BINARY).exe $(PROJECT)

# Cleans our project: deletes binaries
clean:
	$(info )
	$(info clean)
	$(info --------------------)
	if [ -f $(BINARY) ] ; then rm $(BINARY) ; fi
	if [ -f $(BINARY).exe ] ; then rm $(BINARY).exe ; fi
