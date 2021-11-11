package resources

import "goframe/lib/frame/application"

var localYaml string = `
server:
  servlet:
    contextPath: ${APPLICATION_PATH:/api/base}
`

func init() {
	application.GetResourcePool().AddConfigYaml(application.ApplicationLocalYaml,localYaml)
}
