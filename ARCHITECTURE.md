# 🏗️ URL Shortener Architecture

## System Overview

```
┌─────────────────────────────────────────────────────────────────┐
│                         CLIENT LAYER                            │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐  ┌──────────┐       │
│  │  Browser │  │  cURL    │  │ Postman  │  │  React   │       │
│  │          │  │          │  │          │  │   App    │       │
│  └────┬─────┘  └────┬─────┘  └────┬─────┘  └────┬─────┘       │
└───────┼────────────┼─────────────┼──────────────┼──────────────┘
        │            │             │              │
        └────────────┴─────────────┴──────────────┘
                     │
                     ▼
        ┌────────────────────────┐
        │   HTTP(S) Requests     │
        │   JSON API Calls       │
        └────────────┬───────────┘
                     │
                     ▼
┌─────────────────────────────────────────────────────────────────┐
│                      API LAYER (Gin)                            │
│  ┌──────────────────────────────────────────────────────────┐  │
│  │  Middleware: CORS, Logging, Error Handling               │  │
│  └──────────────────┬───────────────────────────────────────┘  │
│                     │                                           │
│  ┌──────────────────▼───────────────────────────────────────┐  │
│  │  Router & Handlers                                        │  │
│  │  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐     │  │
│  │  │   POST      │  │   GET       │  │   GET       │     │  │
│  │  │  /shorten   │  │  /:code     │  │  /stats/:id │     │  │
│  │  └──────┬──────┘  └──────┬──────┘  └──────┬──────┘     │  │
│  └─────────┼────────────────┼────────────────┼─────────────┘  │
└────────────┼────────────────┼────────────────┼────────────────┘
             │                │                │
             ▼                ▼                ▼
┌─────────────────────────────────────────────────────────────────┐
│                     BUSINESS LOGIC LAYER                        │
│  ┌──────────────────────────────────────────────────────────┐  │
│  │  URL Service                                              │  │
│  │  • Generate short codes (crypto/rand)                     │  │
│  │  • Validate URLs                                          │  │
│  │  • Track clicks (async with goroutines)                  │  │
│  │  • Business rules & validation                           │  │
│  └──────────────────┬───────────────────────────────────────┘  │
└─────────────────────┼──────────────────────────────────────────┘
                      │
                      ▼
┌─────────────────────────────────────────────────────────────────┐
│                    STORAGE INTERFACE                            │
│  ┌──────────────────────────────────────────────────────────┐  │
│  │  Storage Interface (Go interface)                         │  │
│  │  • Save(url) error                                        │  │
│  │  • Get(code) (*URL, error)                               │  │
│  │  • Update(url) error                                     │  │
│  │  • Delete(code) error                                    │  │
│  │  • List(limit, offset) ([]*URL, error)                  │  │
│  └──────────────────┬───────────────────────────────────────┘  │
└─────────────────────┼──────────────────────────────────────────┘
                      │
        ┌─────────────┴─────────────┐
        │                           │
        ▼                           ▼
┌───────────────────┐      ┌────────────────────┐
│  In-Memory Store  │      │   SQLite Store     │
│  ┌─────────────┐  │      │  ┌──────────────┐ │
│  │  map[string]│  │      │  │  urls table  │ │
│  │    *URL     │  │      │  │  • id        │ │
│  │             │  │      │  │  • short_code│ │
│  │  + RWMutex  │  │      │  │  • orig_url  │ │
│  │             │  │      │  │  • clicks    │ │
│  └─────────────┘  │      │  │  • created_at│ │
│                   │      │  └──────────────┘ │
│  Fast, volatile   │      │  Persistent       │
└───────────────────┘      └────────────────────┘
```

## Request Flow Examples

### 1. Shortening a URL

```
Client → POST /api/shorten
         {
           "url": "https://example.com",
           "custom_code": "ex123"
         }
    ↓
Middleware (CORS, Logging)
    ↓
Handler.ShortenURL()
    ↓
Service.ShortenURL(url, code)
    • Validates URL format
    • Generates/validates short code
    • Creates URL model
    ↓
Storage.Save(url)
    • SQLite: INSERT INTO urls...
    • Memory: map[code] = url
    ↓
Response ← 201 Created
{
  "short_code": "ex123",
  "short_url": "http://localhost:8080/ex123",
  "original_url": "https://example.com"
}
```

