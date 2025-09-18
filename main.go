package main

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Message string `json:"message"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "John Doe"},
	{ID: 2, Name: "Jane Smith"},
}

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(Response{Message: "Welcome to the User API"})
	})

	app.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(users)
	})

	app.Get("users/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(Response{Message: "Invalid ID"})
		}

		for _, user := range users {
			if user.ID == id {
				return c.JSON(user)
			}
		}

		return c.Status(fiber.StatusNotFound).JSON(Response{Message: "User not found"})
	})

	app.Post("/users", func(c *fiber.Ctx) error {
		var newUser User
		if err := c.BodyParser(&newUser); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(Response{Message: "Invalid request body"})
		}

		newUser.ID = len(users) + 1
		users = append(users, newUser)
		return c.Status(fiber.StatusCreated).JSON(newUser)
	})

	app.Delete("/users/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(Response{Message: "Invalid ID"})
		}

		for i, user := range users {
			if user.ID == id {
				users = append(users[:i], users[i+1:]...)
				return c.JSON(Response{Message: "User deleted"})
			}
		}

		return c.Status(fiber.StatusNotFound).JSON(Response{Message: "User not found"})
	})

	app.Patch("/users/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(Response{Message: "Invalid ID"})
		}

		var updatedData User
		if err := c.BodyParser(&updatedData); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(Response{Message: "Invalid request body"})
		}

		for i, user := range users {
			if user.ID == id {
				if updatedData.Name != "" {
					users[i].Name = updatedData.Name
				}
				return c.JSON(users[i])
			}
		}

		return c.Status(fiber.StatusNotFound).JSON(Response{Message: "User not found"})
	})

	app.Put("/users/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(Response{Message: "Invalid ID"})
		}

		var updatedUser User
		if err := c.BodyParser(&updatedUser); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(Response{Message: "Invalid request body"})
		}

		for i, user := range users {
			if user.ID == id {
				users[i].Name = updatedUser.Name
				return c.JSON(users[i])
			}
		}

		return c.Status(fiber.StatusNotFound).JSON(Response{Message: "User not found"})
	})

	log.Fatal(app.Listen(":3000")) //app.Listen(":3000");

}
