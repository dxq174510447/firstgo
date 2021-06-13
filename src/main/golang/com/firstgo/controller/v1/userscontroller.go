package v1

import (
	"encoding/json"
	"firstgo/src/main/golang/com/firstgo/povo/po"
	"firstgo/src/main/golang/com/firstgo/povo/vo"
	"firstgo/src/main/golang/com/firstgo/service/impl"
	"firstgo/src/main/golang/com/firstgo/util"
	"fmt"
	"github.com/dxq174510447/goframe/lib/frame/application"
	"github.com/dxq174510447/goframe/lib/frame/context"
	"github.com/dxq174510447/goframe/lib/frame/db/dbcore"
	"github.com/dxq174510447/goframe/lib/frame/http"
	"github.com/dxq174510447/goframe/lib/frame/proxy/proxyclass"
	vo2 "github.com/dxq174510447/goframe/lib/frame/vo"
	"unsafe"
)

//	var setting map[string]*DatabaseConfig = make(map[string]*DatabaseConfig)
//	applicationContext.Environment.GetObjectValue("platform.datasource.config", setting)
// UsersController 不要直接初始化 首字母大写代表类
type UsersController struct {
	Logger           application.AppLoger              `FrameAutowired:""`
	UsersServiceImpl *impl.UsersService                `FrameAutowired:""`
	DbConfig         map[string]*dbcore.DatabaseConfig `FrameValue:"${platform.datasource.config}"`
	DefaultDbConfig  *dbcore.DatabaseConfig            `FrameValue:"${platform.datasource.config.default}"`
	ContextPath      string                            `FrameValue:"${server.servlet.contextPath:xxxxx}"`
	Proxy_           *proxyclass.ProxyClass
	Save_            func(local *context.LocalStack, param *vo.UsersAdd, self *UsersController) *vo.UsersVo
	Update_          func(local *context.LocalStack, param *vo.UsersEdit, self *UsersController) *vo.UsersVo
	Delete_          func(local *context.LocalStack, id int, self *UsersController)
	Get_             func(local *context.LocalStack, id int, self *UsersController) *vo.UsersVo
	List_            func(local *context.LocalStack, param *vo.UsersParam, self *UsersController) *vo.UsersPage
	ChangeStatus_    func(local *context.LocalStack, id int, status int, self *UsersController)
}

// Save 新增
func (c *UsersController) Save(local *context.LocalStack, param *vo.UsersAdd) *vo2.JsonResult {

	// test
	fmt.Println("contextPath-->", c.ContextPath)

	s1, _ := json.Marshal(c.DbConfig)
	fmt.Println("DbConfig-->", string(s1))

	s2, _ := json.Marshal(c.DefaultDbConfig)
	fmt.Println("DefaultDbConfig-->", string(s2), uintptr(unsafe.Pointer(c.DefaultDbConfig)))

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

	c.Logger.Info(local, "%s--%s", "123", "456")

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

func (c *UsersController) ProxyTarget() *proxyclass.ProxyClass {
	return c.Proxy_
}

// UsersControllerImpl 控制器单例
var userController UsersController = UsersController{
	Proxy_: &proxyclass.ProxyClass{
		Annotations: []*proxyclass.AnnotationClass{
			http.NewRestAnnotation(util.ConfigUtil.WrapServletPath("/v1/users"), "", "", "", "", ""),
		},
		Methods: []*proxyclass.ProxyMethod{
			{
				Name: "Save",
				Annotations: []*proxyclass.AnnotationClass{
					http.NewRestAnnotation("/", "post", "", "", "", ""),
				},
			},
			{
				Name: "Update",
				Annotations: []*proxyclass.AnnotationClass{
					http.NewRestAnnotation("/", "put", "", "", "", ""),
				},
			},
			{
				Name: "Delete",
				Annotations: []*proxyclass.AnnotationClass{
					http.NewRestAnnotation("/", "delete", "_,id", "", "", ""),
				},
			},
			{
				Name: "Get",
				Annotations: []*proxyclass.AnnotationClass{
					http.NewRestAnnotation("/", "get", "_,id", "", "", ""),
				},
			},
			{
				Name: "List",
				Annotations: []*proxyclass.AnnotationClass{
					http.NewRestAnnotation("/list", "post", "", "", "", ""),
				},
			},
			{
				Name: "ChangeStatus",
				Annotations: []*proxyclass.AnnotationClass{
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
		result, _ := self.UsersServiceImpl.List(local, param)
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

		result, _ := self.UsersServiceImpl.Save(local, data)
		fmt.Println(data.Id)
		return result
	},
	Get_: func(local *context.LocalStack, id int, self *UsersController) *vo.UsersVo {
		result, _ := self.UsersServiceImpl.Get(local, id)
		return result
	},
	Delete_: func(local *context.LocalStack, id int, self *UsersController) {
		self.UsersServiceImpl.Delete(local, id)
	},
	Update_: func(local *context.LocalStack, param *vo.UsersEdit, self *UsersController) *vo.UsersVo {
		var data *po.Users = &po.Users{}
		util.JsonUtil.Copy(param, data)

		result, _ := self.UsersServiceImpl.Update(local, data)
		return result
	},
	ChangeStatus_: func(local *context.LocalStack, id int, status int, self *UsersController) {
		self.UsersServiceImpl.ChangeStatus(local, id, status)
	},
}

func init() {
	application.AddProxyInstance("", proxyclass.ProxyTarger(&userController))
}
