package controller

import (
	"firstgo/frame/context"
	vo2 "firstgo/frame/vo"
	"firstgo/povo/po"
	"firstgo/povo/vo"
	"firstgo/service"
	"firstgo/util"
)

// UsersController 不要直接初始化 首字母大写代表类
type UsersController struct {
	usersService *service.UsersService
}

// Save 新增
func (c *UsersController) Save(local *context.LocalStack, param *vo.UsersAdd) *vo2.JsonResult {
	var data *po.Users = &po.Users{}
	util.JsonUtil.Copy(param, data)

	result := c.usersService.Save(local, data)
	return util.JsonUtil.BuildJsonSuccess(result)
}

// Update 修改
func (c *UsersController) Update(local *context.LocalStack, param *vo.UsersEdit) *vo2.JsonResult {
	var data *po.Users = &po.Users{}
	util.JsonUtil.Copy(param, data)

	result := c.usersService.Update(local, data)
	return util.JsonUtil.BuildJsonSuccess(result)
}

// Delete 删除
func (c *UsersController) Delete(local *context.LocalStack, id int) *vo2.JsonResult {
	c.usersService.Delete(local, id)
	return util.JsonUtil.BuildJsonSuccess(nil)
}

// Get 查看
func (c *UsersController) Get(local *context.LocalStack, id int) *vo2.JsonResult {
	result := c.usersService.Get(local, id)
	return util.JsonUtil.BuildJsonSuccess(result)
}

// List 列表
func (c *UsersController) List(local *context.LocalStack, param *vo.UsersParam) *vo2.JsonResult {
	result, total := c.usersService.List(local, param)
	return util.JsonUtil.BuildJsonArraySuccess(result, total)
}

// ChangeStatus 变更状态
func (c *UsersController) ChangeStatus(local *context.LocalStack, id int, status int) *vo2.JsonResult {
	c.usersService.ChangeStatus(local, id, status)
	return util.JsonUtil.BuildJsonSuccess(nil)
}

// UsersControllerImpl 控制器单例
var userController UsersController = UsersController{}

func GetUserController() *UsersController {
	return &userController
}

func init() {

	// 初始化
	userController.usersService = service.GetUsersService()

	//http.DispatchServlet.AddRequestMapping(&userController)
}
