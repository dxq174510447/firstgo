package dao

import (
	"firstgo/frame/context"
	"firstgo/frame/db/dbcore"
	"firstgo/frame/proxy"
	"firstgo/povo/po"
	"firstgo/povo/vo"
)

type UsersDao struct {
	Proxy_               *proxy.ProxyClass
	Find1_               func(local *context.LocalStack) (int, error)
	Find2_               func(local *context.LocalStack, status int) (int, error)
	Find3_               func(local *context.LocalStack, status int, name string) (int, error)
	Find4_               func(local *context.LocalStack, user *po.Users) (int, error)
	Find5_               func(local *context.LocalStack, user *po.Users, status int) ([]*po.Users, error)
	Find6_               func(local *context.LocalStack, user *po.Users, user1 *po.Users) (int, error)
	Save_                func(local *context.LocalStack, data *po.Users, self *UsersDao) int
	Update_              func(local *context.LocalStack, data *po.Users, self *UsersDao) int
	Delete_              func(local *context.LocalStack, id int, self *UsersDao) int
	Get_                 func(local *context.LocalStack, id int, self *UsersDao) *vo.UsersVo
	ChangeStatus_        func(local *context.LocalStack, id int, status int, self *UsersDao) int
	List_                func(local *context.LocalStack, param *vo.UsersParam, self *UsersDao) *vo.UsersPage
	FindByNameExcludeId_ func(local *context.LocalStack, name string, id int, self *UsersDao) int
	FindByName_          func(local *context.LocalStack, name string, self *UsersDao) int
}

func (c *UsersDao) Save(local *context.LocalStack, data *po.Users) int {
	return c.Save_(local, data, c)
}

func (c *UsersDao) Find1(local *context.LocalStack) (int, error) {
	return c.Find1_(local)
}
func (c *UsersDao) Find2(local *context.LocalStack, status int) (int, error) {
	return c.Find2_(local, status)
}
func (c *UsersDao) Find3(local *context.LocalStack, status int, name string) (int, error) {
	return c.Find3_(local, status, name)
}
func (c *UsersDao) Find4(local *context.LocalStack, user *po.Users) (int, error) {
	return c.Find4_(local, user)
}
func (c *UsersDao) Find5(local *context.LocalStack, user *po.Users, status int) ([]*po.Users, error) {
	return c.Find5_(local, user, status)
}
func (c *UsersDao) Find6(local *context.LocalStack, user *po.Users, user1 *po.Users) (int, error) {
	return c.Find6_(local, user, user1)
}

func (c *UsersDao) Update(local *context.LocalStack, data *po.Users) int {
	return c.Update_(local, data, c)
}

func (c *UsersDao) Delete(local *context.LocalStack, id int) int {
	return c.Delete_(local, id, c)
}

func (c *UsersDao) Get(local *context.LocalStack, id int) *vo.UsersVo {
	return c.Get_(local, id, c)
}

func (c *UsersDao) ChangeStatus(local *context.LocalStack, id int, status int) int {
	return c.ChangeStatus_(local, id, status, c)
}

func (c *UsersDao) List(local *context.LocalStack, param *vo.UsersParam) *vo.UsersPage {
	return c.List_(local, param, c)
}

func (c *UsersDao) FindByNameExcludeId(local *context.LocalStack, name string, id int) int {
	return c.FindByNameExcludeId_(local, name, id, c)
}

func (c *UsersDao) FindByName(local *context.LocalStack, name string) int {
	return c.FindByName_(local, name, c)
}

func (c *UsersDao) ProxyTarget() *proxy.ProxyClass {
	return c.Proxy_
}

