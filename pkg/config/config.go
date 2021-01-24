package config

import (
	"baoctl/pkg/types"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var configs types.Config

func Initialize(options *types.Options) error {
	data, err := ioutil.ReadFile(options.FilePath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &configs)
	if err != nil {
		return err
	}
	return nil
}

func Config() *types.Config {
	return &configs
}
