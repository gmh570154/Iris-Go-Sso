package configs

import (
	"fmt"
	"io/ioutil"
	"iris_master/common/models"

	"gopkg.in/yaml.v2"
)

var AppConfig *models.Config

func InitConfig() {
	yamlFile, err := ioutil.ReadFile("./configs/config.yml")
	if err != nil {
		fmt.Println(err.Error())
	}

	err = yaml.Unmarshal(yamlFile, &AppConfig)
	if err != nil {
		fmt.Println(err.Error())
	}
}
