#!/bin/bash
# URL Shortener - Linux/Mac Startup Script

echo "Starting URL Shortener Service..."
echo ""

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Error: Go is not installed or not in PATH"
    echo "Please install Go from https://golang.org/dl/"
    exit 1
fi

# Build the application
echo "Building application..."
go build -o bin/url-shortener main.go
if [ $? -ne 0 ]; then
    echo "Build failed!"
    exit 1
fi

echo ""
echo "Build successful!"
echo ""
echo "Starting server on http://localhost:8080"
echo "Press Ctrl+C to stop the server"
echo ""

# Run the application
./bin/url-shortener

