package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arthurh0812/gorilla/pkg/view"
)

type Login struct {
	logger *log.Logger
	view *view.View
}

func NewLogin(l *log.Logger) *Login {
	return &Login{
		logger: l,
		view: view.NewView(),
	}
}

func (l *Login) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var err error
	switch req.Method {
	case http.MethodGet:
		err = l.view.Render(w, "login.html", view.Data{
			Title: "Test",
			Header: "Please fill out the fields.",
		})
	case http.MethodPost:
		fName := req.PostFormValue("firstName")
		lName := req.PostFormValue("lastName")

		l.logger.Printf("fName=%s lName=%s", fName, lName)

		err = l.view.Render(w, "success.html", view.Data{
			Title: "Success!",
			Header: "Yaaay!",
			Message: fmt.Sprintf("You are now successfully logged in as %s %s", fName, lName),
			Data: map[string]interface{}{
				"firstName": fName,
				"lastName": lName,
			},
		})
	default:
		w.WriteHeader(http.StatusNotFound)
		err = l.view.NotFound(w, req)
	}
	if err != nil {
		http.Error(w, "something went wrong...", http.StatusInternalServerError)
	}
}