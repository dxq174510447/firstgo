package v1

import (
	"firstgo/frame/context"
	"firstgo/frame/http"
	"firstgo/frame/proxy"
	vo2 "firstgo/frame/vo"
	"firstgo/povo/po"
	"firstgo/povo/vo"
	"firstgo/service/impl"
	"firstgo/util"
	"fmt"
)

// UsersController 不要直接初始化 首字母大写代表类
type UsersController struct {
	usersService  *impl.UsersService
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
func (c *UsersController) ChangeStatus(
	local *context.LocalStack,
	id int, status int,
	requestId int,
	yid int, ystatus int) *vo2.JsonResult {
	fmt.Println(id, status, requestId, yid, ystatus)
	c.ChangeStatus_(local, id, status, c)
	return util.JsonUtil.BuildJsonSuccess(nil)
}

func (c *UsersController) ProxyTarget() *proxy.ProxyClass {
	return c.Proxy_
}

// UsersControllerImpl 控制器单例
var userController UsersController = UsersController{
	Proxy_: &proxy.ProxyClass{
		Annotations: []*proxy.AnnotationClass{
			http.NewRestAnnotation("/v1/users", "", "", "", "", ""),
		},
		Methods: []*proxy.ProxyMethod{
			{
				Name: "Save",
				Annotations: []*proxy.AnnotationClass{
					http.NewRestAnnotation("/", "post", "", "", "", ""),
				},
			},
			{
				Name: "Update",
				Annotations: []*proxy.AnnotationClass{
					http.NewRestAnnotation("/", "put", "", "", "", ""),
				},
			},
			{
				Name: "Delete",
				Annotations: []*proxy.AnnotationClass{
					http.NewRestAnnotation("/", "delete", "_,id", "", "", ""),
				},
			},
			{
				Name: "Get",
				Annotations: []*proxy.AnnotationClass{
					http.NewRestAnnotation("/", "get", "_,id", "", "", ""),
				},
			},
			{
				Name: "List",
				Annotations: []*proxy.AnnotationClass{
					http.NewRestAnnotation("/list", "post", "", "", "", ""),
				},
			},
			{
				Name: "ChangeStatus",
				Annotations: []*proxy.AnnotationClass{
					http.NewRestAnnotation("/change/{yid}/status/{ystatus}", "post",
						"_,id,status,_,_,_",
						"_,_,_,_,yid,ystatus",
						"_,_,_,requestId,_,_",
						""),
				},
			},
		},
	},
	List_: func(local *context.LocalStack, param *vo.UsersParam, self *UsersController) *vo.UsersPage {
		result, _ := self.usersService.List(local, param)
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

		result, _ := self.usersService.Save(local, data)
		return result
	},
	Get_: func(local *context.LocalStack, id int, self *UsersController) *vo.UsersVo {
		result, _ := self.usersService.Get(local, id)
		return result
	},
	Delete_: func(local *context.LocalStack, id int, self *UsersController) {
		self.usersService.Delete(local, id)
	},
	Update_: func(local *context.LocalStack, param *vo.UsersEdit, self *UsersController) *vo.UsersVo {
		var data *po.Users = &po.Users{}
		util.JsonUtil.Copy(param, data)

		result, _ := self.usersService.Update(local, data)
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
	http.AddControllerProxyTarget(proxy.ProxyTarger(&userController))
	// 初始化
	userController.usersService = impl.GetUsersService()

}
