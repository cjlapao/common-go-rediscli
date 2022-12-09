package main

import (
	"github.com/cjlapao/common-go/execution_context"
	"github.com/cjlapao/common-go/helper"
	"github.com/cjlapao/common-go/service_provider"
	"github.com/cjlapao/common-go/version"
	"github.com/cjlapao/rediscli-go/startup"
)

var svc = service_provider.Get()

var ver = "0.0.0"
var services = execution_context.Get().Services

func main() {
	SetVersion()
	svc.Logger.WithTimestamp()

	configFile := helper.GetFlagValue("config", "")
	if configFile != "" {
		svc.Logger.Command("Loading configuration from " + configFile)
		svc.Configuration.LoadFromFile(configFile)
	}

	defer func() {
	}()

	startup.Init()
}

func SetVersion() {
	svc.Version.Name = "Redis Client"
	svc.Version.Author = "Carlos Lapao"
	svc.Version.License = "MIT"
	strVer, err := version.FromString(ver)
	if err == nil {
		services.Version.Major = strVer.Major
		services.Version.Minor = strVer.Minor
		services.Version.Build = strVer.Build
		services.Version.Rev = strVer.Rev
	}

	svc.Version.PrintAnsiHeader()
}
