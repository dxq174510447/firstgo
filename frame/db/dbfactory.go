package db

import (
	"context"
	"database/sql"
	"firstgo/util"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// dbRouter 关键字 对应的db源
var databaseRouter = make(map[string]*sql.DB)

func AddDatabaseRouter(key string, db *sql.DB) {
	databaseRouter[key] = db
}

type DatabaseFactory struct {
	dbUser string
	dbPwd  string
	dbName string
	dbPort string
	dbHost string
}

func (c *DatabaseFactory) NewDatabase() *sql.DB {
	//user:password@tcp(localhost:5555)/dbname?characterEncoding=UTF-8
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		c.dbUser, c.dbPwd, c.dbPort, c.dbHost, c.dbName,
	)
	//fmt.Println(url)
	db, _ := sql.Open("mysql", url)
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(1000)
	db.SetMaxIdleConns(20)
	return db
}

type DatabaseConnection struct {
	Db          *sql.DB
	Con         *sql.Conn
	Ctx         context.Context
	TxOpt       *sql.TxOptions
	Transaction *sql.Tx
}

func (d *DatabaseConnection) Close() {
	//no need toClose
}

func (d *DatabaseConnection) BeginTransaction() {
	//SET autocommit = 0
	//tx,err := d.Con.BeginTx(d.Ctx,d.TxOpt) //好像版本有问题 触发失败
	//if err != nil {
	//	panic(err)
	//}
	//d.Transaction = tx
	d.Con.ExecContext(d.Ctx, "begin")
}

func (d *DatabaseConnection) Commit() {
	//d.Transaction.Commit()
	d.Con.ExecContext(d.Ctx, "commit")
}

func (d *DatabaseConnection) Rollback() {
	//d.Transaction.Rollback()
	d.Con.ExecContext(d.Ctx, "rollback")
}

// OpenSqlConnection 是否只读 1是 0否
func OpenSqlConnection(readOnly int) *DatabaseConnection {

	ctx := context.Background()
	txOpt := sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
		ReadOnly:  readOnly == 1,
	}

	// 取默认的key
	var key string = DataBaseDefaultKey
	var db *sql.DB = databaseRouter[key]
	conn, _ := db.Conn(ctx)
	return &DatabaseConnection{
		Db:    db,
		Con:   conn,
		Ctx:   ctx,
		TxOpt: &txOpt,
	}
}

func init() {

	// 初始化默认数据源
	var defaultFactory DatabaseFactory = DatabaseFactory{
		dbUser: util.ConfigUtil.Get("DB_USER", "platform"),
		dbPwd:  util.ConfigUtil.Get("DB_PASSWORD", "xxcxcx"),
		dbName: util.ConfigUtil.Get("DB_NAME", "plat_base1"),
		dbPort: util.ConfigUtil.Get("DB_PORT", "3306"),
		dbHost: util.ConfigUtil.Get("DB_HOST", "rm-bp1thh63s5tx33q0kio.mysql.rds.aliyuncs.com"),
	}
	db := defaultFactory.NewDatabase()
	AddDatabaseRouter(DataBaseDefaultKey, db)

}
