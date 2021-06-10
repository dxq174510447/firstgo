package resources

import "github.com/dxq174510447/goframe/lib/frame/application"

var logback = `
`

func init() {
	application.AddConfigYaml(application.ApplicationDefaultYaml, logback)
}
