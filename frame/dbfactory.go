package frame

import (
	"context"
	"database/sql"
	"firstgo/util"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// DbDefaultKey db路由选择 默认default 兼容后面路由选择
const DbDefaultKey = "default"

// DbConnectKey 当前goroute的包含的数据库连接key，在变量栈中的key
const DbConnectKey = "_db_connect"

//key --> dbname-tablename-method
//简单都orm
//var entityTpl = make(map[string]string)

// dbRouter 关键字 对应的db源
var dbRouter = make(map[string]*sql.DB)

func AddDbRouter(key string, db *sql.DB) {
	dbRouter[key] = db
}

type DbFactory struct {
	dbuser string
	dbpwd  string
	dbname string
	dbport string
	dbhost string
}

func (c *DbFactory) NewDatabase() *sql.DB {
	//user:password@tcp(localhost:5555)/dbname?characterEncoding=UTF-8
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		c.dbuser, c.dbpwd, c.dbhost, c.dbport, c.dbname,
	)
	//fmt.Println(url)
	db, _ := sql.Open("mysql", url)
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(1000)
	db.SetMaxIdleConns(20)
	return db
}

type DbConnection struct {
	Db          *sql.DB
	Con         *sql.Conn
	Ctx         context.Context
	TxOpt       *sql.TxOptions
	Transaction *sql.Tx
}

func (d *DbConnection) Close() {
	//no need toclose
}

func (d *DbConnection) BeginTransaction() {
	//SET autocommit = 0
	//tx,err := d.Con.BeginTx(d.Ctx,d.TxOpt) //好像版本有问题 触发失败
	//if err != nil {
	//	panic(err)
	//}
	//d.Transaction = tx
	d.Con.ExecContext(d.Ctx, "begin")
}

func (d *DbConnection) Commit() {
	//d.Transaction.Commit()
	d.Con.ExecContext(d.Ctx, "commit")
}

func (d *DbConnection) Rollback() {
	//d.Transaction.Rollback()
	d.Con.ExecContext(d.Ctx, "rollback")
}

//不用关闭连接 db里面有连接池
//是否只读 1是 0否
func OpenSqlConnection(readOnly int) *DbConnection {

	ctx := context.Background()
	txOpt := sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
		ReadOnly:  readOnly == 1,
	}

	// 取默认的key
	var key string = DbDefaultKey
	var db *sql.DB = dbRouter[key]
	conn, _ := db.Conn(ctx)
	return &DbConnection{
		Db:    db,
		Con:   conn,
		Ctx:   ctx,
		TxOpt: &txOpt,
	}
}

func init() {

	// 初始化默认数据源
	var defaultFactory DbFactory = DbFactory{
		dbuser: util.ConfigUtil.Get("DB_USER", "platform"),
		dbpwd:  util.ConfigUtil.Get("DB_PASSWORD", "xxx"),
		dbname: util.ConfigUtil.Get("DB_NAME", "plat_base1"),
		dbport: util.ConfigUtil.Get("DB_PORT", "3306"),
		dbhost: util.ConfigUtil.Get("DB_HOST", "rm-bp1thh63s5tx33q0kio.mysql.rds.aliyuncs.com"),
	}
	db := defaultFactory.NewDatabase()
	AddDbRouter(DbDefaultKey, db)

}
