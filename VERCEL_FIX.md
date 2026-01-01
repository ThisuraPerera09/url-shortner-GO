# 🔧 Fix Vercel Deployment

## 🎯 **Best Solution: Configure in Dashboard**

Instead of using `vercel.json`, configure everything in Vercel dashboard (more reliable for monorepos):

### **Step 1: Remove vercel.json (Optional)**

You can delete `vercel.json` - we'll configure in dashboard instead.

### **Step 2: Configure in Vercel Dashboard**

1. Go to your project in Vercel
2. Click **"Settings"** tab
3. Go to **"General"** section

#### **Framework Preset:**
- Set to: **"Create React App"**

#### **Root Directory:**
- Click **"Edit"**
- Set to: **`frontend`**
- Click **"Save"**

#### **Build and Output Settings:**
- **Build Command:** `npm run build` (auto-filled)
- **Output Directory:** `build` (auto-filled)
- **Install Command:** `npm install` (auto-filled)

### **Step 3: Environment Variables**

Go to **"Environment Variables"**:
```
REACT_APP_API_URL = https://YOUR-RAILWAY-APP.up.railway.app/api
```

### **Step 4: Redeploy**

1. Go to **"Deployments"** tab
2. Click **"Redeploy"** on latest deployment
3. ✅ Should work now!

---

## 🔄 **Alternative: If Dashboard Doesn't Work**

If dashboard configuration doesn't work, try this `vercel.json` in root:

```json
{
  "buildCommand": "cd frontend && npm install && npm run build",
  "outputDirectory": "frontend/build",
  "installCommand": "cd frontend && npm install"
}
```

---

## 🆘 **Common Errors & Fixes**

### **Error: "Cannot find package.json"**
- **Fix:** Make sure Root Directory is set to `frontend` in dashboard

### **Error: "Build command failed"**
- **Fix:** Check that `frontend/package.json` exists
- **Fix:** Try setting Build Command to: `cd frontend && npm install && npm run build`

### **Error: "Output directory not found"**
- **Fix:** Make sure Output Directory is `build` (not `frontend/build`) when Root Directory is `frontend`

---

## ✅ **Recommended Settings**

**In Vercel Dashboard → Settings → General:**

```
Framework Preset:     Create React App
Root Directory:       frontend
Build Command:        npm run build
Output Directory:     build
Install Command:      npm install
```

**Environment Variables:**
```
REACT_APP_API_URL = https://your-backend.up.railway.app/api
```

---

**This should fix it! 🚀**

