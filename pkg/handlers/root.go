package handlers

import (
	"github.com/arthurh0812/gorilla/pkg/view"
	"log"
	"net/http"
)

type Root struct {
	logger *log.Logger
	view *view.View
}

func NewRoot(l *log.Logger) *Root {
	return &Root{logger: l, view: view.NewView()}
}

func (r Root) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var err error
	switch req.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
		err = r.view.Static(w, "base.html")
	default:
		w.WriteHeader(http.StatusNotFound)
		err = r.view.NotFound(w, req)
	}
	if err != nil {
		http.Error(w, "something went wrong...", http.StatusInternalServerError)
	}
}