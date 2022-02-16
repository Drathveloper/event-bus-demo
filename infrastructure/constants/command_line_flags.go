package constants

import (
	"os"
)

const (
	ResourceRootFlag   = "resource-root"
	ConfigPrefixFlag   = "config-prefix"
	ActiveProfilesFlag = "active-profiles"
)

const (
	ResourceRootDefaultValue   = os.PathSeparator
	ConfigPrefixDefaultValue   = "application"
	ActiveProfilesDefaultValue = "default"
)