var usersDao UsersDao = UsersDao{
	Proxy_: &proxy.ProxyClass{
		Annotations: proxy.NewSingleAnnotation(proxy.AnnotationDao, nil),
		Methods: []*proxy.ProxyMethod{
			&proxy.ProxyMethod{
				Name:        "Find2",
				Annotations: dbcore.NewSqlProvierConfigAnnotation("_,status"),
			},
			&proxy.ProxyMethod{
				Name:        "Find3",
				Annotations: dbcore.NewSqlProvierConfigAnnotation("_,status,name"),
			},
			&proxy.ProxyMethod{
				Name:        "Find4",
				Annotations: dbcore.NewSqlProvierConfigAnnotation("_,_"),
			},
			&proxy.ProxyMethod{
				Name:        "Find5",
				Annotations: dbcore.NewSqlProvierConfigAnnotation("_,user,status"),
			}, &proxy.ProxyMethod{
				Name:        "Find6",
				Annotations: dbcore.NewSqlProvierConfigAnnotation("_,user,user1"),
			},
		},
	},
	Save_: func(local *context.LocalStack, data *po.Users, self *UsersDao) int {
		con := local.Get(dbcore.DataBaseConnectKey).(*dbcore.DatabaseConnection)

		stmt, err := con.Con.PrepareContext(con.Ctx, "insert into users(name,password,status) values (?,?,?)")
		if err != nil {
			panic(err)
		}
		defer stmt.Close()
		result, err1 := stmt.Exec(data.Name, data.Password, data.Status)
		if err1 != nil {
			panic(err1)
		}
		ids, _ := result.LastInsertId()
		affect, _ := result.RowsAffected()
		data.Id = int(ids)
		return int(affect)
	},
	Update_: func(local *context.LocalStack, data *po.Users, self *UsersDao) int {
		con := local.Get(dbcore.DataBaseConnectKey).(*dbcore.DatabaseConnection)

		stmt, err := con.Con.PrepareContext(con.Ctx, "update users set name=?,password=?,status=? where id=?")
		if err != nil {
			panic(err)
		}
		defer stmt.Close()
		result, err1 := stmt.Exec(data.Name, data.Password, data.Status, data.Id)
		if err1 != nil {
			panic(err1)
		}
		affect, _ := result.RowsAffected()
		return int(affect)
	},
	Delete_: func(local *context.LocalStack, id int, self *UsersDao) int {
		con := local.Get(dbcore.DataBaseConnectKey).(*dbcore.DatabaseConnection)

		stmt, err := con.Con.PrepareContext(con.Ctx, "delete from users  where id=?")
		if err != nil {
			panic(err)
		}
		defer stmt.Close()
		result, err1 := stmt.Exec(id)
		if err1 != nil {
			panic(err1)
		}
		affect, _ := result.RowsAffected()
		return int(affect)
	},
	Get_: func(local *context.LocalStack, id int, self *UsersDao) *vo.UsersVo {
		con := local.Get(dbcore.DataBaseConnectKey).(*dbcore.DatabaseConnection)

		stmt, err := con.Con.PrepareContext(con.Ctx, "select id,name,status from users where id=?")
		if err != nil {
			panic(err)
		}
		defer stmt.Close()
		result := stmt.QueryRow(id)

		data := vo.UsersVo{}
		if err := result.Scan(&data.Id, &data.Name, &data.Status); err != nil {
			return nil
		}
		return &data
	},
	ChangeStatus_: func(local *context.LocalStack, id int, status int, self *UsersDao) int {
		con := local.Get(dbcore.DataBaseConnectKey).(*dbcore.DatabaseConnection)

		stmt, err := con.Con.PrepareContext(con.Ctx, "update users set status=? where id=?")
		if err != nil {
			panic(err)
		}
		defer stmt.Close()
		result, err1 := stmt.Exec(status, id)
		if err1 != nil {
			panic(err1)
		}
		affect, _ := result.RowsAffected()
		return int(affect)
	},
	List_: func(local *context.LocalStack, param *vo.UsersParam, self *UsersDao) *vo.UsersPage {
		con := local.Get(dbcore.DataBaseConnectKey).(*dbcore.DatabaseConnection)
		var pageSize int = 10
		if param.PageSize > 0 {
			pageSize = param.PageSize
		}
		var pageNo int = 1
		if param.Page >= 1 {
			pageNo = param.Page
		}

		var firstrow int = (pageNo - 1) * pageSize

		stmt, err := con.Con.PrepareContext(con.Ctx, "select id,name,status from users order by id desc limit ?,? ")
		if err != nil {
			panic(err)
		}
		defer func() {
			stmt.Close()
		}()
		result, err1 := stmt.Query(firstrow, pageSize)
		defer func() {
			if result != nil {
				result.Close() //可以关闭掉未scan连接一直占用
			}
		}()

		if err1 != nil {
			panic(err1)
		}
		dd := make([]*vo.UsersVo, pageSize)
		queryCount := 0
		for result.Next() {
			data := vo.UsersVo{}
			result.Scan(&data.Id, &data.Name, &data.Status) //不scan会导致连接不释放
			dd[queryCount] = &data
			queryCount++
		}

		if queryCount > 0 {
			dd = dd[0:queryCount]
		}

		stmt2, err2 := con.Con.PrepareContext(con.Ctx, "select count(id) from  users ")
		if err2 != nil {
			panic(err2)
		}
		defer func() {
			stmt2.Close()
		}()
		result2 := stmt2.QueryRow()

		var totalRow int = 0
		result2.Scan(&totalRow)
		return &vo.UsersPage{Total: totalRow, Data: dd}
	},
	FindByNameExcludeId_: func(local *context.LocalStack, name string, id int, self *UsersDao) int {
		con := local.Get(dbcore.DataBaseConnectKey).(*dbcore.DatabaseConnection)

		stmt2, err2 := con.Con.PrepareContext(con.Ctx, "select count(id) from  users where name = ? and id != ? ")
		if err2 != nil {
			panic(err2)
		}
		defer func() {
			stmt2.Close()
		}()
		result2 := stmt2.QueryRow(name, id)

		var totalRow int = 0
		result2.Scan(&totalRow)
		return totalRow
	},
	FindByName_: func(local *context.LocalStack, name string, self *UsersDao) int {
		con := local.Get(dbcore.DataBaseConnectKey).(*dbcore.DatabaseConnection)

		stmt2, err2 := con.Con.PrepareContext(con.Ctx, "select count(id) from  users where name = ? ")
		if err2 != nil {
			panic(err2)
		}
		defer func() {
			stmt2.Close()
		}()
		result2 := stmt2.QueryRow(name)

		var totalRow int = 0
		result2.Scan(&totalRow)
		return totalRow
	},
}

func GetUsersDao() *UsersDao {
	return &usersDao
}

func init() {
	dbcore.AddMapperProxyTarget(proxy.ProxyTarger(&usersDao), UsersXml)
}
