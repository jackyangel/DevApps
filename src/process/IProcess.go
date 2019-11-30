package process

import (
	"github.com/devApps/src/model/request"
	"github.com/devApps/src/model/response"
)

//IProcess Interfacez para la capa de negocio
type IProcess interface {
	Process(req request.ShuntingRequest) (httpStatus int, res response.ShuntingRequest)
}
