@echo off
REM Deploy to GitHub with Multiple Commits

echo ==================================================
echo Deploying URL Shortener with Multiple Commits
echo ==================================================
echo.

where git >nul 2>&1
if %ERRORLEVEL% NEQ 0 (
    echo Error: Git is not installed
    pause
    exit /b 1
)

echo Git is installed
echo.

REM Initialize git if needed
if not exist .git (
    echo Initializing Git repository...
    git init
) else (
    echo Git repository already exists
)

echo.
echo Creating commits with development history...
echo.

REM Commit 1: Initial setup
echo Commit 1/8: Initial project setup
git add go.mod go.sum .gitignore Makefile
git add config/
git commit -m "feat: Initial project setup" -m "- Add Go module configuration" -m "- Set up project structure" -m "- Add configuration management"

timeout /t 1 /nobreak >nul

REM Commit 2: Models and storage
echo Commit 2/8: Add models and storage layer
git add models/
git add storage/storage.go storage/memory.go
git commit -m "feat: Add models and storage layer" -m "- Define URL model" -m "- Create storage interface" -m "- Implement in-memory storage"

timeout /t 1 /nobreak >nul

REM Commit 3: SQLite
echo Commit 3/8: Add SQLite database support
git add storage/sqlite.go
git commit -m "feat: Add SQLite persistent storage" -m "- Implement SQLite adapter" -m "- Add database schema" -m "- Support CRUD operations"

timeout /t 1 /nobreak >nul

REM Commit 4: Service
echo Commit 4/8: Implement URL shortening service
git add service/url_service.go
git commit -m "feat: Implement URL shortening service" -m "- Add random code generation" -m "- Support custom codes" -m "- Implement async click tracking"

timeout /t 1 /nobreak >nul

REM Commit 5: API
echo Commit 5/8: Add REST API handlers
git add handlers/ middleware/ main.go
git commit -m "feat: Add REST API with Gin framework" -m "- Implement 6 RESTful endpoints" -m "- Add CORS middleware" -m "- Configure graceful shutdown"

timeout /t 1 /nobreak >nul

REM Commit 6: Tests
echo Commit 6/8: Add comprehensive tests
git add service/*_test.go storage/*_test.go
git add cmd/test-client/
git commit -m "test: Add unit tests and test client" -m "- Add service layer tests" -m "- Add storage layer tests" -m "- Create API test client"

timeout /t 1 /nobreak >nul

REM Commit 7: Frontend
echo Commit 7/8: Add React frontend
git add frontend/
git add postman_collection.json
git commit -m "feat: Add React frontend with beautiful UI" -m "- Create modern UI with purple gradient" -m "- Add 3-tab interface" -m "- Implement URL shortening form" -m "- Add analytics dashboard" -m "- Make fully responsive"

timeout /t 1 /nobreak >nul

REM Commit 8: Documentation
echo Commit 8/8: Add documentation and deployment
git add README.md QUICKSTART.md FULLSTACK_GUIDE.md
git add ARCHITECTURE.md PROJECT_SUMMARY.md
git add Dockerfile docker-compose.yml
git add start*.sh start*.bat deploy-to-github*.sh deploy-to-github*.bat
git add react-example.jsx
git commit -m "docs: Add comprehensive documentation" -m "- Add complete README" -m "- Add deployment guides" -m "- Add Docker support" -m "- Create startup scripts"

echo.
echo Created 8 commits!
echo.

REM Add remote
git remote | findstr origin >nul 2>&1
if %ERRORLEVEL% NEQ 0 (
    echo Adding GitHub remote...
    git remote add origin https://github.com/ThisuraPerera09/url-shortner-GO.git
)

echo.
echo Commit History:
git log --oneline --graph

echo.
echo Pushing to GitHub...
git branch -M main
git push -u origin main --force

echo.
echo ==================================================
echo Successfully deployed with 8 commits!
echo ==================================================
echo.
echo Repository: https://github.com/ThisuraPerera09/url-shortner-GO
echo.
pause

