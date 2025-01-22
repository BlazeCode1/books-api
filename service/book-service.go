// book_service.go
package services

import (
	"context"

	"book-api/proto"
)

type BookService struct {
	Client proto.BookServiceClient
}

func (bs *BookService) AddBook(bookName string) (*proto.BookResponse, error) {
	return bs.Client.AddBook(context.Background(), &proto.BookRequest{BookName: bookName})
}

func (bs *BookService) GetBooks() ([]*proto.Book, error) {
	resp, err := bs.Client.GetBooks(context.Background(), &proto.EmptyRequest{})
	if err != nil {
		return nil, err
	}
	return resp.Books, nil
}

func (bs *BookService) DeleteBook(id string) (*proto.BookDeletionResponse, error) {
	return bs.Client.DeleteBook(context.Background(), &proto.BookDeletionRequest{Id: id})
}

func (bs *BookService) UpdateBook(id, bookName string) error {
	// Logic for Kafka message production, or call another service, if needed.
	// Example Kafka logic can be added here as required.
	return nil
}
