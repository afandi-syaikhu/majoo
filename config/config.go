package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/afandi-syaikhu/majoo/constant"
	"github.com/afandi-syaikhu/majoo/model"
)

func GetConfig() (*model.Config, error) {
	cfg := &model.Config{}

	jsonFile, err := ioutil.ReadFile(constant.ConfigFilepath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonFile, &cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
