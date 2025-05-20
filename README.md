
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
📁 go-tasker-api
├── 📁 cmd
│   └── 📁 api
│       └── main.go             
├── 📁 docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── 📁 internal                 
│   ├── 📁 config               
│   │   └── config.go
│   ├── 📁 errors               
│   │   └── errors.go
│   ├── 📁 handlers             
│   │   ├── auth.go
│   │   └── tasks.go
│   ├── 📁 logger             
│   │   └── logger.go
│   ├── 📁 middleware           
│   │   ├── auth.go
│   │   └── middleware.go
│   ├── 📁 models              
│   │   ├── task.go
│   │   └── user.go
│   ├── 📁 repository           
│   │   ├── db.go
│   │   ├── interface.go      
│   │   ├── memory.go         
│   │   ├── postgres.go       
│   │   └── user.go
│   └── 📁 service             
│       ├── auth.go
│       └── tasks.go
├── 📁 migrations
├── 📁 static
