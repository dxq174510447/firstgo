package main

import (
	_ "firstgo/controller"
	_ "firstgo/dao"
	_ "firstgo/service"
	_ "github.com/dxq174510447/goframe/lib/frame"
	"github.com/dxq174510447/goframe/lib/frame/application"
)

//func init() {
//	// 初始化默认数据源
//	var defaultFactory dbcore.DatabaseFactory = dbcore.DatabaseFactory{
//		DbUser: util.ConfigUtil.Get("DB_USER", "platform"),
//		DbPwd:  util.ConfigUtil.Get("DB_PASSWORD", "xxcxcx"),
//		DbName: util.ConfigUtil.Get("DB_NAME", "plat_base1"),
//		DbPort: util.ConfigUtil.Get("DB_PORT", "3306"),
//		DbHost: util.ConfigUtil.Get("DB_HOST", "rm-bp1thh63s5tx33q0kio.mysql.rds.aliyuncs.com"),
//	}
//	db := defaultFactory.NewDatabase()
//	dbcore.AddDatabaseRouter(dbcore.DataBaseDefaultKey, db)
//}

type FirstGo struct {
}

func (f *FirstGo) Run(args []string) {
	application.NewApplication(f).Run(args)
}

func main() {
	// http.ListenAndServe(":8080", nil)
	args := []string{"--appli=123"}

	var instance *FirstGo = &FirstGo{}
	instance.Run(args)
}
