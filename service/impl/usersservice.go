package impl

import (
	"firstgo/dao"
	"firstgo/frame/context"
	"firstgo/frame/db/dbcore"
	"firstgo/frame/exception"
	"firstgo/frame/proxy"
	"firstgo/povo/po"
	"firstgo/povo/vo"
)

type UsersService struct {
	usersDao      *dao.UsersDao
	Proxy_        *proxy.ProxyClass
	Save_         func(local *context.LocalStack, data *po.Users, self *UsersService) *vo.UsersVo
	Update_       func(local *context.LocalStack, data *po.Users, self *UsersService) *vo.UsersVo
	Delete_       func(local *context.LocalStack, id int, self *UsersService)
	Get_          func(local *context.LocalStack, id int, self *UsersService) *vo.UsersVo
	ChangeStatus_ func(local *context.LocalStack, id int, status int, self *UsersService)
	List_         func(local *context.LocalStack, param *vo.UsersParam, self *UsersService) *vo.UsersPage
}

func (c *UsersService) Save(local *context.LocalStack, data *po.Users) *vo.UsersVo {
	return c.Save_(local, data, c)
}

func (c *UsersService) Update(local *context.LocalStack, data *po.Users) *vo.UsersVo {
	return c.Update_(local, data, c)
}

func (c *UsersService) Delete(local *context.LocalStack, id int) {
	c.Delete_(local, id, c)
}

func (c *UsersService) Get(local *context.LocalStack, id int) *vo.UsersVo {
	return c.Get_(local, id, c)
}

func (c *UsersService) ChangeStatus(local *context.LocalStack, id int, status int) {
	c.ChangeStatus_(local, id, status, c)
}
func (c *UsersService) List(local *context.LocalStack, param *vo.UsersParam) *vo.UsersPage {
	return c.List_(local, param, c)
}

func (c *UsersService) ProxyTarget() *proxy.ProxyClass {
	return c.Proxy_
}

var usersService UsersService = UsersService{
	Proxy_: &proxy.ProxyClass{
		Annotations: proxy.NewSingleAnnotation(proxy.AnnotationService, nil),
		Methods: []*proxy.ProxyMethod{
			&proxy.ProxyMethod{
				Name:        "Save",
				Annotations: proxy.NewSingleAnnotation(dbcore.TransactionRequire, nil),
			},
			&proxy.ProxyMethod{
				Name:        "Update",
				Annotations: proxy.NewSingleAnnotation(dbcore.TransactionRequire, nil),
			},
			&proxy.ProxyMethod{
				Name:        "Delete",
				Annotations: proxy.NewSingleAnnotation(dbcore.TransactionRequire, nil),
			},
			&proxy.ProxyMethod{
				Name:        "ChangeStatus",
				Annotations: proxy.NewSingleAnnotation(dbcore.TransactionRequire, nil),
			},
		},
	},
	Get_: func(local *context.LocalStack, id int, self *UsersService) *vo.UsersVo {
		return self.usersDao.Get(local, id)
	},
	Delete_: func(local *context.LocalStack, id int, self *UsersService) {
		self.usersDao.Delete(local, id)
	},
	List_: func(local *context.LocalStack, param *vo.UsersParam, self *UsersService) *vo.UsersPage {
		return self.usersDao.List(local, param)
	},
	Update_: func(local *context.LocalStack, data *po.Users, self *UsersService) *vo.UsersVo {
		var total int = self.usersDao.FindByNameExcludeId(local, data.Name, data.Id)
		if total > 0 {
			panic(exception.NewException(502, "名称重复"))
		}

		self.usersDao.Update(local, data)
		return self.Get(local, data.Id)
	},
	Save_: func(local *context.LocalStack, data *po.Users, self *UsersService) *vo.UsersVo {
		var total int = self.usersDao.FindByName(local, data.Name)
		if total > 0 {
			panic(exception.NewException(502, "名称重复"))
		}

		self.usersDao.Save(local, data)

		if data.Status == -1 {
			panic(exception.NewException(502, "状态不正确"))
		}

		return self.Get(local, data.Id)
	},
	ChangeStatus_: func(local *context.LocalStack, id int, status int, self *UsersService) {
		self.usersDao.ChangeStatus(local, id, status)
	},
}

func GetUsersService() *UsersService {
	return &usersService
}

func init() {

	proxy.AddClassProxy(proxy.ProxyTarger(&usersService))

	// 初始化
	usersService.usersDao = dao.GetUsersDao()

}