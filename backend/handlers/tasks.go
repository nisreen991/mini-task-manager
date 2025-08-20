package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetTasks retrieves all tasks for the authenticated user.
func GetTasks(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.Context().Value("username").(string)

		// Get user ID from username
		var user models.User
		userCollection := client.Database("taskmanager").Collection("users")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
		if err != nil {
			http.Error(w, "User not found", http.StatusInternalServerError)
			return
		}

		taskCollection := client.Database("taskmanager").Collection("tasks")
		cursor, err := taskCollection.Find(ctx, bson.M{"userId": user.ID})
		if err != nil {
			http.Error(w, "Failed to retrieve tasks", http.StatusInternalServerError)
			return
		}
		defer cursor.Close(ctx)

		var tasks []models.Task
		if err = cursor.All(ctx, &tasks); err != nil {
			http.Error(w, "Failed to decode tasks", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(tasks)
	}
}

// GetTaskById retrieves tasks for a given id for authenticated user
func GetTaskById(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.Context().Value("username").(string)

		var user models.User
		userCollection := client.Database("taskmanager").Collection("users")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
		if err != nil {
			http.Error(w, "User not found", http.StatusUnauthorized)
			return
		}
		idParam := r.URL.Path[len("/tasks/get-task/"):]
		id, err := primitive.ObjectIDFromHex(idParam)
		if err != nil {
			http.Error(w, "Invalid task ID", http.StatusBadRequest)
			return
		}
		// oid, err := primitive.ObjectIDFromHex(taskID)
		// if err != nil {
		// 	http.Error(w, "Invalid task ID", http.StatusBadRequest)
		// 	return
		// }
		taskCollection := client.Database("taskmanager").Collection("tasks")
		var task models.Task
		if err := taskCollection.FindOne(ctx, bson.M{"_id": id, "userId": user.ID}).Decode(&task); err != nil {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}
		// Return task as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)
	}
}

// CreateTask handles the creation of new tasks.
func CreateTask(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.Context().Value("username").(string)

		var user models.User
		userCollection := client.Database("taskmanager").Collection("users")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
		if err != nil {
			http.Error(w, "User not found", http.StatusInternalServerError)
			return
		}

		taskCollection := client.Database("taskmanager").Collection("tasks")

		var task models.Task
		err = json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		task.UserID = user.ID

		_, err = taskCollection.InsertOne(ctx, task)
		if err != nil {
			http.Error(w, "Failed to create task", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "Task created successfully"})
	}
}

// TaskHandler handles update, and delete operations for tasks.
func TaskHandler(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.Context().Value("username").(string)

		// Get user ID from username
		var user models.User
		userCollection := client.Database("taskmanager").Collection("users")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
		if err != nil {
			http.Error(w, "User not found", http.StatusInternalServerError)
			return
		}

		taskCollection := client.Database("taskmanager").Collection("tasks")

		switch r.Method {
		case "PUT":
			idParam := r.URL.Path[len("/tasks/"):]
			id, err := primitive.ObjectIDFromHex(idParam)
			if err != nil {
				http.Error(w, "Invalid task ID", http.StatusBadRequest)
				return
			}

			var updatedTask models.Task
			err = json.NewDecoder(r.Body).Decode(&updatedTask)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			filter := bson.M{"_id": id, "userId": user.ID}
			update := bson.M{"$set": bson.M{
				"title":       updatedTask.Title,
				"description": updatedTask.Description,
				"completed":   updatedTask.Completed,
			}}

			_, err = taskCollection.UpdateOne(ctx, filter, update)
			if err != nil {
				http.Error(w, "Failed to update task", http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"message": "Task updated successfully"})

		case "DELETE":
			idParam := r.URL.Path[len("/tasks/"):]
			id, err := primitive.ObjectIDFromHex(idParam)
			if err != nil {
				http.Error(w, "Invalid task ID", http.StatusBadRequest)
				return
			}

			filter := bson.M{"_id": id, "userId": user.ID}
			_, err = taskCollection.DeleteOne(ctx, filter)
			if err != nil {
				http.Error(w, "Failed to delete task", http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted successfully"})

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
	}
}
