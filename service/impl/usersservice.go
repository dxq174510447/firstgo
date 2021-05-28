package impl

import (
	"firstgo/dao"
	"firstgo/povo/po"
	"firstgo/povo/vo"
	"github.com/dxq174510447/goframe/lib/frame/application"
	context "github.com/dxq174510447/goframe/lib/frame/context"
	dbcore "github.com/dxq174510447/goframe/lib/frame/db/dbcore"
	exception "github.com/dxq174510447/goframe/lib/frame/exception"
	proxy "github.com/dxq174510447/goframe/lib/frame/proxy"
)

type UsersService struct {
	usersDao      *dao.UsersDao
	Proxy_        *proxy.ProxyClass
	Save_         func(local *context.LocalStack, data *po.Users, self *UsersService) (*vo.UsersVo, error)
	Update_       func(local *context.LocalStack, data *po.Users, self *UsersService) (*vo.UsersVo, error)
	Delete_       func(local *context.LocalStack, id int, self *UsersService) (int, error)
	Get_          func(local *context.LocalStack, id int, self *UsersService) (*vo.UsersVo, error)
	ChangeStatus_ func(local *context.LocalStack, id int, status int, self *UsersService) (int, error)
	List_         func(local *context.LocalStack, param *vo.UsersParam, self *UsersService) (*vo.UsersPage, error)
}

func (c *UsersService) Save(local *context.LocalStack, data *po.Users) (*vo.UsersVo, error) {
	return c.Save_(local, data, c)
}

func (c *UsersService) Update(local *context.LocalStack, data *po.Users) (*vo.UsersVo, error) {
	return c.Update_(local, data, c)
}

func (c *UsersService) Delete(local *context.LocalStack, id int) {
	c.Delete_(local, id, c)
}

func (c *UsersService) Get(local *context.LocalStack, id int) (*vo.UsersVo, error) {
	return c.Get_(local, id, c)
}

func (c *UsersService) ChangeStatus(local *context.LocalStack, id int, status int) {
	c.ChangeStatus_(local, id, status, c)
}
func (c *UsersService) List(local *context.LocalStack, param *vo.UsersParam) (*vo.UsersPage, error) {
	return c.List_(local, param, c)
}

func (c *UsersService) ProxyTarget() *proxy.ProxyClass {
	return c.Proxy_
}

var usersService UsersService = UsersService{
	Proxy_: &proxy.ProxyClass{
		Annotations: []*proxy.AnnotationClass{
			proxy.NewSingleAnnotation(proxy.AnnotationService, nil),
		},
		Methods: []*proxy.ProxyMethod{
			&proxy.ProxyMethod{
				Name: "Save",
				Annotations: []*proxy.AnnotationClass{
					proxy.NewSingleAnnotation(dbcore.TransactionRequire, nil),
				},
			},
			&proxy.ProxyMethod{
				Name: "Update",
				Annotations: []*proxy.AnnotationClass{
					proxy.NewSingleAnnotation(dbcore.TransactionRequire, nil),
				},
			},
			&proxy.ProxyMethod{
				Name: "Delete",
				Annotations: []*proxy.AnnotationClass{
					proxy.NewSingleAnnotation(dbcore.TransactionRequire, nil),
				},
			},
			&proxy.ProxyMethod{
				Name: "ChangeStatus",
				Annotations: []*proxy.AnnotationClass{
					proxy.NewSingleAnnotation(dbcore.TransactionRequire, nil),
				},
			},
		},
	},
	Get_: func(local *context.LocalStack, id int, self *UsersService) (*vo.UsersVo, error) {
		return self.usersDao.Get1(local, id)
	},
	Delete_: func(local *context.LocalStack, id int, self *UsersService) (int, error) {
		return self.usersDao.Delete1(local, id)
	},
	List_: func(local *context.LocalStack, param *vo.UsersParam, self *UsersService) (*vo.UsersPage, error) {
		return self.usersDao.List1(local, param)
	},
	Update_: func(local *context.LocalStack, data *po.Users, self *UsersService) (*vo.UsersVo, error) {
		total, _ := self.usersDao.FindByNameExcludeId1(local, data.Name, data.Id)
		if total > 0 {
			panic(exception.NewException(502, "名称重复"))
		}

		self.usersDao.Update1(local, data)
		return self.Get(local, data.Id)
	},
	Save_: func(local *context.LocalStack, data *po.Users, self *UsersService) (*vo.UsersVo, error) {
		total, _ := self.usersDao.FindByName1(local, data.Name)
		if total > 0 {
			panic(exception.NewException(502, "名称重复"))
		}

		self.usersDao.Save1(local, data)

		if data.Status == -1 {
			panic(exception.NewException(502, "状态不正确"))
		}

		return self.Get(local, data.Id)
	},
	ChangeStatus_: func(local *context.LocalStack, id int, status int, self *UsersService) (int, error) {
		return self.usersDao.ChangeStatus1(local, id, status)
	},
}

func GetUsersService() *UsersService {
	return &usersService
}

func init() {
	application.AddProxyInstance("", proxy.ProxyTarger(&usersService))
}
