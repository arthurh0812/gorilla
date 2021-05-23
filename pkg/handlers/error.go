package handlers

import (
	"log"
	"net/http"

	"github.com/arthurh0812/gorilla/pkg/view"
)

type Error struct {
	logger *log.Logger
	view *view.View

	status int
}

var errorHandler *Error

func NewError(l *log.Logger, status int) *Error {
	if errorHandler == nil {
		errorHandler = &Error{logger: l, status: status, view: view.NewView()}
	}
	return errorHandler
}

func (e *Error) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var err error
	switch e.status {
	case http.StatusInternalServerError:
		w.WriteHeader(http.StatusInternalServerError)
		err = e.view.Render(w, "404.html", view.Data{
			Title: "Internal Server Error",
			Header: "Oops!",
			Message: "An internal server error occurred...",
		})
	default:
		w.WriteHeader(http.StatusNotFound)
		err = e.view.NotFound(w, req)
	}
	if err != nil {
		e.logger.Println(err)
		http.Error(w, "something went wrong...", http.StatusInternalServerError)
	}
}
