package controller

import (
	"firstgo/frame/context"
	"firstgo/frame/http"
	"firstgo/frame/proxy"
	vo2 "firstgo/frame/vo"
	"firstgo/povo/po"
	"firstgo/povo/vo"
	"firstgo/service"
	"firstgo/util"
	"fmt"
)

// UsersController 不要直接初始化 首字母大写代表类
type UsersController struct {
	usersService  *service.UsersService
	Proxy_        *proxy.ProxyClass
	Save_         func(local *context.LocalStack, param *vo.UsersAdd, self *UsersController) *vo.UsersVo
	Update_       func(local *context.LocalStack, param *vo.UsersEdit, self *UsersController) *vo.UsersVo
	Delete_       func(local *context.LocalStack, id int, self *UsersController)
	Get_          func(local *context.LocalStack, id int, self *UsersController) *vo.UsersVo
	List_         func(local *context.LocalStack, param *vo.UsersParam, self *UsersController) *vo.UsersPage
	ChangeStatus_ func(local *context.LocalStack, id int, status int, self *UsersController)
}

// Save 新增
func (c *UsersController) Save(local *context.LocalStack, param *vo.UsersAdd) *vo2.JsonResult {
	result := c.Save_(local, param, c)
	return util.JsonUtil.BuildJsonSuccess(result)
}

// Update 修改
func (c *UsersController) Update(local *context.LocalStack, param *vo.UsersEdit) *vo2.JsonResult {
	result := c.Update_(local, param, c)
	return util.JsonUtil.BuildJsonSuccess(result)
}

// Delete 删除
func (c *UsersController) Delete(local *context.LocalStack, id int) *vo2.JsonResult {
	c.Delete_(local, id, c)
	return util.JsonUtil.BuildJsonSuccess(nil)
}

// Get 查看
func (c *UsersController) Get(local *context.LocalStack, id int) *vo2.JsonResult {
	result := c.Get_(local, id, c)
	return util.JsonUtil.BuildJsonSuccess(result)
}

// List 列表
func (c *UsersController) List(local *context.LocalStack, param *vo.UsersParam) *vo2.JsonResult {
	result := c.List_(local, param, c)
	return util.JsonUtil.BuildJsonArraySuccess(result.Data, result.Total)
}

// ChangeStatus 变更状态
func (c *UsersController) ChangeStatus(local *context.LocalStack, id int, status int) *vo2.JsonResult {
	c.ChangeStatus_(local, id, status, c)
	return util.JsonUtil.BuildJsonSuccess(nil)
}

func (c *UsersController) ProxyTarget() *proxy.ProxyClass {
	return c.Proxy_
}

// UsersControllerImpl 控制器单例
var userController UsersController = UsersController{
	Proxy_: &proxy.ProxyClass{
		Annotations: http.NewRestAnnotation("/users", "", "", ""),
		Methods: []*proxy.ProxyMethod{
			&proxy.ProxyMethod{
				Name:        "Save",
				Annotations: http.NewRestAnnotation("/", "post", "", ""),
			},
			&proxy.ProxyMethod{
				Name:        "Update",
				Annotations: http.NewRestAnnotation("/", "put", "", ""),
			},
			&proxy.ProxyMethod{
				Name:        "Delete",
				Annotations: http.NewRestAnnotation("/", "delete", "_,id", ""),
			},
			&proxy.ProxyMethod{
				Name:        "Get",
				Annotations: http.NewRestAnnotation("/", "get", "_,id", ""),
			},
			&proxy.ProxyMethod{
				Name:        "List",
				Annotations: http.NewRestAnnotation("/list", "post", "", ""),
			},
			&proxy.ProxyMethod{
				Name:        "ChangeStatus",
				Annotations: http.NewRestAnnotation("/change/status", "post", "_,id,status", ""),
			},
		},
	},
	List_: func(local *context.LocalStack, param *vo.UsersParam, self *UsersController) *vo.UsersPage {
		result := self.usersService.List(local, param)
		return result
	},
	Save_: func(local *context.LocalStack, param *vo.UsersAdd, self *UsersController) *vo.UsersVo {

		user := util.WebUtil.GetThreadUser(local)

		if user == nil {
			fmt.Println("no user")
		} else {
			fmt.Println("current user", user.Name)
		}

		var data *po.Users = &po.Users{}
		util.JsonUtil.Copy(param, data)

		result := self.usersService.Save(local, data)
		return result
	},
	Get_: func(local *context.LocalStack, id int, self *UsersController) *vo.UsersVo {
		result := self.usersService.Get(local, id)
		return result
	},
	Delete_: func(local *context.LocalStack, id int, self *UsersController) {
		self.usersService.Delete(local, id)
	},
	Update_: func(local *context.LocalStack, param *vo.UsersEdit, self *UsersController) *vo.UsersVo {
		var data *po.Users = &po.Users{}
		util.JsonUtil.Copy(param, data)

		result := self.usersService.Update(local, data)
		return result
	},
	ChangeStatus_: func(local *context.LocalStack, id int, status int, self *UsersController) {
		self.usersService.ChangeStatus(local, id, status)
	},
}

func GetUserController() *UsersController {
	return &userController
}

func init() {

	http.AddRequestMapping(proxy.ProxyTarger(&userController))
	// 初始化
	userController.usersService = service.GetUsersService()

}
