@echo off
REM Deploy to GitHub Script

echo ====================================
echo Deploying URL Shortener to GitHub
echo ====================================
echo.

REM Check if git is installed
where git >nul 2>&1
if %ERRORLEVEL% NEQ 0 (
    echo Error: Git is not installed
    pause
    exit /b 1
)

echo Git is installed
echo.

REM Initialize git if not already done
if not exist .git (
    echo Initializing Git repository...
    git init
    echo Git initialized
) else (
    echo Git repository already exists
)

echo.
echo Adding files to Git...

REM Add all files
git add .

echo Files staged
echo.

REM Commit
echo Creating initial commit...
git commit -m "Initial commit: Full-stack URL Shortener" -m "- Go backend with Gin framework" -m "- React frontend with beautiful UI" -m "- RESTful API with 6 endpoints" -m "- SQLite + in-memory storage" -m "- Click tracking and analytics" -m "- Docker support" -m "- Comprehensive documentation"

echo Commit created
echo.

REM Add remote (if not already added)
git remote | findstr origin >nul 2>&1
if %ERRORLEVEL% NEQ 0 (
    echo Adding GitHub remote...
    git remote add origin https://github.com/ThisuraPerera09/url-shortner-GO.git
    echo Remote added
) else (
    echo Remote already exists
)

echo.
echo Pushing to GitHub...
git branch -M main
git push -u origin main

echo.
echo ====================================
echo Successfully deployed to GitHub!
echo ====================================
echo.
echo Your repository: https://github.com/ThisuraPerera09/url-shortner-GO
echo.
echo Next steps:
echo   1. Visit your GitHub repo
echo   2. Add a description
echo   3. Add topics: golang, react, url-shortener, fullstack
echo   4. Share with the world!
echo.
pause

