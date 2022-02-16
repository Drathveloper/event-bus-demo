package configuration

import (
	"event-bus-demo/infrastructure/constants"
	infrastructure "event-bus-demo/infrastructure/error"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

type YamlParser interface {
	ReadConfiguration(arguments Arguments) (ApplicationConfiguration, infrastructure.InfrastructureError)
}

type yamlParser struct {
	validate *validator.Validate
}

func NewYamlParser(validate *validator.Validate) YamlParser {
	return &yamlParser{
		validate: validate,
	}
}

func (parser *yamlParser) ReadConfiguration(arguments Arguments) (ApplicationConfiguration, infrastructure.InfrastructureError) {
	var applicationConfiguration ApplicationConfiguration
	if yamlMap, err := parser.getYamlMap(arguments); err != nil {
		return ApplicationConfiguration{}, infrastructure.NewParseFileError(err.Error())
	} else if err = mapstructure.Decode(yamlMap, &applicationConfiguration); err != nil {
		return ApplicationConfiguration{}, infrastructure.NewParseFileError(err.Error())
	} else if err := parser.validate.Struct(applicationConfiguration); err != nil {
		return ApplicationConfiguration{}, infrastructure.NewParseFileError(err.Error())
	} else {
		return applicationConfiguration, nil
	}
}

func (parser *yamlParser) getYamlMap(arguments Arguments) (map[interface{}]interface{}, error) {
	var err error
	yamlMap := make(map[interface{}]interface{})
	for _, profile := range arguments.ActiveConfigurationProfiles {
		configurationPath := parser.generatePath(arguments.ResourceRoot, arguments.ConfigFilePrefix, profile)
		yamlMap, err = parser.readConfigurationToMap(configurationPath, yamlMap)
		if err != nil && profile == constants.DefaultProfile {
			return nil, err
		}
		// In case that we cannot find any associated config file to given profile we silently skip that file without
		// throwing any error
	}
	return yamlMap, nil
}

func (parser *yamlParser) readConfigurationToMap(configurationPath string, currentConfiguration map[interface{}]interface{}) (map[interface{}]interface{}, error) {
	newConfiguration := make(map[interface{}]interface{})
	if yamlConfiguration, err := ioutil.ReadFile(configurationPath); err != nil {
		return currentConfiguration, err
	} else if err := yaml.Unmarshal(yamlConfiguration, &newConfiguration); err != nil {
		return currentConfiguration, err
	} else {
		return parser.mergeMap(currentConfiguration, newConfiguration), nil
	}
}

func (parser *yamlParser) generatePath(path, prefix string, profile string) string {
	if !strings.HasSuffix(path, fmt.Sprintf("%c", os.PathSeparator)) {
		path = fmt.Sprintf("%s%c", path, os.PathSeparator)
	}
	if profile == "" || profile == constants.DefaultProfile {
		return fmt.Sprintf("%s%s.%s", path, prefix, constants.YamlExtension)
	}
	return fmt.Sprintf("%s%s-%s.%s", path, prefix, profile, constants.YamlExtension)
}

func (parser *yamlParser) mergeMap(map1, map2 map[interface{}]interface{}) map[interface{}]interface{} {
	for k, v := range map2 {
		switch map1[k].(type) {
		case map[interface{}]interface{}:
			map1[k] = parser.mergeMap(map1[k].(map[interface{}]interface{}), v.(map[interface{}]interface{}))
		default:
			map1[k] = v
		}
	}
	return map1
}
