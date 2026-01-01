@echo off
REM URL Shortener - Windows Startup Script

echo Starting URL Shortener Service...
echo.

REM Check if Go is installed
go version >nul 2>&1
if errorlevel 1 (
    echo Error: Go is not installed or not in PATH
    echo Please install Go from https://golang.org/dl/
    pause
    exit /b 1
)

REM Build the application
echo Building application...
go build -o bin\url-shortener.exe main.go
if errorlevel 1 (
    echo Build failed!
    pause
    exit /b 1
)

echo.
echo Build successful!
echo.
echo Starting server on http://localhost:8080
echo Press Ctrl+C to stop the server
echo.

REM Run the application
bin\url-shortener.exe

