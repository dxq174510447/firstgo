package resources

import "goframe/lib/frame/application"

var defaultYaml string = `
server:
  port: ${APPLICATION_PORT:8080}
  servlet:
    contextPath: ${APPLICATION_PATH:/api/base1}
  access:
    inner: ${APPLICATION_INNER:http://127.0.0.1:8081}
    outer: ${APPLICATION_OUTER:https://wx.dev.chelizitech.com}
spring:
  application:
    name: ${APPLICATION_NAME:base-frontend}
  profiles:
    include: platform
platform:
  datasource:
    config:
      default:
        dbUser: ${DB_USER:platform}
        dbPwd: ${DB_PWD:xxxx}
        dbName: ${DB_NAME:plat_base1}
        dbPort: ${DB_PORT:3306}
        dbHost: ${DB_HOST:rm-bp1thh63s5tx33q0kio.mysql.rds.aliyuncs.com}
        prop: 
          a: b
          c: d
        proparray:
        - a
        - b
`

func init() {
	application.GetResourcePool().AddConfigYaml(application.ApplicationDefaultYaml,defaultYaml)
}
