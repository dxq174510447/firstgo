package dao

import (
	"firstgo/frame"
	"firstgo/povo/po"
	"firstgo/povo/vo"
)

type UsersDao struct {
}

func (c *UsersDao) Save(local *frame.FrameStack, data *po.Users) int {
	con := local.Get(frame.DbConnectKey).(*frame.DbConnection)

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
}

func (c *UsersDao) Update(local *frame.FrameStack, data *po.Users) int {
	con := local.Get(frame.DbConnectKey).(*frame.DbConnection)

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
}

func (c *UsersDao) Delete(local *frame.FrameStack, id int) int {
	con := local.Get(frame.DbConnectKey).(*frame.DbConnection)

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
}

func (c *UsersDao) Get(local *frame.FrameStack, id int) *vo.UsersVo {
	con := local.Get(frame.DbConnectKey).(*frame.DbConnection)

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
}

func (c *UsersDao) ChangeStatus(local *frame.FrameStack, id int, status int) int {
	con := local.Get(frame.DbConnectKey).(*frame.DbConnection)

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
}

func (c *UsersDao) List(local *frame.FrameStack, param *vo.UsersParam) ([]*vo.UsersVo, int) {
	con := local.Get(frame.DbConnectKey).(*frame.DbConnection)
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
	dd := []*vo.UsersVo{}
	for result.Next() {
		data := vo.UsersVo{}
		result.Scan(&data.Id, &data.Name, &data.Status) //不scan会导致连接不释放
		dd = append(dd, &data)
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
	return dd, totalRow
}

var UsersDaoImpl UsersDao = UsersDao{}

func init() {

}
