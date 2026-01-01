# 🚀 Vercel Deployment Guide - Frontend

Deploy your React frontend to Vercel in minutes!

---

## 📋 **Prerequisites**

1. ✅ Backend deployed on Railway (or any hosting)
2. ✅ GitHub repository with your code
3. ⬜ Vercel account (free)

---

## 🎯 **Step-by-Step Deployment**

### **Step 1: Sign Up for Vercel**

1. Go to **[vercel.com](https://vercel.com)**
2. Click **"Sign Up"**
3. Choose **"Continue with GitHub"**
4. Authorize Vercel to access your repositories

---

### **Step 2: Import Your Project**

1. In Vercel dashboard, click **"Add New..."** → **"Project"**
2. Find and select: **`url-shortner-GO`** repository
3. Click **"Import"**

---

### **Step 3: Configure Project Settings**

Vercel should auto-detect it's a React app, but verify these settings:

#### **Framework Preset:**
- Select: **"Create React App"** (or "Other" if not listed)

#### **Root Directory:**
- Click **"Edit"** next to Root Directory
- Set to: **`frontend`**
- Click **"Continue"**

#### **Build Settings:**
- **Build Command:** `cd frontend && npm install && npm run build`
- **Output Directory:** `frontend/build`
- (These should auto-fill, but verify)

---

### **Step 4: Add Environment Variable**

**This is critical!** The frontend needs to know where your backend is.

1. In the **"Environment Variables"** section
2. Click **"Add"**
3. Add this variable:

```
Name:  REACT_APP_API_URL
Value: https://YOUR-RAILWAY-APP.up.railway.app/api
```

**Replace `YOUR-RAILWAY-APP` with your actual Railway domain!**

Example:
```
REACT_APP_API_URL = https://url-shortner-go-production.up.railway.app/api
```

⚠️ **Important:** Include `/api` at the end!

---

### **Step 5: Deploy!**

1. Click **"Deploy"**
2. Wait 2-3 minutes for build to complete
3. 🎉 Your app is live!

---

## ✅ **After Deployment**

### **Get Your Frontend URL:**

1. After deployment completes, Vercel shows your URL
2. Example: `url-shortner-go.vercel.app`
3. Click the URL to open your app!

### **Test Your App:**

1. Open your Vercel URL
2. Try shortening a URL
3. Check if it connects to your Railway backend

---

## 🔧 **Troubleshooting**

### **Issue: "Build Failed"**

**Solution:**
- Make sure Root Directory is set to `frontend`
- Check that `frontend/package.json` exists
- Verify build command: `cd frontend && npm install && npm run build`

### **Issue: "Cannot connect to backend"**

**Solution:**
1. Check `REACT_APP_API_URL` is set correctly
2. Make sure it includes `/api` at the end
3. Verify your Railway backend is running
4. Test backend directly: `https://your-backend.up.railway.app/api/health`

### **Issue: "CORS Error"**

**Solution:**
- Your backend should already have CORS enabled
- If not, check `middleware/middleware.go` has CORS middleware

### **Issue: "Blank Page"**

**Solution:**
- Check browser console for errors
- Verify `REACT_APP_API_URL` is correct
- Make sure backend is accessible

---

## 🔄 **Updating Your Deployment**

### **Automatic (Recommended):**

Vercel automatically redeploys when you push to GitHub!

```bash
# Make changes to frontend
git add .
git commit -m "Update frontend"
git push origin main

# Vercel automatically rebuilds! ✅
```

### **Manual Redeploy:**

1. Go to Vercel dashboard
2. Click on your project
3. Go to **"Deployments"** tab
4. Click **"Redeploy"** on latest deployment

---

## 📊 **Project Structure on Vercel**

```
Your Repository (url-shortner-GO)
│
├── frontend/          ← Vercel builds this
│   ├── src/
│   ├── package.json
│   └── build/         ← Vercel serves this
│
├── main.go            ← Not used by Vercel
├── Dockerfile         ← Not used by Vercel
└── vercel.json        ← Vercel config (optional)
```

---

## 🎨 **Custom Domain (Optional)**

### **Add Your Own Domain:**

1. In Vercel project, go to **"Settings"** → **"Domains"**
2. Enter your domain (e.g., `urlshortener.com`)
3. Follow DNS instructions
4. Vercel provides free SSL! ✅

---

## 💰 **Cost**

| Feature | Free Tier |
|---------|-----------|
| Hosting | ✅ Unlimited |
| Bandwidth | ✅ 100GB/month |
| Builds | ✅ Unlimited |
| Custom Domain | ✅ Free SSL |
| **Total** | **$0** 🎉 |

---

## 🔗 **Quick Reference**

### **Environment Variables:**

```
REACT_APP_API_URL = https://your-backend.up.railway.app/api
```

### **Vercel URLs:**

- **Production:** `your-app.vercel.app`
- **Preview:** `your-app-git-branch.vercel.app` (for each branch)

---

## 📝 **Complete Setup Checklist**

- [ ] Sign up on Vercel
- [ ] Import GitHub repository
- [ ] Set Root Directory to `frontend`
- [ ] Add `REACT_APP_API_URL` environment variable
- [ ] Deploy
- [ ] Test frontend URL
- [ ] Verify connection to Railway backend
- [ ] (Optional) Add custom domain

---

## 🎯 **Final Architecture**

```
┌─────────────────────────────────────────────┐
│           USER'S BROWSER                    │
└─────────────────────────────────────────────┘
         │                    │
         │                    │
    ┌────▼────┐         ┌─────▼─────┐
    │ Vercel  │         │  Railway  │
    │Frontend │ ──────► │  Backend  │
    │(React)  │  API    │   (Go)    │
    └─────────┘         └───────────┘
    vercel.app          railway.app
```

---

## 🆘 **Need Help?**

Common issues:
- **Backend not responding:** Check Railway logs
- **CORS errors:** Verify backend CORS middleware
- **Build fails:** Check Root Directory is `frontend`
- **Environment variable not working:** Make sure it starts with `REACT_APP_`

---

**Happy Deploying! 🚀**

