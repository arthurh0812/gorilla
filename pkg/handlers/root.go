package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/arthurh0812/gorilla/pkg/api"
)

type API struct {
	logger log.Logger
}

func NewAPI(l log.Logger) API {
	return API{logger: l}
}

func (a API) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(&api.Response{
			Status: http.StatusOK,
			Message: "This is the homepage",
			Data: map[string]interface{}{
				"name": "Monchi",
				"age": 9,
			},
		})
	case http.MethodPost:
		a.logger.Printf("Route does not exist yet.\n")
	}
}