package controllers

import (
	"github.com/BlazeCode1/books-api/services"
	"github.com/gofiber/fiber/v2"
)

type BookController struct {
	bookService services.BookService
}

func NewBookController(service services.BookService) *BookController {
	return &BookController{
		bookService: service,
	}
}

func (bc *BookController) AddBook(c *fiber.Ctx) error {
	var request struct {
		BookName string `json:"book_name"`
	}

	if err := c.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}

	response, err := bc.bookService.AddBook(request.BookName)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to process book name")
	}

	return c.JSON(response)
}

func (bc *BookController) GetBooks(c *fiber.Ctx) error {
	books, err := bc.bookService.GetBooks()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(books)
}

func (bc *BookController) DeleteBook(c *fiber.Ctx) error {
	var request struct {
		Id string `json:"id"`
	}

	if err := c.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}

	response, err := bc.bookService.DeleteBook(request.Id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete book")
	}

	return c.JSON(response)
}

func (bc *BookController) UpdateBook(c *fiber.Ctx) error {
	var request struct {
		Id       string `json:"id"`
		BookName string `json:"book_name"`
	}

	if err := c.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}

	err := bc.bookService.UpdateBook(request.Id, request.BookName)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to update book")
	}

	return c.JSON(fiber.Map{
		"message": "Book updated successfully",
	})
}
