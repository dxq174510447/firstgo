package resources

import (
	"github.com/dxq174510447/goframe/lib/frame/application"
	"github.com/dxq174510447/goframe/lib/frame/log"
)

var logConfig = `
<configuration> 
  
  <appender name="CONSOLE" class="console">
    <encoder>
      <pattern>%date %date{2006-01-02} %date{2006-01-02T15:04:05Z07:00} %-5thread  %logger %logger{5} %-30logger{5} %thread %-5line %file %msg %n</pattern>
    </encoder>
  </appender>

  <root level="debug">
	<appender-ref ref="CONSOLE"/>
  </root>
</configuration>
`

func init() {
	application.GetResourcePool().AddLogConfig(log.PlatLogKey,logConfig)
}
