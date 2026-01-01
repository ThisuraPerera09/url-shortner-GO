@echo off
REM Quick start backend with in-memory storage (no CGO needed)

echo ========================================
echo Starting URL Shortener Backend
echo Using In-Memory Storage (No Database)
echo ========================================
echo.

set USE_IN_MEMORY=true
go run main.go

