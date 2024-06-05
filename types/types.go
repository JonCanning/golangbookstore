//go:generate go run ../gen.go
package types

import (
	"context"
	"strings"
)

// primitives with semantic meaning
type Author string
type Id uint32
type Title string
type Section uint32

// poor mans enum
const (
	Fiction Section = iota
	NonFiction
)

// extend local types with receiver functions (methods)
func (author Author) ToUpper() string {
	return strings.ToUpper(string(author))
}

// struct with embedded fields, the type is the name of the field
type Book struct {
	Author
	Id
	Title
	Section
}

// constructor (can use promoted fields)
func NewBook(id Id, title Title, author Author, section Section) Book {
	return Book{
		Author:  author,
		Id:      id,
		Title:   title,
		Section: section,
	}
}

// marker interfaces
type Request interface {
	request()
}

type Response interface {
	response()
}

/*
the input port for the service

context.Context is a type that carries deadlines, cancelation signals, and other request-scoped values across API boundaries and between processes.
We won't use it in this example but it is a good practice to include it in the function signature since an adapter will likely have context to pass to the service.
*/
type RequestHandler func(context.Context, Request) Response

type CreateRequest struct {
	Author
	Title
	Section
}

// this is a receiver function that implements the Request interface
func (CreateRequest) request() {}

type CreateResponse struct {
	Id
}

// this is a receiver function that implements the Response interface
func (CreateResponse) response() {}

type ReadRequest struct {
	Id
}

func (ReadRequest) request() {}

type ReadResponse struct {
	Book
}

func (ReadResponse) response() {}

type UpdateRequest struct {
	Author
	Id
	Title
	Section
}

func (UpdateRequest) request() {}

type UpdateResponse struct{}

func (UpdateResponse) response() {}

type DeleteRequest struct {
	Id
}

func (DeleteRequest) request() {}

type DeleteResponse struct{}

func (DeleteResponse) response() {}

type ErrorResponse struct {
	Error error
}

func (ErrorResponse) response() {}
