#!/bin/bash
# Verification script to ensure everything is set up correctly

echo "🔍 Verifying URL Shortener Project Setup"
echo "=========================================="
echo ""

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Check Go installation
echo -n "Checking Go installation... "
if command -v go &> /dev/null; then
    echo -e "${GREEN}✓${NC} $(go version)"
else
    echo -e "${RED}✗${NC} Go not found"
    exit 1
fi

# Check project structure
echo -n "Checking project structure... "
required_dirs=("config" "handlers" "middleware" "models" "service" "storage")
all_exist=true
for dir in "${required_dirs[@]}"; do
    if [ ! -d "$dir" ]; then
        echo -e "${RED}✗${NC} Missing directory: $dir"
        all_exist=false
    fi
done
if [ "$all_exist" = true ]; then
    echo -e "${GREEN}✓${NC}"
fi

# Check dependencies
echo -n "Checking dependencies... "
if go list -m all &> /dev/null; then
    echo -e "${GREEN}✓${NC}"
else
    echo -e "${YELLOW}⚠${NC} Running go mod download..."
    go mod download
fi

# Run tests
echo ""
echo "Running tests..."
if go test ./service/... ./storage/... -v; then
    echo -e "${GREEN}✓${NC} All tests passed"
else
    echo -e "${RED}✗${NC} Some tests failed"
fi

# Try to build
echo ""
echo -n "Building application... "
if go build -o bin/url-shortener main.go; then
    echo -e "${GREEN}✓${NC}"
else
    echo -e "${RED}✗${NC} Build failed"
    exit 1
fi

# Summary
echo ""
echo "=========================================="
echo -e "${GREEN}✅ Project verification complete!${NC}"
echo ""
echo "Next steps:"
echo "  1. Start the server: ./start.sh (or start.bat on Windows)"
echo "  2. Test the API: go run cmd/test-client/main.go"
echo "  3. Read the docs: README.md or PROJECT_SUMMARY.md"
echo ""
echo "Happy coding! 🚀"

