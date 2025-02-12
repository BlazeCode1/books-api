package controllers

import (
	"github.com/BlazeCode1/books-api/app/book/model/Book"
	"github.com/BlazeCode1/books-api/app/book/services"
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
	var request Book.Book

	if err := c.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}

	book := Book.Book{
		BookName: request.BookName,
		Author:   request.Author,
	}
	response, err := bc.bookService.AddBook(book)
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
	var request Book.Book

	if err := c.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}

	book := request

	err := bc.bookService.UpdateBook(book.ID, book)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to update book")
	}

	return c.JSON(fiber.Map{
		"message": "Book updated successfully",
	})
}
