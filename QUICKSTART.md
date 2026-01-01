# URL Shortener - Quick Start Guide

## 🎯 What You'll Build

A complete URL shortening service with:
- RESTful API for creating and managing short URLs
- Click tracking and analytics
- Choice of in-memory or persistent (SQLite) storage
- Production-ready with Docker support

## 🚦 Getting Started (5 Minutes)

### Step 1: Install Dependencies

```bash
# Make sure you have Go installed (go version)
go mod download
```

### Step 2: Run the Server

```bash
# Quick start with in-memory storage
go run main.go

# Or with persistent storage
# (Creates urlshortener.db file)
USE_IN_MEMORY=false go run main.go
```

The server starts at `http://localhost:8080`

### Step 3: Test It Out

**Shorten your first URL:**
```bash
curl -X POST http://localhost:8080/api/shorten \
  -H "Content-Type: application/json" \
  -d '{"url": "https://github.com"}'
```

**Response:**
```json
{
  "short_code": "a1b2c3",
  "short_url": "http://localhost:8080/a1b2c3",
  "original_url": "https://github.com"
}
```

**Try the redirect:**
```bash
# Opens GitHub in your browser
curl -L http://localhost:8080/a1b2c3
```

**Check statistics:**
```bash
curl http://localhost:8080/api/stats/a1b2c3
```

## 📚 Core API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/shorten` | Create short URL |
| GET | `/:code` | Redirect to original URL |
| GET | `/api/stats/:code` | Get click statistics |
| GET | `/api/urls` | List all URLs |
| DELETE | `/api/urls/:code` | Delete a URL |

## 🎨 Connecting to React Frontend

```javascript
// Example React component
import { useState } from 'react';

function URLShortener() {
  const [url, setUrl] = useState('');
  const [shortUrl, setShortUrl] = useState('');

  const shortenURL = async () => {
    const response = await fetch('http://localhost:8080/api/shorten', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ url })
    });
    const data = await response.json();
    setShortUrl(data.short_url);
  };

  return (
    <div>
      <input 
        value={url} 
        onChange={(e) => setUrl(e.target.value)} 
        placeholder="Enter long URL"
      />
      <button onClick={shortenURL}>Shorten</button>
      {shortUrl && <p>Short URL: <a href={shortUrl}>{shortUrl}</a></p>}
    </div>
  );
}
```

## 🐳 Docker Quick Start

```bash
# Build and run with Docker Compose
docker-compose up -d

# View logs
docker-compose logs -f

# Stop
docker-compose down
```

## 🔧 Configuration Options

Create a `.env` file:
```bash
PORT=8080
BASE_URL=http://localhost:8080
DATABASE_PATH=./urlshortener.db
SHORT_CODE_LEN=6
USE_IN_MEMORY=false
```

## 📖 Learning Path

1. **Start Simple**: Run with in-memory storage, test with cURL
2. **Add Persistence**: Switch to SQLite storage
3. **Frontend Integration**: Connect your React app
4. **Deploy**: Use Docker and deploy to cloud
5. **Enhance**: Add features like expiration, rate limiting

## 🎯 Next Steps

- [ ] Test all API endpoints with Postman/cURL
- [ ] Build a React frontend
- [ ] Add custom short codes
- [ ] Implement rate limiting
- [ ] Deploy to production

See `README.md` for complete documentation!

