package main

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/omdxp/todo-api/data"
	_ "github.com/omdxp/todo-api/docs"
)

type Response struct {
	Message string `json:"message"`
}

// @title Todo API
// @description This is a sample server Todo API server.
// @version 1.0.0
// @contact.name Omdxp
// @host localhost:8080
// @BasePath /
func main() {
	app := fiber.New()

	// CORS
	app.Use(cors.New(cors.ConfigDefault))

	// Swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Routes
	app.Get("/todos", GetTodos)
	app.Get("/todos/:id", GetTodo)
	app.Post("/todos", CreateTodo)
	app.Put("/todos/:id", UpdateTodo)
	app.Delete("/todos/:id", DeleteTodo)

	app.Listen(":8080")
}

// GetTodos returns all todos
// @Summary Get all todos
// @Description Get all todos
// @ID get-todos
// @Produce  json
// @Success 200 {array} data.Todo
// @Failure 400 {object} Response
// @Failure 404 {object} Response
// @Failure 500 {object} Response
// @Router /todos [get]
func GetTodos(c *fiber.Ctx) error {
	todos, status, err := data.GetTodos()
	if err != nil {
		return c.Status(status).JSON(Response{Message: err.Error()})
	}
	return c.Status(status).JSON(todos)
}

// GetTodo returns a todo by id
// @Summary Get a todo by id
// @Description Get a todo by id
// @ID get-todo
// @Produce  json
// @Param id path string true "Todo ID"
// @Success 200 {object} data.Todo
// @Failure 400 {object} Response
// @Failure 404 {object} Response
// @Failure 500 {object} Response
// @Router /todos/{id} [get]
func GetTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	todo, status, err := data.GetTodo(id)
	if err != nil {
		return c.Status(status).JSON(Response{Message: err.Error()})
	}
	return c.Status(status).JSON(todo)
}

// CreateTodo creates a new todo
// @Summary Create a new todo
// @Description Create a new todo
// @ID create-todo
// @Accept  json
// @Produce  json
// @Param request-body body data.CreateTodoRequest true "CreateTodoRequest"
// @Success 200 {object} data.Todo
// @Failure 400 {object} Response
// @Failure 404 {object} Response
// @Failure 500 {object} Response
// @Router /todos [post]
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
// @Summary Update a todo by id
// @Description Update a todo by id
// @ID update-todo
// @Accept  json
// @Produce  json
// @Param id path string true "Todo ID"
// @Param request-body body data.UpdateTodoRequest true "UpdateTodoRequest"
// @Success 200 {object} data.Todo
// @Failure 400 {object} Response
// @Failure 404 {object} Response
// @Failure 500 {object} Response
// @Router /todos/{id} [put]
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
// @Summary Delete a todo by id
// @Description Delete a todo by id
// @ID delete-todo
// @Produce  json
// @Param id path string true "Todo ID"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 404 {object} Response
// @Failure 500 {object} Response
// @Router /todos/{id} [delete]
func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	status, err := data.DeleteTodo(id)
	if err != nil {
		return c.Status(status).JSON(Response{Message: err.Error()})
	}
	return c.Status(status).JSON(Response{Message: "Todo deleted"})
}
