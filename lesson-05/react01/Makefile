# Version: 1.0.0

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build -v
GOTESTS=$(GOCMD) test
GOCLEAN=$(GOCMD) clean

VERSION=0.63.0

DIRECTORY_BIN=bin

all: install tests build

install:
	# Build client application
	cd app && npm install --save && npm run build

	# Create public directory
	if [ ! -d "./public" ]; then mkdir -p ./public; fi

	# Copy client application to public directory
	cp -r ./app/dist/* ./public

	# Copy favicon to public directory
	cp ./assets/* ./public

tests:
	# Run tests
	$(GOTESTS) ./...

build:
	# Create bin directory
	if [ ! -d "./$(DIRECTORY_BIN)" ]; then mkdir $(DIRECTORY_BIN); fi

	# Native build
	$(GOBUILD) $(BUILDFLAGS) -o $(DIRECTORY_BIN)/server cmd/*.go
	
	# Linux AMD64 architecture build
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(BUILDFLAGS) -o $(DIRECTORY_BIN)/server.linux-amd64 cmd/*.go

clean:
	rm -rf public
	rm -rf app/node_modules
	rm -rf $(DIRECTORY_BIN)
