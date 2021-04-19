package controller

import (
	"firstgo/frame"
	vo2 "firstgo/frame/vo"
	"firstgo/povo/po"
	"firstgo/povo/vo"
	"firstgo/service"
	"firstgo/util"
)

type UsersController struct {
	usersService *service.UsersService
}

// Save 新增
func (c *UsersController) Save(local *frame.FrameStack, param *vo.UsersAdd) *vo2.JsonResult {
	var data *po.Users = &po.Users{}
	util.JsonUtil.Copy(param, data)
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
		return c.usersService.Save(local, data)
	}
	result := f()
	return util.JsonUtil.BuildJsonSuccess(result)
}

// Update 修改
func (c *UsersController) Update(local *frame.FrameStack, param *vo.UsersEdit) *vo2.JsonResult {
	var data *po.Users = &po.Users{}
	util.JsonUtil.Copy(param, data)
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
		return c.usersService.Update(local, data)
	}
	result := f()
	return util.JsonUtil.BuildJsonSuccess(result)
}

// Delete 删除
func (c *UsersController) Delete(local *frame.FrameStack, id int) *vo2.JsonResult {
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
		c.usersService.Delete(local, id)
	}
	f()
	return util.JsonUtil.BuildJsonSuccess(nil)
}

// Get 查看
func (c *UsersController) Get(local *frame.FrameStack, id int) *vo2.JsonResult {
	var f func() *vo.UsersVo = func() *vo.UsersVo {
		conn := frame.OpenSqlConnection(1)

		local.Push()
		local.Set(frame.DbConnectKey, conn)

		defer func() {
			local.Pop()
		}()
		return c.usersService.Get(local, id)
	}
	result := f()
	return util.JsonUtil.BuildJsonSuccess(result)
}

// List 列表
func (c *UsersController) List(local *frame.FrameStack, param *vo.UsersParam) *vo2.JsonResult {
	var f func() ([]*vo.UsersVo, int) = func() ([]*vo.UsersVo, int) {
		conn := frame.OpenSqlConnection(1)

		local.Push()
		local.Set(frame.DbConnectKey, conn)

		defer func() {
			local.Pop()
		}()
		return c.usersService.List(local, param)
	}
	result, total := f()
	return util.JsonUtil.BuildJsonArraySuccess(result, total)
}

// ChangeStatus 变更状态
func (c *UsersController) ChangeStatus(local *frame.FrameStack, id int, status int) *vo2.JsonResult {
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
		c.usersService.ChangeStatus(local, id, status)
	}
	f()
	return util.JsonUtil.BuildJsonSuccess(nil)
}

// UsersControllerImpl 控制器单例
var UsersControllerImpl UsersController = UsersController{}

// UsersRequestController 控制器请求配置
var UsersRequestController frame.RequestController = frame.RequestController{
	HttpPath: "/users",
	Target:   &UsersControllerImpl,
	Methods: []frame.RequestMethod{
		{
			HttpMethod: "post",
			HttpPath:   "/",
			MethodName: "Save",
		},
		{
			HttpMethod: "put",
			HttpPath:   "/",
			MethodName: "Update",
		},
		{
			HttpMethod:     "delete",
			HttpPath:       "/",
			MethodName:     "Delete",
			MethodParamter: "_,id",
		},
		{
			HttpMethod:     "get",
			HttpPath:       "/",
			MethodName:     "Get",
			MethodParamter: "_,id",
		},
		{
			HttpMethod: "post",
			HttpPath:   "/list",
			MethodName: "List",
		},
		{
			HttpMethod:     "post",
			HttpPath:       "/status/change",
			MethodName:     "ChangeStatus",
			MethodParamter: "_,id,status",
		},
	},
}

func init() {

	// 初始化service
	UsersControllerImpl.usersService = &service.UsersServiceImpl

	// 初始化请求路由
	frame.DispatchServlet.AddRequestMapping(&UsersRequestController)
}