### 2. Redirecting (Click Tracking)

```
Client → GET /ex123
    ↓
Handler.RedirectURL()
    ↓
Service.GetURL(code)
    • Retrieves URL from storage
    • Increments clicks (async goroutine)
    • Updates last_accessed
    ↓
Storage.Get(code) → returns URL
    ↓
[Async] Storage.Update(url) 
    • Updates clicks in background
    ↓
Response ← 301 Redirect
Location: https://example.com
```

### 3. Getting Statistics

```
Client → GET /api/stats/ex123
    ↓
Handler.GetStats()
    ↓
Service.GetStats(code)
    ↓
Storage.Get(code)
    ↓
Response ← 200 OK
{
  "short_code": "ex123",
  "original_url": "https://example.com",
  "clicks": 42,
  "created_at": "2026-01-01T10:00:00Z",
  "last_accessed": "2026-01-01T15:30:00Z"
}
```

## Data Models

### URL Model
```go
type URL struct {
    ID           int64      // Auto-incrementing primary key
    ShortCode    string     // Unique short code (6 chars)
    OriginalURL  string     // Full original URL
    Clicks       int64      // Number of times accessed
    CreatedAt    time.Time  // Creation timestamp
    LastAccessed *time.Time // Last access time (nullable)
}
```

### Request/Response DTOs
```go
// Shorten request
type ShortenRequest struct {
    URL        string `json:"url" binding:"required,url"`
    CustomCode string `json:"custom_code,omitempty"`
}

// Shorten response
type ShortenResponse struct {
    ShortCode   string `json:"short_code"`
    ShortURL    string `json:"short_url"`
    OriginalURL string `json:"original_url"`
}

// Stats response
type StatsResponse struct {
    ShortCode    string     `json:"short_code"`
    OriginalURL  string     `json:"original_url"`
    Clicks       int64      `json:"clicks"`
    CreatedAt    time.Time  `json:"created_at"`
    LastAccessed *time.Time `json:"last_accessed,omitempty"`
}
```

## Concurrency Model

```
┌─────────────────────────────────────────────────┐
│  HTTP Server (Gin)                              │
│  ┌───────────────────────────────────────────┐ │
│  │  Connection Pool                          │ │
│  │  Each request → New Goroutine             │ │
│  └───────────────────────────────────────────┘ │
└────────┬─────────────────────────────┬──────────┘
         │                             │
    Request 1                      Request 2
    Goroutine 1                    Goroutine 2
         │                             │
         ▼                             ▼
  ┌──────────────┐              ┌──────────────┐
  │  Handler     │              │  Handler     │
  │  + Service   │              │  + Service   │
  └──────┬───────┘              └──────┬───────┘
         │                             │
         ▼                             ▼
  ┌──────────────────────────────────────────┐
  │  Storage (Thread-safe)                   │
  │  • In-memory: sync.RWMutex              │
  │  • SQLite: Database locking             │
  └──────────────────────────────────────────┘

Async Click Tracking:
  User requests redirect
      ↓
  Get URL from storage (sync)
      ↓
  Return redirect immediately
      ↓
  go func() { storage.Update() } ← Background goroutine
```

## Scalability Considerations

### Current Setup (Single Instance)
- **Throughput**: ~10K req/sec
- **Storage**: SQLite (good for < 1M URLs)
- **Concurrency**: Go routines (thousands)

### Scaling Up (Vertical)
```
┌────────────────────┐
│   Load Balancer    │
└─────────┬──────────┘
          │
    ┌─────┴─────┐
    │           │
    ▼           ▼
┌────────┐  ┌────────┐
│ App #1 │  │ App #2 │
└───┬────┘  └───┬────┘
    └───────┬───┘
            ▼
    ┌──────────────┐
    │  PostgreSQL  │
    │  + Redis     │
    └──────────────┘
```

