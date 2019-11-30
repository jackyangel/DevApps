package controllerfactory

import (
	"github.com/devApps/src/controller"
	"github.com/devApps/src/process"
	"github.com/devApps/src/shunting"
	"github.com/olebedev/config"
)

//BeneficiaryCFactory estrucutra que implementa la interfaz IControllerFactory
type ShuntingControllerFactory struct {
	Cfg *config.Config
}

//ReturnController metodo implentado de la interfaz IControllerFactory
func (shuc ShuntingControllerFactory) ReturnController() controller.IController {

	var shu shunting.IShuntingYard
	shu = shunting.ShuntingImp{}

	var pro process.IProcess
	pro = process.ShuntingProcess{Shun: shu}

	return controller.ShuntingController{Pro: pro}
}
