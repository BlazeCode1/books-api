package services

import (
	"context"
	"github.com/BlazeCode1/books-api/app/book/client/book"
	"log"

	"github.com/segmentio/kafka-go"
	grpc "google.golang.org/grpc"
)

type BookService interface {
	AddBook(bookName string) (*book.BookResponse, error)
	GetBooks() ([]*book.BookListResponse, error)
	DeleteBook(id string) (*book.BookResponse, error)
	UpdateBook(id, bookName string) error
}

type bookService struct {
	client   book.BookServiceClient
	producer *kafka.Writer
}

func NewBookService(conn *grpc.ClientConn) BookService {
	client := book.NewBookServiceClient(conn)
	producer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
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

func (s *bookService) UpdateBook(id, bookName string) error {
	message := kafka.Message{
		Key: []byte(id),
		// todo: put an object not only a bookname
		Value: []byte(bookName),
	}
	if err := s.producer.WriteMessages(context.Background(), message); err != nil {
		log.Printf("Failed to send message to Kafka: %v", err)
		return err
	}
	log.Println("Message sent to Kafka")
	return nil
}
