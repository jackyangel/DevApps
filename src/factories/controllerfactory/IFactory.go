package controllerfactory

import (
	"github.com/devApps/src/controller"
)

//IControllerFactory interfaz que retorna controllers
type IControllerFactory interface {
	ReturnController() controller.IController
}
