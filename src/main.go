package main

import (
	"github.com/devApps/src/configurations"
	"github.com/devApps/src/controller"
	"github.com/devApps/src/factories/controllerfactory"
	"github.com/devApps/src/handler"
	"github.com/olebedev/config"
)

var (
	port              string
	path              string
	cfg               *config.Config
	control           controller.IController
	factoryController controllerfactory.IControllerFactory
)

func init() {
	cfg, _ = configurations.GetConfig()
	port, _ = cfg.String("server.port")
	path, _ = cfg.String("server.path")
	factoryController = controllerfactory.ShuntingControllerFactory{Cfg: cfg}
	control = factoryController.ReturnController()
}

func main() {
	var h handler.IHandlerFunction

	h = handler.ShuntingHandler{
		Port:    port,
		Path:    path,
		Control: control}

	h.Run()
}
