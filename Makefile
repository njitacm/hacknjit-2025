# PHONY does not correspond to a real file; Used to run commands w/o a corresponding file
.PHONY: all update clean

# Default 
all: hacknjit-server

# Install & Update Node.js and golang dependencies
update:
	npm install 
	npm update
	go install
	go get -u ./

# Build frontend and Go backend
hacknjit-server:
	npm install
	go install
	npm run build
	go build -o hacknjit-server main.go

# Clean build artifacts
clean:
	rm -rf node_modules
	go clean
	rm -f hacknjit-server
