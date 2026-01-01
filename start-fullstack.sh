#!/bin/bash
# Full Stack URL Shortener - Startup Script

echo "🚀 Starting Full Stack URL Shortener"
echo "====================================="
echo ""

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Error: Go is not installed"
    exit 1
fi

# Check if Node is installed
if ! command -v node &> /dev/null; then
    echo "❌ Error: Node.js is not installed"
    exit 1
fi

echo "✓ Go and Node.js are installed"
echo ""

# Start backend
echo "📡 Starting Backend API on http://localhost:8080"
cd "$(dirname "$0")"
go run main.go &
BACKEND_PID=$!

# Wait for backend to start
sleep 3

# Start frontend
echo "🎨 Starting Frontend on http://localhost:3000"
cd frontend
npm start &
FRONTEND_PID=$!

echo ""
echo "✅ Application Started!"
echo "====================================="
echo "📡 Backend:  http://localhost:8080"
echo "🎨 Frontend: http://localhost:3000"
echo "====================================="
echo ""
echo "Press Ctrl+C to stop both servers"
echo ""

# Wait for Ctrl+C
trap "echo ''; echo '🛑 Stopping servers...'; kill $BACKEND_PID $FRONTEND_PID; exit" INT
wait

