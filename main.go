package main

import (
	"log"

	"github.com/BlazeCode1/books-api/controllers"
	"github.com/BlazeCode1/books-api/services"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	app := fiber.New()

	// Serve frontend files
	app.Static("/", "./client")

	// gRPC connection
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Initialize services and controllers
	bookService := services.NewBookService(conn)
	bookController := controllers.NewBookController(bookService)

	// Define routes
	app.Post("/book", bookController.AddBook)
	app.Get("/book", bookController.GetBooks)
	app.Delete("/book", bookController.DeleteBook)
	app.Put("/book", bookController.UpdateBook)

	log.Println("Server is running on http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
