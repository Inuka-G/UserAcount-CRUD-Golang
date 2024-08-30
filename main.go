package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	id      int    `json:"id"`
	body    string `json:"body"`
	isAdmin bool   `json:"isAdmin"`
}

func main() {
	todos := []Todo{}
	fmt.Println("data3")
	app := fiber.New()
	
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"data": "hello world"})
	})
	app.Post("/api/add", func(c *fiber.Ctx) error {

		singleTodo := &Todo{}
		if err := c.BodyParser(singleTodo); err != nil {
			return err
		}
		singleTodo.id = len(todos) + 1
		todos = append(todos, *singleTodo)
		fmt.Println(singleTodo)
		return c.Status(201).JSON(todos)
	})
	app.Listen(":4000")
}
