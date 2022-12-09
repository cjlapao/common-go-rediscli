package controller

import (
	"encoding/json"
	"net/http"
)

func TestController(w http.ResponseWriter, r *http.Request) {
	users := struct {
		Message string
	}{
		Message: "All working",
	}

	json.NewEncoder(w).Encode(users)
}

// func TestContextController(w http.ResponseWriter, r *http.Request) {
// 	ctx := execution_context.Get()
// 	context := struct {
// 		ContextUser   authorization_context.ContextUser
// 		CorrelationId string
// 	}{
// 		ContextUser:   *ctx.Authorization.User,
// 		CorrelationId: ctx.Authorization.CorrelationId,
// 	}

// 	json.NewEncoder(w).Encode(context)
// }
