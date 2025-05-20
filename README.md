
## GoTasker is a REST API for task management built with Go and PostgreSQL. It provides user authentication and CRUD operations for tasks with data isolation between users.

## API Endpoints
# Authentication
- POST /register - Register new user (email, password)

- POST /login - Login user (returns JWT token)

## Tasks (require JWT auth)
- GET /tasks - Get all user's tasks

- POST /tasks - Create new task (title, description optional)

- GET /tasks/{id} - Get task details

- PUT /tasks/{id} - Update task

- DELETE /tasks/{id} - Delete task

## Configuration (.env)
- DB_URL=your-db-url
- PORT=8080
- JWT_SECRET=your_random_secret_key
- JWT_EXPIRE_HOURS=24

 
## Project Structure
```
go-tasker/
├── cmd/
│   └── api/
│       └── main.go          # Entry point
├── internal/
│   ├── handlers/            # HTTP handlers
│   │   ├── auth.go          # Auth endpoints
│   │   └── tasks.go         # Task endpoints
│   ├── models/              # Data models
│   │   ├── task.go          # Task model  
│   │   └── user.go          # User model
│   ├── repository/          # Database layer
│   │   ├── postgres.go      # PostgreSQL impl
│   │   └── interface.go     # Repository interfaces
│   ├── service/             # Business logic  
│   │   ├── auth.go          # Auth service
│   │   └── tasks.go         # Task service
│   └── middleware/          # HTTP middleware
│       └── auth.go          # JWT auth  
├── migrations/              # Database migrations
│   ├── 0001_init.up.sql     # Initial schema
│   └── 0001_init.down.sql   # Rollback
├── .env.example             # Env template
├── go.mod                   # Go dependencies
└── README.md                # This file
