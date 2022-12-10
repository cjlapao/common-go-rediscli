package controller

import (
	"encoding/json"
	"net/http"

	redisclient "github.com/cjlapao/common-go-redis"
	restapi "github.com/cjlapao/common-go-restapi"
	"github.com/cjlapao/common-go/execution_context"
	"github.com/cjlapao/common-go/helper/http_helper"
	"github.com/cjlapao/common-go/log"
	"github.com/cjlapao/rediscli-go/constants"
	"github.com/cjlapao/rediscli-go/entities"
	"github.com/gorilla/mux"
)

var listener *restapi.HttpListener
var logger = log.Get()
var connectionString string
var ctx = execution_context.Get()

func Init() {
	listener = restapi.GetHttpListener()
	listener.AddJsonContent().AddLogger().AddHealthCheck()
	connectionString = ctx.Configuration.GetString(constants.REDIS_CONNECTION_STRING_ENVIRONMENT_VAR)

	listener.AddController(SetConfig, "/config", "POST")

	// Generic
	listener.AddController(DeleteKey, "/key/{key}", "DELETE")
	listener.AddController(GetKeys, "/keys", "GET")

	// Strings
	listener.AddController(SetStringKey, "/key/string", "POST")
	listener.AddController(GetStringKey, "/key/string/{key}", "GET")

	// Lists
	listener.AddController(SetListValues, "/key/list", "POST")
	listener.AddController(DeleteKey, "/key/list/{key}", "DELETE")
	listener.AddController(CountListKeys, "/key/list/{key}/count", "GET")
	listener.AddController(ListPopQueue, "/key/list/{key}/pop-queue", "GET")
	listener.AddController(ListPopStack, "/key/list/{key}/pop-stack", "GET")
	listener.AddController(TrimList, "/key/list/{key}/trim", "POST")

	listener.Start()
}

func DeleteKey(w http.ResponseWriter, r *http.Request) {
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

	_, err := redisCli.Delete(key)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := entities.ApiErrorResponse{
			Error:       "invalid_request",
			Description: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetKeys(w http.ResponseWriter, r *http.Request) {
	var body entities.RedisGetKeysRequest

	http_helper.MapRequestBody(r, &body)

	if body.Pattern == "" {
		w.WriteHeader(http.StatusBadRequest)
		response := entities.ApiErrorResponse{
			Error:       "invalid_pattern",
			Description: "Pattern cannot be empty, use * to get all",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	if body.Pattern == "*" {
		body.Pattern = ""
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

	keys, err := redisCli.GetAllKeys(body.Pattern)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := entities.ApiErrorResponse{
			Error:       "invalid_request",
			Description: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	response := entities.RedisGetKeysResponse{
		Keys: keys,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
