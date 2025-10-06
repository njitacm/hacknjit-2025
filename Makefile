# PHONY does not correspond to a real file; Used to run commands w/o a corresponding file
.PHONY: all install clean 

# Default 
all: hacknjit-server

# Install Node.js and golang dependencies
install:
	npm install
	go install

# Build frontend and Go backend
hacknjit-server: install
	npm run build
	go build -o hacknjit-server main.go

# Clean build artifacts
clean:
	rm -rf node_modules
	go clean
	rm -f hacknjit-server
