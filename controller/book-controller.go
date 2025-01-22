// book_controller.go
package controllers

import (


	"github.com/gofiber/fiber/v2"
	"github.com/BlazeCode1/books-api/service"
)

type BookController struct {
	BookHandler services.BookHandler
}

func (bc *BookController) CreateBook(c *fiber.Ctx) error {
	var data struct {
		BookName string `json:"book_name"`
	}
	if err := c.BodyParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}

	resp, err := bc.BookHandler.AddBook(data.BookName)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to process book")
	}

	return c.JSON(fiber.Map{"message": resp.Message})
}

func (bc *BookController) ListBooks(c *fiber.Ctx) error {
	books, err := bc.BookHandler.GetBooks()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch books")
	}

	return c.JSON(books)
}
