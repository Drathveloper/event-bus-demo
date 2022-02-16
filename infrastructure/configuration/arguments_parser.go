package configuration

import (
	"event-bus-demo/infrastructure/constants"
	"flag"
	"fmt"
	"strings"
)

const resourceRootHint = "path where resources are stored."
const configPrefixHint = "prefix of yaml configuration file."
const activeProfilesHint = "profiles to be loaded from configuration comma separated. Priority given by order, having the first the less priority and the last the most priority."

func ParseInputArguments() Arguments {
	resourceRootFlag := flag.String(constants.ResourceRootFlag, fmt.Sprintf(".%c", constants.ResourceRootDefaultValue), resourceRootHint)
	configPrefixFlag := flag.String(constants.ConfigPrefixFlag, constants.ConfigPrefixDefaultValue, configPrefixHint)
	activeProfilesFlag := flag.String(constants.ActiveProfilesFlag, constants.ActiveProfilesDefaultValue, activeProfilesHint)
	flag.Parse()
	activeProfiles := strings.Split(*activeProfilesFlag, ",")
	return Arguments{
		*resourceRootFlag,
		*configPrefixFlag,
		activeProfiles,
	}
}
