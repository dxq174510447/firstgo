package main

import (
	_ "firstgo/controller"
	_ "firstgo/dao"
	_ "firstgo/service"
	"firstgo/util"
	_ "github.com/dxq174510447/goframe/lib/frame"
	"github.com/dxq174510447/goframe/lib/frame/db/dbcore"
	http "net/http"
)

func init() {
	// 初始化默认数据源
	var defaultFactory dbcore.DatabaseFactory = dbcore.DatabaseFactory{
		DbUser: util.ConfigUtil.Get("DB_USER", "platform"),
		DbPwd:  util.ConfigUtil.Get("DB_PASSWORD", "xxcxcx"),
		DbName: util.ConfigUtil.Get("DB_NAME", "plat_base1"),
		DbPort: util.ConfigUtil.Get("DB_PORT", "3306"),
		DbHost: util.ConfigUtil.Get("DB_HOST", "rm-bp1thh63s5tx33q0kio.mysql.rds.aliyuncs.com"),
	}
	db := defaultFactory.NewDatabase()
	dbcore.AddDatabaseRouter(dbcore.DataBaseDefaultKey, db)
}

func main() {
	//defer statsd.StartHttpHandlerStatsCollector("service/klook-p2pbackend/p2pbackend", "01",
	//	";;nats://ip-172-31-20-207.ap-southeast-1.compute.internal:6222").Stop()
	http.ListenAndServe(":8080", nil)
}
