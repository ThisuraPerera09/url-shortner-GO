# URL Shortener Service 🔗

A production-ready URL shortener built with Go, featuring a RESTful API, SQLite/in-memory storage, click tracking, and containerization support.

## 🌟 Features

- **URL Shortening**: Generate short codes for long URLs
- **Custom Short Codes**: Support for user-defined short codes
- **Click Tracking**: Monitor access statistics for each shortened URL
- **Dual Storage**: Choose between SQLite (persistent) or in-memory storage
- **RESTful API**: Clean, documented API endpoints
- **High Performance**: Built with Go's concurrency model
- **Docker Support**: Easy containerization and deployment
- **CORS Enabled**: Ready for frontend integration

## 🚀 Quick Start

### Prerequisites

- Go 1.21+ (for local development)
- Docker & Docker Compose (for containerized deployment)

### Local Development

1. **Clone the repository**
```bash
git clone <your-repo-url>
cd url-shortener
```

2. **Install dependencies**
```bash
go mod download
```

3. **Configure environment** (optional)
```bash
cp .env.example .env
# Edit .env with your configuration
```

4. **Run the application**
```bash
# With SQLite (persistent storage)
go run main.go

# With in-memory storage
USE_IN_MEMORY=true go run main.go
```

5. **Server will start on** `http://localhost:8080`

### Docker Deployment

1. **Using Docker Compose (recommended)**
```bash
docker-compose up -d
```

2. **Using Docker directly**
```bash
# Build the image
docker build -t url-shortener .

# Run the container
docker run -p 8080:8080 -v $(pwd)/data:/data url-shortener
```

## 📚 API Documentation

### Base URL
```
http://localhost:8080
```

### Endpoints

#### 1. Health Check
Check if the service is running.

```http
GET /api/health
```

**Response:**
```json
{
  "status": "healthy",
  "service": "url-shortener"
}
```

---

#### 2. Shorten URL
Create a shortened URL.

```http
POST /api/shorten
Content-Type: application/json

{
  "url": "https://www.example.com/very/long/url/path",
  "custom_code": "my-link"  // Optional
}
```

**Response:**
```json
{
  "short_code": "my-link",
  "short_url": "http://localhost:8080/my-link",
  "original_url": "https://www.example.com/very/long/url/path"
}
```

**Status Codes:**
- `201 Created` - URL shortened successfully
- `400 Bad Request` - Invalid request body
- `409 Conflict` - Custom code already exists

---

#### 3. Redirect to Original URL
Access a shortened URL and get redirected.

```http
GET /:shortCode
```

**Response:**
- `301 Moved Permanently` - Redirects to original URL
- `404 Not Found` - Short code doesn't exist

**Example:**
```bash
curl -L http://localhost:8080/my-link
# Redirects to: https://www.example.com/very/long/url/path
```

---

#### 4. Get URL Statistics
Retrieve analytics for a shortened URL.

```http
GET /api/stats/:shortCode
```

**Response:**
```json
{
  "short_code": "my-link",
  "original_url": "https://www.example.com/very/long/url/path",
  "clicks": 42,
  "created_at": "2026-01-01T10:00:00Z",
  "last_accessed": "2026-01-01T15:30:00Z"
}
```

---

#### 5. List All URLs
Get a paginated list of all shortened URLs.

```http
GET /api/urls?limit=10&offset=0
```

**Response:**
```json
{
  "urls": [
    {
      "id": 1,
      "short_code": "abc123",
      "original_url": "https://example.com",
      "clicks": 10,
      "created_at": "2026-01-01T10:00:00Z"
    }
  ],
  "limit": 10,
  "offset": 0,
  "count": 1
}
```

---

#### 6. Delete URL
Remove a shortened URL.

```http
DELETE /api/urls/:shortCode
```

**Response:**
```json
{
  "message": "URL deleted successfully"
}
```

**Status Codes:**
- `200 OK` - Deleted successfully
- `404 Not Found` - Short code doesn't exist

---

## 🛠️ Configuration

