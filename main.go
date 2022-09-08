package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/omdxp/todo-api/data"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	app := fiber.New()

	// CORS
	app.Use(cors.New(cors.ConfigDefault))

	app.Get("/todos", GetTodos)
	app.Get("/todos/:id", GetTodo)
	app.Post("/todos", CreateTodo)
	app.Put("/todos/:id", UpdateTodo)
	app.Delete("/todos/:id", DeleteTodo)

	app.Listen(":8080")
}

// GetTodos returns all todos
func GetTodos(c *fiber.Ctx) error {
	todos, status, err := data.GetTodos()
	if err != nil {
		return c.Status(status).JSON(Response{Message: err.Error()})
	}
	return c.Status(status).JSON(todos)
}

// GetTodo returns a todo by id
func GetTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	todo, status, err := data.GetTodo(id)
	if err != nil {
		return c.Status(status).JSON(Response{Message: err.Error()})
	}
	return c.Status(status).JSON(todo)
}

// CreateTodo creates a new todo
func CreateTodo(c *fiber.Ctx) error {
	var todo data.Todo
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Message: err.Error()})
	}
	todo, status, err := data.CreateTodo(todo)
	if err != nil {
		return c.Status(status).JSON(Response{Message: err.Error()})
	}
	return c.Status(status).JSON(todo)
}

// UpdateTodo updates a todo by id
func UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo data.Todo
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Message: err.Error()})
	}
	todo, status, err := data.UpdateTodo(id, todo)
	if err != nil {
		return c.Status(status).JSON(Response{Message: err.Error()})
	}
	return c.Status(status).JSON(todo)
}

// DeleteTodo deletes a todo by id
func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	status, err := data.DeleteTodo(id)
	if err != nil {
		return c.Status(status).JSON(Response{Message: err.Error()})
	}
	return c.Status(status).JSON(Response{Message: "Todo deleted"})
}
