package resources

import "github.com/dxq174510447/goframe/lib/frame/application"

var localYaml string = `
server:
  servlet:
    contextPath: ${APPLICATION_PATH:/api/base}
`

func init() {
	application.GetResourcePool().AddConfigYaml(application.ApplicationLocalYaml,localYaml)
}
