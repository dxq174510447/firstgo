package service

import (
	"firstgo/frame"
	"firstgo/povo/po"
	"firstgo/povo/vo"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

import _ "firstgo/frame"

func TestUsersService_Save(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	c1 := rand.Intn(100)
	c2 := rand.Intn(100)
	name := fmt.Sprintf("hello world%d-%d", c1, c2)
	data := &po.Users{Name: name, Password: "111", Status: 1}
	local := frame.NewFrameStack()

	var f func() *vo.UsersVo = func() *vo.UsersVo {
		conn := frame.OpenSqlConnection(0)
		conn.BeginTransaction()

		local.Push()
		local.Set(frame.DbConnectKey, conn)

		defer func() {
			local.Pop()
			if err := recover(); err != nil {
				conn.Rollback()
				panic(err)
			} else {
				conn.Commit()
			}
		}()
		return UsersServiceImpl.Save(local, data)
	}
	result := f()

	t.Log(result)
}

func TestUsersService_Update(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	c1 := rand.Intn(100)
	c2 := rand.Intn(100)
	name := fmt.Sprintf("hello world%d-%d", c1, c2)
	data := &po.Users{Name: name, Password: "111", Status: 1, Id: 68}
	local := frame.NewFrameStack()

	var f func() *vo.UsersVo = func() *vo.UsersVo {
		conn := frame.OpenSqlConnection(0)
		conn.BeginTransaction()

		local.Push()
		local.Set(frame.DbConnectKey, conn)

		defer func() {
			local.Pop()
			if err := recover(); err != nil {
				conn.Rollback()
				panic(err)
			} else {
				conn.Commit()
			}
		}()
		return UsersServiceImpl.Update(local, data)
	}
	result := f()

	t.Log(result)
}

func TestUsersService_Delete(t *testing.T) {

	local := frame.NewFrameStack()

	var f func() = func() {
		conn := frame.OpenSqlConnection(0)
		conn.BeginTransaction()

		local.Push()
		local.Set(frame.DbConnectKey, conn)

		defer func() {
			local.Pop()
			if err := recover(); err != nil {
				conn.Rollback()
				panic(err)
			} else {
				conn.Commit()
			}
		}()
		UsersServiceImpl.Delete(local, 67)
	}
	f()

	t.Log("delete", 67)
}

func TestUsersService_Get(t *testing.T) {
	local := frame.NewFrameStack()
	var f func() *vo.UsersVo = func() *vo.UsersVo {
		conn := frame.OpenSqlConnection(1)

		local.Push()
		local.Set(frame.DbConnectKey, conn)

		defer func() {
			local.Pop()
		}()
		return UsersServiceImpl.Get(local, 68)
	}
	result := f()
	t.Log(result)
}

func TestUsersService_ChangeStatus(t *testing.T) {
	local := frame.NewFrameStack()

	var f func() = func() {
		conn := frame.OpenSqlConnection(0)
		conn.BeginTransaction()

		local.Push()
		local.Set(frame.DbConnectKey, conn)

		defer func() {
			local.Pop()
			if err := recover(); err != nil {
				conn.Rollback()
				panic(err)
			} else {
				conn.Commit()
			}
		}()
		UsersServiceImpl.ChangeStatus(local, 68, 2)
	}
	f()

	t.Log("change status", 68)
}

func TestUsersService_List(t *testing.T) {

	local := frame.NewFrameStack()

	var f func() ([]*vo.UsersVo, int) = func() ([]*vo.UsersVo, int) {
		conn := frame.OpenSqlConnection(1)

		local.Push()
		local.Set(frame.DbConnectKey, conn)

		defer func() {
			local.Pop()
		}()
		param := &vo.UsersParam{Page: 1, PageSize: 10}
		return UsersServiceImpl.List(local, param)
	}
	result, total := f()
	t.Log(total, len(result))

}
