#!/bin/bash
# Deploy to GitHub Script

echo "🚀 Deploying URL Shortener to GitHub"
echo "======================================"
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
echo "📝 Adding files to Git..."

# Create .gitignore if it doesn't exist or update it
cat > .gitignore << 'EOF'
# Binaries
*.exe
*.exe~
*.dll
*.so
*.dylib
bin/
dist/

# Test binary
*.test

# Output
*.out

# Dependency directories
vendor/

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
*.swo
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
frontend/yarn-error.log*
EOF

echo "✓ .gitignore updated"

# Add all files
git add .

echo "✓ Files staged"
echo ""

# Commit
echo "💾 Creating initial commit..."
git commit -m "Initial commit: Full-stack URL Shortener

- Go backend with Gin framework
- React frontend with beautiful UI
- RESTful API with 6 endpoints
- SQLite + in-memory storage
- Click tracking and analytics
- Docker support
- Comprehensive documentation

Features:
- URL shortening (random + custom codes)
- URL management and listing
- Detailed analytics dashboard
- Responsive design
- CORS enabled
- Production-ready"

echo "✓ Commit created"
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
echo "🚀 Pushing to GitHub..."
git branch -M main
git push -u origin main

echo ""
echo "======================================"
echo "✅ Successfully deployed to GitHub!"
echo "======================================"
echo ""
echo "🔗 Your repository: https://github.com/ThisuraPerera09/url-shortner-GO"
echo ""
echo "Next steps:"
echo "  1. Visit your GitHub repo"
echo "  2. Add a description"
echo "  3. Add topics: golang, react, url-shortener, fullstack"
echo "  4. Share with the world! 🌍"
echo ""

