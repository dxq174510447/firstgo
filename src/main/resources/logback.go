package resources

import "github.com/dxq174510447/goframe/lib/frame/application"

var logback = `
<configuration> 
  
  <appender name="CONSOLE" class="console">
    <encoder>
      <pattern>%date %date{2006-01-02} %date{2006-01-02T15:04:05Z07:00} %-5thread  %logger %logger{5} %-30logger{5} %thread %-5line %file %msg %n</pattern>
    </encoder>
  </appender>

	<appender name="CONSOLE1" class="console">
    <encoder>
      <pattern>%date [%-5thread] [%-5level] %-10logger{0}.%-5line %msg %n</pattern>
    </encoder>
  </appender>
  
  <logger name="goframe.lib.frame" level="debug" additivity="false">
	<appender-ref ref="CONSOLE1"/>
  </logger>
  <root level="debug">
	<appender-ref ref="CONSOLE1"/>
  </root>
</configuration>
`

func init() {
	application.AddAppLogConfig(application.ApplicationDefaultYaml, logback)
}
