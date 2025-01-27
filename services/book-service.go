package services

import (
	"context"
	"log"

	pb "github.com/BlazeCode1/books-api/server/proto"
	"github.com/segmentio/kafka-go"
	grpc "google.golang.org/grpc"
)

type BookService interface {
	AddBook(bookName string) (*pb.BookResponse, error)
	GetBooks() ([]*pb.BookListResponse, error)
	DeleteBook(id string) (*pb.BookResponse, error)
	UpdateBook(id, bookName string) error
}

type bookService struct {
	client   pb.BookServiceClient
	producer *kafka.Writer
}

func NewBookService(conn *grpc.ClientConn) BookService {
	client := pb.NewBookServiceClient(conn)
	producer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "book-events",
	})
	return &bookService{client: client, producer: producer}
}

func (s *bookService) AddBook(bookName string) (*pb.BookResponse, error) {
	response, err := s.client.AddBook(context.Background(), &pb.BookRequest{BookName: bookName})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *bookService) GetBooks() ([]*pb.BookListResponse, error) {
	response, err := s.client.GetBooks(context.Background(), &pb.EmptyRequest{})
	if err != nil {
		return nil, err
	}
	return []*pb.BookListResponse{response}, nil
}

func (s *bookService) DeleteBook(id string) (*pb.BookResponse, error) {
	response, err := s.client.DeleteBook(context.Background(), &pb.BookDeletionRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *bookService) UpdateBook(id, bookName string) error {
	message := kafka.Message{
		Key:   []byte(id),
		Value: []byte(bookName),
	}
	if err := s.producer.WriteMessages(context.Background(), message); err != nil {
		log.Printf("Failed to send message to Kafka: %v", err)
		return err
	}
	log.Println("Message sent to Kafka")
	return nil
}
