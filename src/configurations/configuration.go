package configurations

import (
	"io/ioutil"
	"path/filepath"

	"github.com/olebedev/config"
)

func GetConfig() (cfg *config.Config, err error) {

	//Validamos las configuraciones de la aplicación

	absPath, _ := filepath.Abs("src/github.com/myService/src/resources/appConfig.yml")
	fileconfig, err := ioutil.ReadFile(absPath)
	if err != nil {
		panic(err)
	}
	yamlString := string(fileconfig)
	if err != nil {
		panic(err)
	}
	cfg, err = config.ParseYaml(yamlString)
	if err != nil {
		panic(err)
	}

	return
}
