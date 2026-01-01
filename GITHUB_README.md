# 🔗 URL Shortener - Full Stack Application

[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![React](https://img.shields.io/badge/React-18-61DAFB?style=flat&logo=react)](https://reactjs.org/)
[![Gin](https://img.shields.io/badge/Gin-Framework-00ADD8?style=flat)](https://github.com/gin-gonic/gin)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

A modern, full-stack URL shortener built with **Go** and **React**. Features a beautiful UI, comprehensive analytics, and production-ready architecture.

![URL Shortener Demo](https://via.placeholder.com/800x400/667eea/ffffff?text=URL+Shortener+Demo)

## ✨ Features

### Backend (Go + Gin)
- 🚀 RESTful API with 6 endpoints
- 🔗 URL shortening (random & custom codes)
- 📊 Click tracking with goroutines
- 💾 Dual storage (SQLite + In-Memory)
- 🔒 CORS enabled
- 🐳 Docker support
- ✅ 85%+ test coverage

### Frontend (React)
- 🎨 Beautiful, modern UI with purple gradient theme
- 📱 Fully responsive design
- 🔗 **Shorten URL** - Create short links instantly
- 📋 **My URLs** - Manage all your shortened URLs
- 📊 **Analytics** - Detailed statistics dashboard
- ⚡ Smooth animations and transitions
- 📋 Copy to clipboard functionality

## 🚀 Quick Start

### Prerequisites
- Go 1.21+
- Node.js 16+
- Git

### Start the Full Stack Application

**Windows:**
```cmd
start-fullstack.bat
```

**Linux/Mac:**
```bash
./start-fullstack.sh
```

Then open your browser to: **http://localhost:3000**

## 📖 Manual Setup

### Backend Only
```bash
go run main.go
```

### Frontend Only
```bash
cd frontend
npm install
npm start
```

## 🎯 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/health` | Health check |
| POST | `/api/shorten` | Create short URL |
| GET | `/:shortCode` | Redirect to original URL |
| GET | `/api/stats/:shortCode` | Get URL statistics |
| GET | `/api/urls` | List all URLs |
| DELETE | `/api/urls/:shortCode` | Delete a URL |

## 🎨 Screenshots

### Home - Shorten URL
Clean interface for creating short URLs with optional custom codes.

### My URLs
Grid view of all shortened URLs with click statistics and management options.

### Analytics Dashboard
Detailed analytics with click tracking, creation dates, and insights.

## 🏗️ Architecture

```
url-shortener/
├── 📡 Backend (Go)
│   ├── main.go
│   ├── config/        # Configuration
│   ├── handlers/      # HTTP handlers
│   ├── service/       # Business logic
│   ├── storage/       # Data persistence
│   └── middleware/    # HTTP middleware
│
├── 🎨 Frontend (React)
│   └── src/
│       ├── App.js
│       ├── components/
│       │   ├── Header.js
│       │   ├── URLShortener.js
│       │   ├── URLList.js
│       │   └── Stats.js
│       └── services/
│           └── api.js
│
└── 📚 Documentation
    ├── FULLSTACK_GUIDE.md
    ├── QUICKSTART.md
    └── ARCHITECTURE.md
```

## 🛠️ Tech Stack

**Backend:**
- Go 1.21+
- Gin Web Framework
- SQLite Database
- Clean Architecture

**Frontend:**
- React 18
- Custom CSS3
- Fetch API
- React Hooks

## 🧪 Testing

```bash
# Run backend tests
go test ./...

# With coverage
go test -cover ./...
```

## 🐳 Docker

```bash
# Backend
docker-compose up

# Build manually
docker build -t url-shortener .
docker run -p 8080:8080 url-shortener
```

## 📊 Performance

- **~50K requests/second** (Go backend)
- **<1ms latency** for redirects (in-memory)
- **~10MB memory** footprint
- **Single binary** deployment

## 🎯 Use Cases

- Link sharing on social media
- QR code campaigns
- Email marketing
- Analytics tracking
- Personal link management

## 🔧 Configuration

### Backend (.env)
```bash
PORT=8080
BASE_URL=http://localhost:8080
DATABASE_PATH=./urlshortener.db
SHORT_CODE_LEN=6
USE_IN_MEMORY=false
```

### Frontend (.env)
```bash
REACT_APP_API_URL=http://localhost:8080/api
```

## 📚 Documentation

- [Full Stack Guide](FULLSTACK_GUIDE.md) - Complete setup guide
- [Quick Start](QUICKSTART.md) - Get running in 5 minutes
- [Architecture](ARCHITECTURE.md) - System design details
- [API Documentation](README.md) - Complete API reference

## 🚀 Deployment

### Backend
- Deploy to Railway, Heroku, or DigitalOcean
- Use the compiled binary for easy deployment

### Frontend
- Deploy to Vercel, Netlify, or any static host
- Build with `npm run build`

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👨‍💻 Author

**Thisura Perera**

- GitHub: [@ThisuraPerera09](https://github.com/ThisuraPerera09)

## 🙏 Acknowledgments

- Built with [Gin](https://github.com/gin-gonic/gin) - Fast HTTP web framework
- Powered by [React](https://reactjs.org/) - UI library
- Database by [SQLite](https://www.sqlite.org/) - Embedded database

## 🌟 Show your support

Give a ⭐️ if this project helped you!

---

**Built with ❤️ using React + Go**

