#!/bin/bash
# Quick start backend with in-memory storage

echo "========================================"
echo "Starting URL Shortener Backend"
echo "Using In-Memory Storage (No Database)"
echo "========================================"
echo ""

USE_IN_MEMORY=true go run main.go

