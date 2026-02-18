## ⏱ 3–5 days
## 🎯 Goal: You don’t fight your tools later
## Difficulty * *

# Step 0.1 – Install global tools
# Required
1. Go ≥ 1.22
2. Node ≥ 18
3. PostgreSQL
4. Git
5. Backend CLI tools
- go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
- go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Step 0.2 – Backend bootstrap
1. mkdir backend && cd backend
- go mod init social-platform
- go get github.com/go-chi/chi/v5
- go get github.com/jackc/pgx/v5

2. Create:
- cmd/api/main.go

3. Minimal server:
- health check /health
- DB connection
- graceful shutdown

# Step 0.3 – Database setup
1. Create:
- migrations/

2. First migration:
- users table
- id (UUID)
- email (unique)
- password_hash
- created_at

3. Test:
- migrate up