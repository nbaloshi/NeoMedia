🔵 PHASE 0 — ENVIRONMENT & FOUNDATION

⏱ 3–5 days
🎯 Goal: You don’t fight your tools later

Step 0.1 – Install global tools
Required

Go ≥ 1.22

Node ≥ 18

PostgreSQL

Git

Backend CLI tools
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest

Step 0.2 – Backend bootstrap
mkdir backend && cd backend
go mod init social-platform
go get github.com/go-chi/chi/v5
go get github.com/jackc/pgx/v5


Create:

cmd/api/main.go


Minimal server:

health check /health

DB connection

graceful shutdown

Step 0.3 – Database setup

Create:

migrations/


First migration:

users table

id (UUID)

email (unique)

password_hash

created_at

Test:

migrate up


✅ STOP HERE
If this feels comfortable, continue.

🔵 PHASE 1 — AUTH & SESSIONS (CRITICAL)

⏱ 1–2 weeks
🎯 Goal: Secure authentication like a real app

Step 1.1 – Passwords

Install:

go get golang.org/x/crypto/bcrypt


Implement:

Hash on register

Compare on login

Step 1.2 – Sessions

Create table:

sessions
- id (UUID)
- user_id
- expires_at


Middleware:

Read session cookie

Load user

Reject if expired

Cookie flags:

HttpOnly

SameSite=Lax

Secure (in prod)

Step 1.3 – CSRF Protection

Generate CSRF token

Send via cookie or header

Validate on POST/PUT/DELETE

Step 1.4 – Frontend setup (NOW)
cd ../
npm create vite@latest frontend -- --template react-ts
cd frontend
npm install


Install frontend libs:

npm install @tanstack/react-query zod
npm install -D tailwindcss postcss autoprefixer


Setup:

Auth pages

Login/Register forms

Cookie-based auth (no tokens!)

✅ Resume value: Secure auth system

🔵 PHASE 2 — SOCIAL CORE

⏱ 1–2 weeks
🎯 Goal: Real CRUD with relationships

Step 2.1 – Data modeling

Add tables:

posts

comments

likes (polymorphic)

Indexes:

post_id

user_id

created_at

Step 2.2 – Backend APIs

Endpoints:

POST /posts

GET /feed

POST /comments

POST /likes

DELETE /likes

Use:

Transactions

Authorization checks

Step 2.3 – Frontend data layer

Install:

npm install @tanstack/react-query-devtools


Use:

useInfiniteQuery for feed

Optimistic likes

Loading & error states

✅ Resume value: Business logic & data integrity

🔵 PHASE 3 — REALTIME CHAT

⏱ 1 week
🎯 Goal: Concurrency + sockets

Step 3.1 – Backend WebSocket

Install:

go get github.com/gorilla/websocket


Implement:

Auth via cookie

User connection map

Message broadcasting

Save messages to DB

Step 3.2 – Frontend WebSocket

Connect on login

Reconnect logic

Message queue

Presence indicator

✅ Resume value: Realtime systems

🔵 PHASE 4 — POLISH & HARDENING

⏱ 1 week
🎯 Goal: Production readiness

Step 4.1 – Security & limits

Add:

Rate limiting (login)

Input validation

Error sanitization

Step 4.2 – Tests

Backend:

Auth tests

Post creation test

Frontend:

Basic component tests (optional)

Step 4.3 – UX polish

Empty states

Skeleton loaders

Toasts

🔵 PHASE 5 — DEPLOYMENT

⏱ 2–3 days
🎯 Goal: Shareable link

Step 5.1 – Backend deploy

Choose:

Fly.io (recommended)

Railway

Setup:

Env vars

DB migrations

HTTPS

Step 5.2 – Frontend deploy

Choose:

Vercel

Netlify

Point frontend to backend URL.

✅ Resume value: Shipped product

🔵 PHASE 6 — INTERVIEW PREP (VERY IMPORTANT)

Prepare answers for:

Why cookies over JWT?

How WebSocket auth works

DB indexing strategy

Scaling chat

Tradeoffs made

Write:

README

Architecture diagram

Known limitations

🎯 FINAL RESULT

When done:

You are hireable

You can defend design choices

You have one serious project instead of many weak ones