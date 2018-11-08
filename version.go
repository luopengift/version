package version

import (
	"fmt"
)

// export build var
var (
	VERSION string
	TIME    string
	APP     string
	GIT     string
)

// String to string
func String() string {
	return fmt.Sprintf("version: %v, buildTime: %v, buildTag: %v", VERSION, TIME, GIT)
}
