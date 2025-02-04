package services

import (
	"context"
	"encoding/json"
	"log"

	"github.com/BlazeCode1/books-api/app/book/client/book"

	"github.com/BlazeCode1/books-api/app/book/model/Book"
	"github.com/segmentio/kafka-go"
	grpc "google.golang.org/grpc"
)

type BookService interface {
	AddBook(bookName string) (*book.BookResponse, error)
	GetBooks() ([]*book.BookListResponse, error)
	DeleteBook(id string) (*book.BookResponse, error)
	UpdateBook(id string, bookObject Book.Book) error
}

type bookService struct {
	client   book.BookServiceClient
	producer *kafka.Writer
}

func NewBookService(conn *grpc.ClientConn) BookService {
	client := book.NewBookServiceClient(conn)
	producer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"broker:29092"},
		Topic:   "book-events",
	})
	return &bookService{client: client, producer: producer}
}

func (s *bookService) AddBook(bookName string) (*book.BookResponse, error) {
	response, err := s.client.AddBook(context.Background(), &book.BookRequest{BookName: bookName})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *bookService) GetBooks() ([]*book.BookListResponse, error) {
	response, err := s.client.GetBooks(context.Background(), &book.EmptyRequest{})
	if err != nil {
		return nil, err
	}
	return []*book.BookListResponse{response}, nil
}

func (s *bookService) DeleteBook(id string) (*book.BookResponse, error) {
	response, err := s.client.DeleteBook(context.Background(), &book.BookDeletionRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *bookService) UpdateBook(id string, book Book.Book) error {
	// Serialize the entire book struct into JSON
	bookData, err := json.Marshal(book)
	if err != nil {
		log.Printf("Failed to serialize book: %v", err)
		return err
	}

	// Create Kafka message with ID as the key and the serialized book as the value
	message := kafka.Message{
		Key:   []byte(id),
		Value: bookData,
	}
	// Send the message to Kafka
	if err := s.producer.WriteMessages(context.Background(), message); err != nil {
		log.Printf("Failed to send message to Kafka: %v", err)
		return err
	}

	log.Println("Message sent to Kafka")
	return nil
}
