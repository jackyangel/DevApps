package process

import (
	"net/http"

	"github.com/devApps/src/model/request"
	"github.com/devApps/src/model/response"
	"github.com/devApps/src/shunting"
)

//ShuntingProcess estructura que implementa la interfaz IProcess
type ShuntingProcess struct {
	Shun shunting.IShuntingYard
}

//Process Metodo de la interaz IProcess
func (sbp ShuntingProcess) Process(req request.ShuntingRequest) (httpStatus int, res response.ShuntingRequest) {

	postfix := sbp.Shun.Evaluate(req.Exp)

	res.Infix = req.Exp
	res.Postfix = postfix
	res.Result = 0

	httpStatus = http.StatusOK
	return
}
