package config

import "github.com/jinzhu/configor"

const ymlFilePath = "config/app.yml"

func NewYaml() (app Application, err error) {
	err = configor.Load(&app, ymlFilePath)
	return
}
