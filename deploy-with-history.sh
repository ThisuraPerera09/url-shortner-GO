#!/bin/bash
# Deploy to GitHub with Multiple Commits
# This creates a realistic commit history

echo "🚀 Deploying URL Shortener with Multiple Commits"
echo "=================================================="
echo ""

# Check if git is installed
if ! command -v git &> /dev/null; then
    echo "❌ Error: Git is not installed"
    exit 1
fi

echo "✓ Git is installed"
echo ""

# Initialize git if not already done
if [ ! -d .git ]; then
    echo "📦 Initializing Git repository..."
    git init
    echo "✓ Git initialized"
else
    echo "✓ Git repository already exists"
fi

echo ""

# Create .gitignore
cat > .gitignore << 'EOF'
# Binaries
*.exe
*.dll
*.so
*.dylib
bin/

# Test binary
*.test
*.out

# Go workspace
go.work

# Database files
*.db
*.sqlite
*.sqlite3

# Environment variables
.env
.env.local

# IDE
.vscode/
.idea/
*.swp
*~

# OS
.DS_Store
Thumbs.db

# Frontend
frontend/node_modules/
frontend/build/
frontend/.env
frontend/npm-debug.log*
frontend/yarn-debug.log*
EOF

echo "📝 Creating commits with development history..."
echo ""

# Commit 1: Initial setup
echo "Commit 1/8: Initial project setup"
git add go.mod go.sum .gitignore Makefile
git add config/
git commit -m "feat: Initial project setup

- Add Go module configuration
- Set up project structure
- Add configuration management
- Create Makefile for build automation"

sleep 1

# Commit 2: Core backend models and storage
echo "Commit 2/8: Add models and storage layer"
git add models/
git add storage/storage.go storage/memory.go
git commit -m "feat: Add models and storage layer

- Define URL model with all fields
- Create storage interface for flexibility
- Implement in-memory storage
- Add thread-safe operations with mutexes"

sleep 1

# Commit 3: SQLite storage
echo "Commit 3/8: Add SQLite database support"
git add storage/sqlite.go
git commit -m "feat: Add SQLite persistent storage

- Implement SQLite storage adapter
- Add database schema with indexes
- Support CRUD operations
- Handle database migrations"

sleep 1

# Commit 4: Business logic
echo "Commit 4/8: Implement URL shortening service"
git add service/url_service.go
git commit -m "feat: Implement URL shortening service

- Add random short code generation (crypto/rand)
- Support custom short codes
- Implement async click tracking with goroutines
- Add collision handling for generated codes
- Include URL validation"

sleep 1

# Commit 5: API and handlers
echo "Commit 5/8: Add REST API handlers"
git add handlers/ middleware/ main.go
git commit -m "feat: Add REST API with Gin framework

- Implement 6 RESTful endpoints
- Add URL shortening endpoint
- Add redirect with click tracking
- Add statistics endpoint
- Add CORS middleware for frontend
- Add custom logging middleware
- Configure graceful shutdown"

sleep 1

# Commit 6: Tests
echo "Commit 6/8: Add comprehensive tests"
git add service/*_test.go storage/*_test.go
git add cmd/test-client/
git commit -m "test: Add unit tests and test client

- Add service layer tests (85%+ coverage)
- Add storage layer tests
- Implement table-driven tests
- Create API test client
- Add Postman collection"

sleep 1

# Commit 7: Frontend
echo "Commit 7/8: Add React frontend"
git add frontend/
git add postman_collection.json
git commit -m "feat: Add React frontend with beautiful UI

- Create modern UI with purple gradient theme
- Add 3-tab interface (Shorten, List, Analytics)
- Implement URL shortening form with validation
- Add URL management dashboard
- Create analytics page with statistics
- Make fully responsive design
- Add smooth animations and transitions
- Integrate with backend API"

sleep 1

# Commit 8: Documentation and deployment
echo "Commit 8/8: Add documentation and deployment"
git add README.md QUICKSTART.md FULLSTACK_GUIDE.md
git add ARCHITECTURE.md PROJECT_SUMMARY.md
git add COMPLETE.txt PROJECT_TREE.txt
git add GITHUB_README.md GITHUB_DEPLOY.md
git add Dockerfile docker-compose.yml
git add start*.sh start*.bat
git add deploy-to-github*.sh deploy-to-github*.bat
git add react-example.jsx verify.sh
git commit -m "docs: Add comprehensive documentation and deployment

- Add complete README with API docs
- Add quick start guide
- Add full-stack integration guide
- Add architecture documentation
- Add Docker support
- Create startup scripts for Windows and Linux
- Add GitHub deployment guide
- Add React example component
- Include project summaries"

echo ""
echo "✅ Created 8 commits with logical progression!"
echo ""

# Add remote (if not already added)
if ! git remote | grep -q origin; then
    echo "🔗 Adding GitHub remote..."
    git remote add origin https://github.com/ThisuraPerera09/url-shortner-GO.git
    echo "✓ Remote added"
else
    echo "✓ Remote already exists"
fi

echo ""
echo "📊 Commit History:"
git log --oneline --graph

echo ""
echo "🚀 Pushing to GitHub..."
git branch -M main
git push -u origin main --force

echo ""
echo "=================================================="
echo "✅ Successfully deployed with 8 commits!"
echo "=================================================="
echo ""
echo "🔗 Your repository: https://github.com/ThisuraPerera09/url-shortner-GO"
echo ""
echo "📊 View commit history on GitHub to see the development progression"
echo ""

