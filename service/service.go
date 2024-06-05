package service

import (
	"context"
	"fmt"

	"github.com/JonCanning/golangbookstore/database"
	"github.com/JonCanning/golangbookstore/types"
)

// a custom error type
type invalidRequestError struct {
	types.Request
}

// implement the error interface
func (err invalidRequestError) Error() string {
	return fmt.Sprintf("invalid request: %+v", err.Request)
}

// a higher order function that returns a function
func NewRequestHandler(db database.Database) types.RequestHandler {
	return func(ctx context.Context, request types.Request) types.Response {
		switch request := request.(type) {

		case types.CreateRequest:
			response := <-db.Create(ctx, request.Title, request.Author, request.Section)
			if response.Error != nil {
				return types.ErrorResponse{Error: response.Error}
			}
			return types.CreateResponse{Id: *response.Id}

		case types.ReadRequest:
			response := <-db.Read(ctx, request.Id)
			if response.Error != nil {
				return types.ErrorResponse{Error: response.Error}
			}
			return types.ReadResponse{Book: *response.Book}

		case types.UpdateRequest:
			response := <-db.Update(ctx, request.Id, request.Title, request.Author, request.Section)
			if response.Error != nil {
				return types.ErrorResponse{Error: response.Error}
			}
			return types.UpdateResponse{}

		case types.DeleteRequest:
			response := <-db.Delete(ctx, request.Id)
			if response.Error != nil {
				return types.ErrorResponse{Error: response.Error}
			}
			return types.DeleteResponse{}
		}

		return types.ErrorResponse{Error: invalidRequestError{Request: request}}
	}
}
