# PHONY does not correspond to a real file; Used to run commands w/o a corresponding file
.PHONY: all install build clean dev

# Default 
all: install build

# Install Node.js and golang dependencies
install:
	npm install
	go install

# Build frontend and Go backend
build: install
	npm run build
	go build -o hacknjit-server main.go

# Clean build artifacts
clean:
	rm -rf node_modules
	go clean
	rm -f hacknjit-server

# Setup a development server for the backend
dev: install build
	./hacknjit-server -dev
