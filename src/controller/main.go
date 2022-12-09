package controller

import (
	restapi "github.com/cjlapao/common-go-restapi"
)

var listener *restapi.HttpListener

func Init() {
	// userCtx := UserContext{}
	listener = restapi.GetHttpListener()
	listener.AddJsonContent().AddLogger().AddHealthCheck()
	// listener.WithAuthentication(userCtx)

	listener.AddController(TestController, "/test", "GET")

	// listener.AddAuthorizedController(TestContextController, "/test/context", "GET")
	listener.Start()
}
