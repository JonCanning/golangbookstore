package database

import (
	"context"
	"fmt"

	"github.com/JonCanning/golangbookstore/types"
)

type NotFoundError struct {
	types.Id
}

func (err NotFoundError) Error() string {
	return fmt.Sprintf("NotFound: %d", err.Id)
}

type CreateResponse struct {
	*types.Id //pointer to the Id because it might be nil
	Error     error
}

type ReadResponse struct {
	*types.Book       //pointer to the Book because it might be nil
	Error       error //could be NotFoundError or something more catastrophic
}

type ErrorResponse struct {
	Error error
}

/*
a receive only chan uses the arrow before the type (<-chan)
a send only chan uses the arrow after the type (chan<-)
a bidirectional chan does not have an arrow (chan)
*/
type Database struct {
	Create func(context.Context, types.Title, types.Author, types.Section) <-chan CreateResponse
	Read   func(context.Context, types.Id) <-chan ReadResponse
	Update func(context.Context, types.Id, types.Title, types.Author, types.Section) <-chan ErrorResponse
	Delete func(context.Context, types.Id) <-chan ErrorResponse
}
