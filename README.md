
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
│       └── main.go         
├── internal/
│   ├── handlers/          
│   │   ├── auth.go          
│   │   └── tasks.go         
│   ├── models/            
│   │   ├── task.go       
│   │   └── user.go       
│   ├── repository/        
│   │   ├── postgres.go    
│   │   └── interface.go   
│   ├── service/         
│   │   ├── auth.go         
│   │   └── tasks.go        
│   └── middleware/     
│       └── auth.go       
├── migrations/           
│   ├── 0001_init.up.sql    
│   └── 0001_init.down.sql  
├── .env.example      
├── go.mod 
└── README.md    