### Scaling Out (Horizontal)
```
       ┌──────────────────┐
       │  CDN / Cloudflare│
       └────────┬─────────┘
                │
       ┌────────▼─────────┐
       │  API Gateway     │
       └────────┬─────────┘
                │
    ┌───────────┼───────────┐
    ▼           ▼           ▼
┌────────┐ ┌────────┐ ┌────────┐
│ App #1 │ │ App #2 │ │ App #3 │
└───┬────┘ └───┬────┘ └───┬────┘
    └──────────┴──────────┘
               │
    ┌──────────┴──────────┐
    │                     │
    ▼                     ▼
┌──────────┐      ┌──────────────┐
│  Redis   │      │  PostgreSQL  │
│  Cache   │      │  (Primary DB)│
└──────────┘      └──────────────┘
```

## Technology Stack

```
┌──────────────────────────────────────────────┐
│  Language & Runtime                          │
│  • Go 1.21+ (compiled, statically typed)    │
└──────────────────────────────────────────────┘

┌──────────────────────────────────────────────┐
│  Web Framework                               │
│  • Gin (HTTP routing & middleware)          │
│  • Fast, minimal, production-ready          │
└──────────────────────────────────────────────┘

┌──────────────────────────────────────────────┐
│  Storage                                     │
│  • SQLite (embedded database)               │
│  • In-memory (development)                  │
│  • Extensible via interface                 │
└──────────────────────────────────────────────┘

┌──────────────────────────────────────────────┐
│  Testing                                     │
│  • Go testing package                       │
│  • Table-driven tests                       │
│  • Unit & integration tests                 │
└──────────────────────────────────────────────┘

┌──────────────────────────────────────────────┐
│  Deployment                                  │
│  • Docker (containerization)                │
│  • Docker Compose (orchestration)           │
│  • Binary deployment (single file)          │
└──────────────────────────────────────────────┘
```

## Design Patterns Used

1. **Repository Pattern**
   - Storage interface abstracts data access
   - Easy to swap implementations

2. **Dependency Injection**
   - Services receive dependencies via constructor
   - Improves testability

3. **Factory Pattern**
   - NewURLService(), NewSQLiteStorage()
   - Encapsulates object creation

4. **Strategy Pattern**
   - Multiple storage strategies (Memory, SQLite)
   - Same interface, different behaviors

5. **MVC-like Architecture**
   - Handlers (Controllers)
   - Service (Business Logic)
   - Storage (Data Access)
   - Models (Data Structures)

## Configuration Management

```
Environment Variables → config.Load()
                              ↓
                       ┌──────────────┐
                       │  Config      │
                       │  • Port      │
                       │  • BaseURL   │
                       │  • DBPath    │
                       │  • CodeLen   │
                       └──────────────┘
                              ↓
              Injected into services at startup
```

## Error Handling Flow

```
Storage Layer
    ↓ returns error
Service Layer
    ↓ wraps/handles error
Handler Layer
    ↓ converts to HTTP response
Client
    ↓ receives JSON error

Example:
ErrNotFound → 404 Not Found
ErrAlreadyExists → 409 Conflict
validation error → 400 Bad Request
other errors → 500 Internal Server Error
```

## Security Considerations

✅ **Current Implementations:**
- URL validation (prevents invalid URLs)
- CORS configured (frontend integration)
- No SQL injection (prepared statements)
- Random short code generation (crypto/rand)

🔒 **Future Enhancements:**
- Rate limiting (prevent abuse)
- Authentication (user accounts)
- HTTPS enforcement
- Input sanitization
- URL blacklisting
- DDoS protection

## Performance Optimizations

1. **Async Click Tracking**
   - Don't wait for DB update on redirect
   - Background goroutine handles update

2. **Database Indexing**
   - Index on short_code for O(1) lookups

3. **Connection Pooling**
   - SQLite connection reuse

4. **Lightweight Framework**
   - Gin is one of fastest Go frameworks

5. **Compiled Binary**
   - Native performance (no VM overhead)

## Monitoring & Observability

**Current:**
- Console logging
- Health check endpoint

**Recommended Additions:**
- Prometheus metrics
- Structured logging (JSON)
- Distributed tracing
- Error tracking (Sentry)
- Performance monitoring

---

This architecture supports:
✅ High concurrency
✅ Easy testing
✅ Simple deployment
✅ Horizontal scaling
✅ Code maintainability

