package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/devApps/src/model/request"
	"github.com/devApps/src/process"
)

type ShuntingController struct {
	Pro process.IProcess
}

func (shu ShuntingController) Execute(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var request request.ShuntingRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	httpStatus, res := shu.Pro.Process(request)
	w.WriteHeader(httpStatus)

	json.NewEncoder(w).Encode(res)
}
