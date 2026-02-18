## ⏱ 1–2 weeks
## 🎯 Goal: Real CRUD with relationships
## Difficulty * * * *

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