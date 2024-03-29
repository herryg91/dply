package serviceYaml

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type ServiceYAML struct {
	Name      string `yaml:"name"`
	Project   string `yaml:"project"`
	Lang      string `yaml:"lang"`
	Category  string `yaml:"category"`
	BuildArgs []Args `yaml:"build_args"`
}

type Args struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

func GetServiceYAML(filepath string) (ServiceYAML, error) {
	var result ServiceYAML
	filedata, err := ioutil.ReadFile(filepath)
	if err != nil {
		return result, err
	}

	err = yaml.Unmarshal(filedata, &result)
	if err != nil {
		return result, err
	}
	return result, err
}
