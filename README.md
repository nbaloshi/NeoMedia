# NeoMedia

NeoMedia is a full-stack social media platform built from scratch as a personal project.

The project focuses on building a complete web application with a Go backend, PostgreSQL database, and React/TypeScript frontend.

## Features

* User registration and login
* Session-based authentication
* Secure password hashing
* Session cookies
* Protected routes
* Create and view posts
* Create and view comments
* Like and unlike posts
* Like and unlike comments
* User session management
* PostgreSQL database with SQL migrations
* Client-side data fetching and caching with TanStack React Query
* Responsive frontend built with Tailwind CSS

## Tech Stack

### Backend

* Go
* Chi
* PostgreSQL
* pgx
* SQL migrations
* HTTP APIs
* Cookie-based sessions
* Password hashing

### Frontend

* React
* TypeScript
* Vite
* React Router
* TanStack React Query
* Tailwind CSS
* Zod

## Architecture

NeoMedia is structured as a separate frontend and backend application.

```text
NeoMedia
├── backend
│   ├── cmd
│   ├── db
│   │   └── migrations
│   └── internal
│       ├── handlers
│       ├── routes
│       ├── sessions
│       └── utils
│
└── frontend
    └── src
        ├── api
        ├── components
        ├── context
        ├── pages
        └── utils
```

The Go backend handles HTTP requests, authentication, sessions, database operations, and application logic. The React frontend communicates with the backend through HTTP APIs.

## Database

NeoMedia uses PostgreSQL with versioned SQL migrations.

The current schema includes tables for:

* Users
* Sessions
* Posts
* Comments
* Post likes
* Comment likes

## Getting Started

### Prerequisites

Install:

* Go
* Node.js and npm
* PostgreSQL

### Backend

From the project root:

```bash
cd backend
go mod download
```

Configure PostgreSQL and the database connection required by the application, then start the backend:

```bash
go run ./cmd/api
```

### Frontend

In a separate terminal:

```bash
cd frontend
npm install
npm run dev
```

The Vite development server will provide the local frontend URL.

## Project Status

NeoMedia is an actively developed personal project.

The current version provides the core social media functionality, authentication, sessions, posts, comments, and likes.

## Planned Improvements

* User list
* Real-time chat
* Docker containerization
* CI/CD pipeline
* Cloud deployment
* Production configuration
* Automated testing
* Further application and infrastructure improvements
