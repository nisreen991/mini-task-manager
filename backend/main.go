package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"backend/handlers"
	"backend/middleware"

	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func connectDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	fmt.Println("Connected to MongoDB!")
}

func main() {
	connectDB()
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	mux := http.NewServeMux()

	mux.HandleFunc("/auth/signup", handlers.Signup(client))
	mux.HandleFunc("/auth/login", handlers.Login(client))

	mux.Handle("/tasks", middleware.AuthMiddleware(http.HandlerFunc(handlers.GetTasks(client))))
	mux.Handle("/tasks/create-task", middleware.AuthMiddleware(http.HandlerFunc(handlers.CreateTask(client))))
	mux.Handle("/tasks/", middleware.AuthMiddleware(http.HandlerFunc(handlers.TaskHandler(client))))

	// CORS setup
	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"}, // Allow your Angular frontend origin
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Authorization"},
		AllowCredentials: true,
		Debug:            true, // Enable for debugging CORS issues
	})

	h := corsOptions.Handler(mux)

	log.Println("Server starting on port 8080...")
	err := http.ListenAndServe(":8080", h)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
