#!/bin/bash
if [ -f .env ]; then export $(cat .env | xargs); fi

echo "💻 Building on Surface (ARM64) for Droplet (x64)..."

# Build Frontend (Result is just JS/HTML/CSS)
npm run build

# Build Backend (Force x64 Linux target)
echo "🐹 Cross-compiling Go binary..."
GOOS=linux GOARCH=amd64 go build -o hacknjit-server main.go

# Verify the file type (optional but helpful)
# This should say: ELF 64-bit LSB executable, x86-64
file hacknjit-server

echo "📤 Pushing to $DROPLET_IP..."
rsync -avz ./dist/ $DROPLET_USER@$DROPLET_IP:$DROPLET_DEST/dist/
scp ./hacknjit-server $DROPLET_USER@$DROPLET_IP:$DROPLET_DEST/

echo "🚀 Restarting..."
ssh $DROPLET_USER@$DROPLET_IP "sudo systemctl restart hacknjit"