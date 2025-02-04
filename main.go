package main

import (
	"log"

	"github.com/BlazeCode1/books-api/app/book/controllers"
	"github.com/BlazeCode1/books-api/app/book/services"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	app := fiber.New()
	// Add CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173", //  React app URL
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))
	// Serve frontend files
	//app.Static("/", "./client")

	// gRPC connection
	conn, err := grpc.NewClient("host.docker.internal:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC client: %v", err)
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
