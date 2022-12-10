package controller

import (
	"encoding/json"
	"net/http"

	redisclient "github.com/cjlapao/common-go-redis"
	"github.com/cjlapao/common-go/helper/http_helper"
	"github.com/cjlapao/rediscli-go/entities"
	"github.com/gorilla/mux"
)

func SetListValues(w http.ResponseWriter, r *http.Request) {
	var request entities.RedisListSetRequest

	http_helper.MapRequestBody(r, &request)

	if request.Key == "" {
		w.WriteHeader(http.StatusBadRequest)
		response := entities.ApiErrorResponse{
			Error:       "invalid_key",
			Description: "key cannot be empty",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	redisCli := redisclient.New(connectionString)
	if redisCli == nil {
		w.WriteHeader(http.StatusBadRequest)
		response := entities.ApiErrorResponse{
			Error:       "invalid_client",
			Description: "Redis server did not responded",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	values := make([]interface{}, len(request.Values))
	for i, v := range request.Values {
		values[i] = v
	}

	err := redisCli.AddToList(request.Key, values...)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := entities.ApiErrorResponse{
			Error:       "invalid_request",
			Description: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(request)
}

func ListPopQueue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		response := entities.ApiErrorResponse{
			Error:       "invalid_key",
			Description: "Key name cannot be empty",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	if connectionString == "" {
		w.WriteHeader(http.StatusBadRequest)
		response := entities.ApiErrorResponse{
			Error:       "invalid_connection_string",
			Description: "Connection string cannot be empty",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	redisCli := redisclient.Get(connectionString)
	if redisCli == nil {
		w.WriteHeader(http.StatusBadRequest)
		response := entities.ApiErrorResponse{
			Error:       "invalid_client",
			Description: "Redis server did not responded",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	value, err := redisCli.PopQueueList(key)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := entities.ApiErrorResponse{
			Error:       "invalid_request",
			Description: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	response := entities.RedisGetKeyResponse{
		Key:   key,
		Value: value,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func ListPopStack(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		response := entities.ApiErrorResponse{
			Error:       "invalid_key",
			Description: "Key name cannot be empty",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	if connectionString == "" {
		w.WriteHeader(http.StatusBadRequest)
		response := entities.ApiErrorResponse{
			Error:       "invalid_connection_string",
			Description: "Connection string cannot be empty",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	redisCli := redisclient.Get(connectionString)
	if redisCli == nil {
		w.WriteHeader(http.StatusBadRequest)
		response := entities.ApiErrorResponse{
			Error:       "invalid_client",
			Description: "Redis server did not responded",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	value, err := redisCli.PopStackList(key)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := entities.ApiErrorResponse{
			Error:       "invalid_request",
			Description: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	response := entities.RedisGetKeyResponse{
		Key:   key,
		Value: value,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func CountListKeys(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		response := entities.ApiErrorResponse{
			Error:       "invalid_key",
			Description: "Key name cannot be empty",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	if connectionString == "" {
		w.WriteHeader(http.StatusBadRequest)
		response := entities.ApiErrorResponse{
			Error:       "invalid_connection_string",
			Description: "Connection string cannot be empty",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	redisCli := redisclient.Get(connectionString)
	if redisCli == nil {
		w.WriteHeader(http.StatusBadRequest)
		response := entities.ApiErrorResponse{
			Error:       "invalid_client",
			Description: "Redis server did not responded",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	value, err := redisCli.GetListCount(key)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := entities.ApiErrorResponse{
			Error:       "invalid_request",
			Description: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	response := entities.RedisListCountResponse{
		Key:   key,
		Count: value,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func TrimList(w http.ResponseWriter, r *http.Request) {
	var request entities.RedisListTrimRequest

	http_helper.MapRequestBody(r, &request)
	vars := mux.Vars(r)
	key := vars["key"]

	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		response := entities.ApiErrorResponse{
			Error:       "invalid_key",
			Description: "Key name cannot be empty",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	if connectionString == "" {
		w.WriteHeader(http.StatusBadRequest)
		response := entities.ApiErrorResponse{
			Error:       "invalid_connection_string",
			Description: "Connection string cannot be empty",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	redisCli := redisclient.Get(connectionString)
	if redisCli == nil {
		w.WriteHeader(http.StatusBadRequest)
		response := entities.ApiErrorResponse{
			Error:       "invalid_client",
			Description: "Redis server did not responded",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	err := redisCli.TrimList(key, request.From, request.To)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := entities.ApiErrorResponse{
			Error:       "invalid_request",
			Description: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	count, err := redisCli.GetListCount(key)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := entities.ApiErrorResponse{
			Error:       "invalid_request",
			Description: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	response := entities.RedisListCountResponse{
		Key:   key,
		Count: count,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
