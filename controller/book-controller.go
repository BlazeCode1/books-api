// book_controller.go
package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/blazecode1/books-api/services"
)

type BookController struct {
	BookService services.BookService
}

func (bc *BookController) CreateBook(c *fiber.Ctx) error {
	var data struct {
		BookName string `json:"book_name"`
	}
	if err := c.BodyParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}

	resp, err := bc.BookService.AddBook(data.BookName)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to process book")
	}

	return c.JSON(fiber.Map{"message": resp.Message})
}

func (bc *BookController) ListBooks(c *fiber.Ctx) error {
	books, err := bc.BookService.GetBooks()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch books")
	}

	return c.JSON(books)
}

func (bc *BookController) DeleteBook(c *fiber.Ctx) error {
	var data struct {
		Id string `json:"id"`
	}
	if err := c.BodyParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}

	resp, err := bc.BookService.DeleteBook(data.Id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete book")
	}

	return c.JSON(fiber.Map{"message": resp.Message})
}

func (bc *BookController) UpdateBook(c *fiber.Ctx) error {
	var data struct {
		Id       string `json:"id"`
		BookName string `json:"book_name"`
	}
	if err := c.BodyParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}

	err := bc.BookService.UpdateBook(data.Id, data.BookName)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to update book")
	}

	return c.JSON(fiber.Map{"message": "Book updated successfully"})
}
