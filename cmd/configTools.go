package cmd

import (
	"log"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

func yamlStringSettings() string {
	c := viper.AllSettings()
	bs, err := yaml.Marshal(c)
	if err != nil {
		log.Fatalf("unable to marshal config to YAML: %v", err)
	}
	return string(bs)
}

func UnmarshalConfigKey(key string, out interface{}) error {
	settings := viper.AllSettings()
	return mapstructure.Decode(settings[key], out)
}
