package controller

import (
	"encoding/json"
	"net/http"

	"github.com/cjlapao/common-go/helper/http_helper"
	"github.com/cjlapao/rediscli-go/entities"
)

func SetConfig(w http.ResponseWriter, r *http.Request) {
	var request entities.Config

	http_helper.MapRequestBody(r, &request)

	if request.ConnectionString == "" {
		w.WriteHeader(http.StatusBadRequest)
		response := entities.ApiErrorResponse{
			Error:       "invalid_connection_string",
			Description: "Connection string cannot be empty",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	connectionString = request.ConnectionString
	logger.Info("New connection string was setup correctly")

	w.WriteHeader(http.StatusAccepted)
}
