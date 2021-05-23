package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/arthurh0812/gorilla/pkg/api"
	"github.com/arthurh0812/gorilla/pkg/view"
)

type API struct {
	logger *log.Logger
	view *view.View
}

var apiHandler *API

func NewAPI(l *log.Logger) *API {
	if apiHandler == nil {
		apiHandler = &API{logger: l, view: view.NewView()}
	}
	return apiHandler
}

func (a API) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var err error
	switch req.Method {
	case http.MethodGet:
		err = json.NewEncoder(w).Encode(&api.Response{
			Status: http.StatusOK,
			Message: "This is the homepage",
			Data: map[string]interface{}{
				"name": "Monchi",
				"age": 9,
			},
		})
	default:
		w.WriteHeader(http.StatusNotFound)
		err = a.view.NotFound(w, req)
	}
	if err != nil {
		a.logger.Println(err)
		http.Error(w, "something went wrong...", http.StatusInternalServerError)
	}
}