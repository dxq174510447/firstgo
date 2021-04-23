package service

import (
	"firstgo/dao"
	"firstgo/frame/context"
	"firstgo/frame/db"
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
		Annotations: []*proxy.AnnotationClass{
			&proxy.AnnotationClass{
				Name: proxy.AnnotationService,
			},
		},
		Methods: []*proxy.ProxyMethod{
			&proxy.ProxyMethod{
				Name: "Save",
				Annotations: []*proxy.AnnotationClass{
					&proxy.AnnotationClass{
						Name: db.TransactionRequire,
					},
				},
			},
			&proxy.ProxyMethod{
				Name: "Update",
				Annotations: []*proxy.AnnotationClass{
					&proxy.AnnotationClass{
						Name: db.TransactionRequire,
					},
				},
			},
			&proxy.ProxyMethod{
				Name: "Delete",
				Annotations: []*proxy.AnnotationClass{
					&proxy.AnnotationClass{
						Name: db.TransactionRequire,
					},
				},
			},
			&proxy.ProxyMethod{
				Name: "ChangeStatus",
				Annotations: []*proxy.AnnotationClass{
					&proxy.AnnotationClass{
						Name: db.TransactionRequire,
					},
				},
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
