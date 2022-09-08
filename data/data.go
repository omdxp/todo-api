package data

import (
	"context"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ctx        context.Context
	collection *mongo.Collection
)

func init() {
	ctx = context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %s", err.Error())
	}
	collection = client.Database("db").Collection("todos")
}

// Todo is a struct that represents a todo item
type Todo struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title     string             `bson:"title" json:"title"`
	Completed bool               `bson:"completed" json:"completed"`
	CreatedAt string             `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt string             `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

// Todos is a slice of Todo
type Todos []Todo

// GetTodos returns all todos
func GetTodos() (Todos, int, error) {
	var todos Todos
	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return Todos{}, http.StatusInternalServerError, err
	}
	defer cur.Close(ctx)
	if err := cur.All(ctx, &todos); err != nil {
		return Todos{}, http.StatusInternalServerError, err
	}
	if todos == nil {
		return Todos{}, http.StatusNotFound, nil
	}
	return todos, http.StatusOK, nil
}

// GetTodo returns a todo by id
func GetTodo(id string) (Todo, int, error) {
	var todo Todo
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Todo{}, http.StatusBadRequest, err
	}
	if err := collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&todo); err != nil {
		return Todo{}, http.StatusNotFound, err
	}
	return todo, http.StatusOK, nil
}

// CreateTodo creates a new todo
func CreateTodo(todo Todo) (Todo, int, error) {
	result, err := collection.InsertOne(ctx, todo)
	if err != nil {
		return Todo{}, http.StatusInternalServerError, err
	}
	todo.ID = result.InsertedID.(primitive.ObjectID)
	todo.CreatedAt = time.Now().UTC().Format(time.RFC3339)
	todo.UpdatedAt = todo.CreatedAt
	return todo, http.StatusCreated, nil
}

// UpdateTodo updates a todo by id
func UpdateTodo(id string, todo Todo) (Todo, int, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Todo{}, http.StatusBadRequest, err
	}
	// check if todo exists
	if err := collection.FindOne(ctx, bson.M{"_id": objID}).Err(); err != nil {
		return Todo{}, http.StatusNotFound, err
	}
	if _, err := collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": todo}); err != nil {
		return Todo{}, http.StatusInternalServerError, err
	}
	todo.ID = objID
	todo.UpdatedAt = time.Now().UTC().Format(time.RFC3339)
	return todo, http.StatusOK, nil
}

// DeleteTodo deletes a todo by id
func DeleteTodo(id string) (int, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return http.StatusBadRequest, err
	}
	// check if todo exists
	if err := collection.FindOne(ctx, bson.M{"_id": objID}).Err(); err != nil {
		return http.StatusNotFound, err
	}
	if _, err := collection.DeleteOne(ctx, bson.M{"_id": objID}); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
