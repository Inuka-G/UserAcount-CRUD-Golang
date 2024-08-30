package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Account struct {
	Id       int    `json:"id"`
	UserName string `json:"userName"`
	IsAdmin  bool   `json:"isAdmin"`
}

func main() {
	allUserAccounts := []Account{}
	fmt.Println("data3")
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON("Get all user account")
	})

	app.Get("/accounts", func(c *fiber.Ctx) error {
		fmt.Println(allUserAccounts)
		return c.Status(200).JSON(allUserAccounts)
	})
	app.Post("/api/add", func(c *fiber.Ctx) error {

		singleAccount := new(Account)
		if err := c.BodyParser(singleAccount); err != nil {
			return err
		}
		singleAccount.Id = len(allUserAccounts) + 1

		// singleAccount.userName = "string(c.Body())"
		allUserAccounts = append(allUserAccounts, *singleAccount)
		fmt.Println(*singleAccount)
		return c.Status(201).JSON(*singleAccount)
	})
	app.Patch("/api/account/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for i, eachAccount := range allUserAccounts {
			if fmt.Sprint(eachAccount.Id) == id {
				allUserAccounts[i].IsAdmin = true

				return c.Status(200).JSON(allUserAccounts[i])
			}
		}
		return c.SendStatus(501)
	})
	app.Delete("/api/account/:id", func(c *fiber.Ctx) error {

		id := c.Params("id")
		for index, eachAccount := range allUserAccounts {
			if fmt.Sprint(eachAccount.Id) == id {
				allUserAccounts = append(allUserAccounts[:index], allUserAccounts[index+1:]...)

				return c.Status(200).JSON(allUserAccounts[index])
			}
		}
		return c.SendStatus(404)
	})
	app.Listen(":4000")
}