Environment variables (see `.env.example`):

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8080` | Server port |
| `BASE_URL` | `http://localhost:8080` | Base URL for short links |
| `DATABASE_PATH` | `./urlshortener.db` | SQLite database file path |
| `SHORT_CODE_LEN` | `6` | Length of generated short codes |
| `USE_IN_MEMORY` | `false` | Use in-memory storage instead of SQLite |

---

## 📖 Usage Examples

### Using cURL

**Shorten a URL:**
```bash
curl -X POST http://localhost:8080/api/shorten \
  -H "Content-Type: application/json" \
  -d '{"url": "https://github.com/your-repo"}'
```

**Get statistics:**
```bash
curl http://localhost:8080/api/stats/abc123
```

**List all URLs:**
```bash
curl http://localhost:8080/api/urls?limit=5
```

### Using JavaScript (Frontend)

```javascript
// Shorten URL
async function shortenURL(longUrl) {
  const response = await fetch('http://localhost:8080/api/shorten', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ url: longUrl })
  });
  const data = await response.json();
  return data.short_url;
}

// Get stats
async function getStats(shortCode) {
  const response = await fetch(`http://localhost:8080/api/stats/${shortCode}`);
  return await response.json();
}
```

---

## 🏗️ Architecture

```
url-shortener/
├── config/          # Configuration management
├── handlers/        # HTTP request handlers
├── middleware/      # Custom middleware (logging, CORS)
├── models/          # Data models
├── service/         # Business logic
├── storage/         # Storage layer (interface + implementations)
│   ├── storage.go   # Storage interface
│   ├── memory.go    # In-memory implementation
│   └── sqlite.go    # SQLite implementation
├── main.go          # Application entry point
├── Dockerfile       # Docker configuration
└── docker-compose.yml
```

### Design Patterns

- **Repository Pattern**: Clean separation between business logic and data storage
- **Dependency Injection**: Easy testing and swapping implementations
- **Interface-based Design**: Storage layer can be extended (PostgreSQL, Redis, etc.)

---

## 🧪 Testing

### Manual Testing with Postman/Thunder Client

Import this collection or use the examples above.

### Writing Unit Tests

```go
// Example test structure (create test files as needed)
package service_test

import (
    "testing"
    "url-shortener/service"
    "url-shortener/storage"
)

func TestShortenURL(t *testing.T) {
    store := storage.NewInMemoryStorage()
    svc := service.NewURLService(store, 6)
    
    url, err := svc.ShortenURL("https://example.com", "")
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }
    
    if len(url.ShortCode) != 6 {
        t.Errorf("Expected short code length 6, got %d", len(url.ShortCode))
    }
}
```

---

## 🚀 Deployment

### Production Considerations

1. **Use a reverse proxy** (Nginx, Caddy) for SSL/TLS
2. **Set `BASE_URL`** to your production domain
3. **Back up your database** regularly
4. **Monitor logs** and set up alerts
5. **Consider rate limiting** for abuse prevention

### Example Nginx Configuration

```nginx
server {
    listen 80;
    server_name short.yourdomain.com;

    location / {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

---

## 🎯 Next Steps & Enhancements

### Easy Additions
- [ ] Rate limiting middleware
- [ ] URL expiration dates
- [ ] QR code generation for short URLs
- [ ] Password-protected links

### Advanced Features
- [ ] User authentication & accounts
- [ ] Custom domains
- [ ] Analytics dashboard (React frontend)
- [ ] PostgreSQL support
- [ ] Redis caching layer
- [ ] Microservices architecture

### Learning Extensions
- [ ] Add unit and integration tests
- [ ] Implement CI/CD pipeline
- [ ] Add Prometheus metrics
- [ ] Deploy to Kubernetes
- [ ] Add GraphQL API

---

## 🤝 Contributing

Contributions are welcome! Feel free to:
- Report bugs
- Suggest features
- Submit pull requests

---

## 📝 License

MIT License - feel free to use this project for learning and commercial purposes.

---

## 🙏 Acknowledgments

Built with:
- [Gin](https://github.com/gin-gonic/gin) - HTTP web framework
- [SQLite](https://www.sqlite.org/) - Database
- [Go](https://golang.org/) - Programming language

---

## 📧 Support

For questions or issues, please open a GitHub issue or reach out!

Happy URL shortening! 🎉
