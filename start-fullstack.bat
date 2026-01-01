@echo off
REM Full Stack URL Shortener - Windows Startup Script

echo ================================
echo Starting Full Stack URL Shortener
echo ================================
echo.

REM Check Go
where go >nul 2>&1
if %ERRORLEVEL% NEQ 0 (
    echo Error: Go is not installed
    pause
    exit /b 1
)

REM Check Node
where node >nul 2>&1
if %ERRORLEVEL% NEQ 0 (
    echo Error: Node.js is not installed
    pause
    exit /b 1
)

echo Backend API: http://localhost:8080
echo Frontend UI:  http://localhost:3000
echo.
echo Starting servers...
echo.

REM Start backend in new window
start "URL Shortener - Backend" cmd /k "go run main.go"

REM Wait a bit for backend to start
timeout /t 3 /nobreak >nul

REM Start frontend in new window
start "URL Shortener - Frontend" cmd /k "cd frontend && npm start"

echo.
echo ================================
echo Application Started!
echo ================================
echo Backend:  http://localhost:8080
echo Frontend: http://localhost:3000
echo ================================
echo.
echo Close the command windows to stop the servers
echo.
pause

