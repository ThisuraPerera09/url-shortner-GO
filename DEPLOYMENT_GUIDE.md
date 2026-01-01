# Railway Deployment Guide

## 🚂 Deploy to Railway.app (FREE!)

### Prerequisites:
- GitHub account
- Railway.app account (sign up at railway.app)
- Your code pushed to GitHub

### Step 1: Push to GitHub
```bash
# If not already done:
./deploy-fresh.sh
```

### Step 2: Deploy Backend to Railway

1. Go to https://railway.app
2. Click "Start a New Project"
3. Select "Deploy from GitHub repo"
4. Choose your repository: `url-shortner-GO`
5. Railway auto-detects it's a Go project!
6. Click "Deploy"

### Step 3: Configure Environment Variables

In Railway dashboard:
- Go to your project
- Click "Variables"
- Add these:

```
PORT=8080
BASE_URL=https://your-app.up.railway.app
USE_IN_MEMORY=true
SHORT_CODE_LEN=6
```

### Step 4: Get Your URL

Railway gives you a URL like:
```
https://url-shortener-production-abc123.up.railway.app
```

Update `BASE_URL` in Railway with this URL!

### Step 5: Deploy Frontend to Vercel

1. Go to https://vercel.com
2. Import your GitHub repo
3. Set build settings:
   - Root Directory: `frontend`
   - Build Command: `npm run build`
   - Output Directory: `build`
4. Add environment variable:
   ```
   REACT_APP_API_URL=https://your-backend.up.railway.app/api
   ```
5. Deploy!

### Your App is Live! 🎉

Backend:  https://your-app.up.railway.app
Frontend: https://your-frontend.vercel.app

---

## 🆓 Cost Breakdown

**Railway:**
- Free $5/month credit
- Resets every month
- More than enough for development!

**Vercel:**
- 100% free for personal projects
- Unlimited bandwidth

**Total Cost: $0/month** ✨

---

## 🔧 Troubleshooting

### Backend not working?
- Check environment variables in Railway
- Make sure `USE_IN_MEMORY=true` (SQLite needs volume on Railway Pro)
- Check logs in Railway dashboard

### Frontend can't connect to backend?
- Update `REACT_APP_API_URL` in Vercel
- Check CORS is enabled in backend (it is!)
- Make sure backend URL is correct

---

## 🌐 Custom Domain (Optional)

### Free Options:
1. **Freenom** - Free .tk, .ml, .ga domains
2. **GitHub Student Pack** - Free .me domain (if you're a student)
3. **Railway subdomain** - Already included!

### Point your domain:
- In your domain registrar (Freenom, etc.)
- Add CNAME record:
  ```
  Type: CNAME
  Name: @
  Value: your-app.up.railway.app
  ```

---

## 📊 Alternative Combinations

### Option A: Everything on Render
Backend + Frontend both on Render.com
- Single platform
- Free tier available

### Option B: Everything on Fly.io
- Great Docker support
- 3 free VMs
- Best performance

### Option C: Railway + Vercel (Recommended!)
- Railway: Backend (Go)
- Vercel: Frontend (React)
- Best of both worlds!

---

## 🚀 Quick Deploy Commands

### For Railway (using CLI):
```bash
# Install Railway CLI
npm i -g @railway/cli

# Login
railway login

# Initialize
railway init

# Deploy
railway up
```

### For Vercel (using CLI):
```bash
# Install Vercel CLI
npm i -g vercel

# Deploy
cd frontend
vercel
```

---

## ✅ Checklist

Before deploying:
- [ ] Code pushed to GitHub
- [ ] Backend works locally
- [ ] Frontend works locally
- [ ] Environment variables documented

After deploying:
- [ ] Backend URL works
- [ ] Frontend URL works
- [ ] Test creating short URL
- [ ] Test redirect works
- [ ] Update BASE_URL in backend config

---

## 🎉 You're Live!

Once deployed, you'll have:
- ✅ Live backend API
- ✅ Beautiful frontend UI
- ✅ Free hosting
- ✅ HTTPS (secure!)
- ✅ Auto-deploy on push to GitHub

Share your app with the world! 🌍

