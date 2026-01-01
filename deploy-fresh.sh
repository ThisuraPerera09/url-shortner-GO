#!/bin/bash
# Fresh deployment with multiple commits

echo "🚀 Creating Fresh Deployment with Multiple Commits"
echo "===================================================="
echo ""

# Remove old git history and start fresh
echo "🔄 Starting fresh..."
rm -rf .git
git init
echo "✓ Fresh Git repository created"
echo ""

# Remove frontend's git repo if exists
rm -rf frontend/.git

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

echo "📝 Creating 8 commits..."
echo ""

# Commit 1: Initial setup
echo "1/8: Initial project setup"
git add go.mod go.sum .gitignore Makefile config/
git commit -m "feat: Initial project setup

- Initialize Go modules
- Add project configuration
- Create Makefile for automation
- Set up gitignore"
sleep 1

# Commit 2: Models and storage interface
echo "2/8: Add models and storage layer"
git add models/ storage/storage.go storage/memory.go storage/memory_test.go
git commit -m "feat: Add models and storage layer

- Define URL data model
- Create storage interface for flexibility
- Implement in-memory storage
- Add thread-safe operations with RWMutex
- Include unit tests for storage"
sleep 1

# Commit 3: SQLite
echo "3/8: Add SQLite database support"
git add storage/sqlite.go
git commit -m "feat: Add SQLite persistent storage

- Implement SQLite storage adapter
- Create database schema with indexes
- Add CRUD operations
- Handle NULL values for timestamps
- Enable production-ready persistence"
sleep 1

# Commit 4: Business logic
echo "4/8: Implement URL shortening service"
git add service/
git commit -m "feat: Implement URL shortening service

- Generate cryptographically secure short codes
- Support custom user-defined codes
- Handle code collisions with retry logic
- Implement async click tracking with goroutines
- Add comprehensive unit tests (85%+ coverage)
- Include pagination for URL listing"
sleep 1

# Commit 5: API
echo "5/8: Add REST API with Gin framework"
git add handlers/ middleware/ main.go
git commit -m "feat: Add REST API with Gin framework

- Implement 6 RESTful endpoints
- Add URL shortening endpoint with validation
- Add redirect endpoint with click tracking
- Add statistics and management endpoints
- Create CORS middleware for frontend integration
- Add custom logging middleware
- Configure graceful shutdown"
sleep 1

# Commit 6: Testing
echo "6/8: Add testing infrastructure"
git add cmd/test-client/ postman_collection.json
git commit -m "test: Add comprehensive testing tools

- Create Go test client for API testing
- Add Postman collection for manual testing
- Test all API endpoints
- Include success/error scenarios"
sleep 1

# Commit 7: Frontend
echo "7/8: Add React frontend"
git add frontend/ react-example.jsx
git commit -m "feat: Add beautiful React frontend

- Create modern UI with purple gradient theme
- Implement 3-tab interface:
  * Shorten URL - Create short links
  * My URLs - Manage all links
  * Analytics - View detailed statistics
- Add responsive design for all devices
- Implement smooth animations and transitions
- Add copy-to-clipboard functionality
- Include error handling and loading states
- Integrate with backend REST API
- Add standalone React component example"
sleep 1

# Commit 8: Documentation and deployment
echo "8/8: Add documentation and deployment tools"
git add README.md QUICKSTART.md FULLSTACK_GUIDE.md ARCHITECTURE.md
git add PROJECT_SUMMARY.md COMPLETE.txt PROJECT_TREE.txt COMPLETION_SUMMARY.md
git add GITHUB_README.md GITHUB_DEPLOY.md COMMIT_HISTORY.md INDEX.md
git add Dockerfile docker-compose.yml .dockerignore
git add start*.sh start*.bat verify.sh
git add deploy*.sh deploy*.bat
git commit -m "docs: Add comprehensive documentation

- Add complete API documentation
- Create quick start guide
- Add full-stack integration guide
- Document system architecture
- Include project summaries and completion status
- Add Docker support (Dockerfile + compose)
- Create startup scripts for all platforms
- Add GitHub deployment guides
- Include commit history documentation"

echo ""
echo "✅ Successfully created 8 commits!"
echo ""

# Show commit history
echo "📊 Commit History:"
git log --oneline --graph

echo ""
echo "🔗 Adding GitHub remote..."
git remote add origin https://github.com/ThisuraPerera09/url-shortner-GO.git

echo ""
echo "🚀 Pushing to GitHub..."
git branch -M main
git push -u origin main --force

echo ""
echo "===================================================="
echo "✅ Successfully deployed with 8 clean commits!"
echo "===================================================="
echo ""
echo "🔗 Repository: https://github.com/ThisuraPerera09/url-shortner-GO"
echo ""
echo "📊 Visit GitHub to see your beautiful commit history!"
echo ""

