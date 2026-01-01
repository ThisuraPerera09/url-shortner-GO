# 🚀 Deployment Guide

This guide covers deploying your URL Shortener to various platforms.

---

## 📊 **Storage Options Explained**

### **Option 1: In-Memory Storage (Current Default)**
- ✅ Fast and simple
- ✅ No database setup needed
- ❌ **Data lost on restart** (happens daily on free hosting)
- 🎯 **Use for:** Quick demos, testing deployment

### **Option 2: PostgreSQL (Recommended for Production)**
- ✅ Data persists across restarts
- ✅ Free tier available on Railway
- ✅ Production-ready
- 🎯 **Use for:** Real applications

---

## 🚂 **Deploy to Railway (Recommended)**

Railway offers **free PostgreSQL database** and is perfect for Go apps!

### **Step 1: Sign Up**
1. Go to [Railway.app](https://railway.app)
2. Sign in with GitHub

### **Step 2: Deploy Backend**

#### **2a. Deploy with PostgreSQL (Recommended)**

```bash
# 1. Push your code to GitHub (if not already done)
git add .
git commit -m "Ready for deployment"
git push origin main

# 2. On Railway Dashboard:
#    - Click "New Project"
#    - Select "Deploy from GitHub repo"
#    - Choose your repository
#    - Railway will auto-detect it's a Go app!

# 3. Add PostgreSQL Database:
#    - Click "+ New" → "Database" → "Add PostgreSQL"
#    - Railway automatically sets DATABASE_URL environment variable
#    - Your app will detect and use it automatically!

# 4. Configure Environment Variables:
#    - Click on your backend service
#    - Go to "Variables" tab
#    - Add:
PORT=8080
BASE_URL=https://your-app-name.up.railway.app
```

**That's it!** Your app will:
- Detect `DATABASE_URL` and use PostgreSQL automatically
- Create database tables on first run
- Persist all URLs across restarts

#### **2b. Deploy with In-Memory (Quick Test)**

If you just want to test deployment quickly:

```bash
# On Railway, just add:
PORT=8080
BASE_URL=https://your-app-name.up.railway.app
USE_IN_MEMORY=true

# Note: All URLs will be lost when Railway restarts your app (happens daily)
```

### **Step 3: Deploy Frontend**

Railway can host the React frontend too, or use Vercel (easier):

#### **Option A: Vercel (Recommended for Frontend)**

```bash
# 1. Install Vercel CLI
npm install -g vercel

# 2. Deploy frontend
cd frontend
vercel

# 3. Set environment variable in Vercel dashboard:
REACT_APP_API_URL=https://your-backend.up.railway.app
```

#### **Option B: Railway Frontend**

```bash
# 1. In Railway, click "+ New" → "GitHub Repo"
# 2. Set Root Directory to: frontend
# 3. Add environment variable:
REACT_APP_API_URL=https://your-backend.up.railway.app
```

---

## 🌐 **Deploy to Render.com**

Another great free hosting option!

### **Backend with PostgreSQL:**

```bash
# 1. Push to GitHub
git push origin main

# 2. Go to Render Dashboard:
#    - New → Web Service
#    - Connect your GitHub repo
#    - Settings:
Build Command: go build -o main
Start Command: ./main
Environment: Go

# 3. Add PostgreSQL:
#    - Dashboard → New → PostgreSQL
#    - Copy Internal Database URL

# 4. Environment Variables:
DATABASE_URL=your_postgres_url
PORT=8080
BASE_URL=https://your-app.onrender.com
```

### **Frontend on Render:**

```bash
# Settings:
Build Command: cd frontend && npm install && npm run build
Start Command: cd frontend && npm start
Environment: Node

# Environment Variable:
REACT_APP_API_URL=https://your-backend.onrender.com
```

---

## ☁️ **Deploy to Fly.io**

Great for PostgreSQL apps!

```bash
# 1. Install flyctl
curl -L https://fly.io/install.sh | sh

# 2. Login
fly auth login

# 3. Create app
fly launch

# 4. Create PostgreSQL database
fly postgres create

# 5. Attach database to app
fly postgres attach <postgres-app-name>

# 6. Deploy
fly deploy

# Fly.io automatically sets DATABASE_URL!
```

---

## 🐳 **Deploy with Docker**

You already have a Dockerfile!

### **Docker Hub + Any Cloud:**

```bash
# 1. Build image
docker build -t your-username/url-shortener .

# 2. Push to Docker Hub
docker push your-username/url-shortener

# 3. Deploy to any cloud that supports Docker
#    (AWS, GCP, Azure, DigitalOcean, etc.)
```

### **Run with PostgreSQL locally:**

```bash
# docker-compose.yml (already in your project!)
docker-compose up -d

# This starts:
# - Backend on http://localhost:8080
# - PostgreSQL database
```

---

## 🔧 **Environment Variables Reference**

### **Backend:**
```bash
# Database (pick one)
DATABASE_URL=postgres://user:pass@host:5432/dbname  # Use PostgreSQL
USE_IN_MEMORY=true                                   # Use in-memory (no DB)

# Server config
PORT=8080                                            # Server port
BASE_URL=https://your-domain.com                    # Your deployed URL

# Optional
SHORT_CODE_LEN=6                                    # Length of short codes
```

### **Frontend:**
```bash
REACT_APP_API_URL=https://your-backend.com          # Backend API URL
```

---

## ✅ **Checklist for Production**

- [ ] Use PostgreSQL (not in-memory)
- [ ] Set correct `BASE_URL` environment variable
- [ ] Enable HTTPS (most platforms do this automatically)
- [ ] Set up custom domain (optional)
- [ ] Monitor database size (free tiers have limits)
- [ ] Set up backups (for paid tiers)

---

## 🔍 **Testing Your Deployment**

After deploying:

```bash
# Test API health
curl https://your-app.com/api/health

# Shorten a URL
curl -X POST https://your-app.com/api/shorten \
  -H "Content-Type: application/json" \
  -d '{"url": "https://google.com"}'

# Test redirect (in browser)
https://your-app.com/abc123

# Open frontend
https://your-frontend.com
```

---

## 💡 **Cost Breakdown**

| Platform | Backend | Database | Free Tier |
|----------|---------|----------|-----------|
| **Railway** | Free | PostgreSQL Free | ✅ 500 hrs/month |
| **Render** | Free | PostgreSQL Free | ✅ 750 hrs/month |
| **Fly.io** | Free | PostgreSQL Free | ✅ 3 VMs free |
| **Vercel** | - | - | ✅ Unlimited (frontend) |

**Recommended Combo:**
- Backend + PostgreSQL: **Railway** (easiest setup)
- Frontend: **Vercel** (fastest, best for React)
- **Total Cost: $0** 🎉

---

## 🆘 **Troubleshooting**

### **"Connection refused" errors:**
- Make sure backend is deployed and running first
- Update `REACT_APP_API_URL` in frontend settings

### **"Data disappears after restart":**
- You're using in-memory storage
- Switch to PostgreSQL (see above)

### **"Database connection failed":**
- Check `DATABASE_URL` is set correctly
- Make sure database and app are in same region

### **"Build failed" on Railway:**
- Check `go.mod` is in repository root
- Make sure `go.sum` is committed

---

## 📚 **Next Steps**

1. **Add Custom Domain:**
   - Most platforms support free SSL with custom domains
   - Just point your domain's DNS to their servers

2. **Set Up Monitoring:**
   - Railway/Render provide basic logs
   - Consider adding Sentry for error tracking

3. **Scale Up:**
   - Free tiers handle ~10K requests/day
   - Upgrade to paid tier for more capacity

---

Happy Deploying! 🚀

