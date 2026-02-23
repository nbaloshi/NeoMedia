## ⏱ 1–2 weeks
## 🎯 Goal: Secure authentication like a real app
## Difficulty * * * * *

## Step 1.1 – Passwords

# Install:
- go get golang.org/x/crypto/bcrypt

# Implement:
1. Hash on register
2. Compare on login

# Step 1.2 – Sessions
1. Create table:
sessions
- id (UUID)
- user_id
- expires_at

2. Middleware:
- Read session cookie
- Load user
- Reject if expired

3. Cookie flags:
- HttpOnly
- SameSite=Lax
- Secure (in prod)

# Step 1.3 – CSRF Protection
1. Generate CSRF token
2.  Send via cookie or header
3.  Validate on POST/PUT/DELETE

### Cleanup required
- database reset
- goroutine checks
- routing readjustment
- error handling
- redundancy check
- structure check
- file and folder naming
- consider connection pool
- optimization

# Step 1.4 – Frontend setup (NOW)
1. cd ../
- npm create vite@latest frontend -- --template react-ts
- cd frontend
- npm install

2. Install frontend libs:
- npm install @tanstack/react-query zod
- npm install -D tailwindcss postcss autoprefixer

3. Setup:
- Auth pages
- Login/Register forms
- Cookie-based auth (no tokens!)
- ✅ Resume value: Secure auth system