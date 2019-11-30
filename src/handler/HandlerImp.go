package handler

import (
	"log"
	"net/http"

	c "github.com/devApps/src/controller"

	"github.com/gorilla/mux"
)

type ShuntingHandler struct {
	Port    string
	Path    string
	Control c.IController
}

func (shuH ShuntingHandler) Run() {
	// Iniciando ruteo
	r := mux.NewRouter()

	r.HandleFunc(shuH.Path, shuH.Control.Execute).Methods("POST")

	log.Println("Run services Port: " + shuH.Port)

	log.Fatal(http.ListenAndServe(shuH.Port, r))

}
