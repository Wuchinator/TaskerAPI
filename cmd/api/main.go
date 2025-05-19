package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Wuchinator/GoTasker/internal/handlers"
	"github.com/Wuchinator/GoTasker/internal/middleware"
	//"github.com/Wuchinator/GoTasker/internal/models"
	"github.com/Wuchinator/GoTasker/internal/repository"
	"github.com/Wuchinator/GoTasker/internal/service"
	//"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	//"golang.org/x/crypto/bcrypt"
	
)

func runMigrations(dbURL string) {
	m, err := migrate.New(
		"file://migrations",
		dbURL,
	)
	if err != nil {
		log.Fatal("Migration failed: ", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Migration up error: ", err)
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
	dbURL := os.Getenv("DB_URL")
	port := os.Getenv("PORT")
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET not set in .env")
	}
	runMigrations(dbURL)
	db, err := repository.NewPostgresDB(dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	taskRepo := repository.NewPostgresTaskRepository(db)
	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo, jwtSecret)
	taskService := service.NewTaskService(taskRepo)
	authHandler := handlers.NewAuthHandler(authService)
	taskHandler := handlers.NewTaskHandler(taskService)
	r := mux.NewRouter()
	r.HandleFunc("/register", authHandler.Register).Methods("POST")
	r.HandleFunc("/login", authHandler.Login).Methods("POST")
	secured := r.PathPrefix("/").Subrouter()
	secured.Use(middleware.AuthMiddleware(jwtSecret))
	secured.HandleFunc("/tasks", taskHandler.GetTasks).Methods("GET")
	secured.HandleFunc("/tasks/{id}", taskHandler.GetTask).Methods("GET")
	secured.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
	secured.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PUT")
	secured.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./static/index.html")
})


	server := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Server starting on :%s...\n", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Server failed: ", err)
	}
}