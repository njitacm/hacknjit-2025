# Run this build target regardless of if file exists or not
.PHONY: hacknjit-server

# Default 
all: hacknjit-server

front-end:
	npm run build

# Build frontend and Go backend
hacknjit-server: front-end
	go build -o hacknjit-server main.go

# Install Node.js and golang packages 
install:
	npm install 
	go install

# Install & Update Node.js and golang dependencies
update: install
	npm update
	go get -u ./

# Clean build artifacts
clean:
	rm -rf node_modules
	go clean
	rm -f hacknjit-server
