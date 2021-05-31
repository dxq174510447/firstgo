package resources

import "github.com/dxq174510447/goframe/lib/frame/application"

var localYaml string = `
server:
  servlet:
    contextPath: ${APPLICATION_PATH:/api/v1}
`

func init() {
	application.AddConfigYaml(application.ApplicationLocalYaml, localYaml)
}
