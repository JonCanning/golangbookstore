package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/JonCanning/golangbookstore/database"
	"github.com/JonCanning/golangbookstore/types"
	"github.com/stretchr/testify/assert"
)

// there's too much behavior in the test database, but it allows us to demonstrate variadic arguments
func newTestDatabase(books ...types.Book) database.Database {
	data := make(map[types.Id]types.Book)
	for _, book := range books {
		data[book.Id] = book
	}

	create := func(context.Context, types.Title, types.Author, types.Section) <-chan database.CreateResponse {
		id := types.Id(len(data) + 1)
		response := make(chan database.CreateResponse)
		go func() {
			response <- database.CreateResponse{Id: &id}
			close(response)
		}()
		return response
	}

	read := func(ctx context.Context, id types.Id) <-chan database.ReadResponse {
		response := make(chan database.ReadResponse)
		go func() {
			book, ok := data[id]
			if !ok {
				response <- database.ReadResponse{Error: database.NotFoundError{Id: id}}
			} else {
				response <- database.ReadResponse{Book: &book}
			}
			close(response)
		}()
		return response
	}

	update := func(context.Context, types.Id, types.Title, types.Author, types.Section) <-chan database.ErrorResponse {
		response := make(chan database.ErrorResponse)
		go func() {
			response <- database.ErrorResponse{}
			close(response)
		}()
		return response
	}

	delete := func(context.Context, types.Id) <-chan database.ErrorResponse {
		response := make(chan database.ErrorResponse)
		go func() {
			response <- database.ErrorResponse{}
			close(response)
		}()
		return response
	}

	return database.Database{
		Create: create,
		Read:   read,
		Update: update,
		Delete: delete,
	}
}

func NewTestErrorDatabase() database.Database {
	error := fmt.Errorf("error")
	create := func(context.Context, types.Title, types.Author, types.Section) <-chan database.CreateResponse {
		response := make(chan database.CreateResponse)
		go func() { response <- database.CreateResponse{Error: error} }()
		return response
	}
	read := func(context.Context, types.Id) <-chan database.ReadResponse {
		response := make(chan database.ReadResponse)
		go func() { response <- database.ReadResponse{Error: error} }()
		return response
	}
	update := func(context.Context, types.Id, types.Title, types.Author, types.Section) <-chan database.ErrorResponse {
		response := make(chan database.ErrorResponse)
		go func() { response <- database.ErrorResponse{Error: error} }()
		return response
	}
	delete := func(context.Context, types.Id) <-chan database.ErrorResponse {
		response := make(chan database.ErrorResponse)
		go func() { response <- database.ErrorResponse{Error: error} }()
		return response
	}
	return database.Database{
		Create: create,
		Read:   read,
		Update: update,
		Delete: delete,
	}
}

func Test_create(t *testing.T) {
	db := newTestDatabase()
	requestHandler := NewRequestHandler(db)
	request := types.CreateRequest{
		Title:   "The Hobbit",
		Author:  "J.R.R. Tolkien",
		Section: types.Fiction,
	}
	ctx := context.Background()
	response := requestHandler(ctx, request)
	expected := types.CreateResponse{
		Id: 1}
	assert.Equal(t, expected, response)
}

func Test_create_error(t *testing.T) {
	database := NewTestErrorDatabase()
	requestHandler := NewRequestHandler(database)
	request := types.CreateRequest{
		Title:   "The Hobbit",
		Author:  "J.R.R. Tolkien",
		Section: types.Fiction,
	}
	ctx := context.Background()
	response := requestHandler(ctx, request)
	expected := types.ErrorResponse{
		Error: fmt.Errorf("error")}
	assert.Equal(t, expected, response)
}

func Test_read(t *testing.T) {
	book := types.NewBook(1, "The Hobbit", "J.R.R. Tolkien", types.Fiction)
	db := newTestDatabase(book)
	requestHandler := NewRequestHandler(db)
	readRequest := types.ReadRequest{
		Id: 1,
	}
	ctx := context.Background()
	response := requestHandler(ctx, readRequest)
	expected := types.ReadResponse{
		Book: book}
	assert.Equal(t, expected, response)
}

func Test_read_error(t *testing.T) {
	db := NewTestErrorDatabase()
	requestHandler := NewRequestHandler(db)
	request := types.ReadRequest{Id: 1}
	ctx := context.Background()
	response := requestHandler(ctx, request)
	expected := types.ErrorResponse{
		Error: fmt.Errorf("error")}
	assert.Equal(t, expected, response)
}

func Test_read_not_found(t *testing.T) {
	db := newTestDatabase()
	requestHandler := NewRequestHandler(db)
	request := types.ReadRequest{Id: 1}
	ctx := context.Background()
	response := requestHandler(ctx, request)
	expected := types.ErrorResponse{
		Error: database.NotFoundError{Id: 1}}
	assert.Equal(t, expected, response)
}
