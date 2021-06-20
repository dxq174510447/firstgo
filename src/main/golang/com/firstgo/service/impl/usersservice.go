package impl

import (
	"encoding/json"
	"firstgo/src/main/golang/com/firstgo/dao"
	"firstgo/src/main/golang/com/firstgo/povo/po"
	"firstgo/src/main/golang/com/firstgo/povo/vo"
	"fmt"
	"github.com/dxq174510447/goframe/lib/frame/application"
	context "github.com/dxq174510447/goframe/lib/frame/context"
	dbcore "github.com/dxq174510447/goframe/lib/frame/db/dbcore"
	exception "github.com/dxq174510447/goframe/lib/frame/exception"
	"github.com/dxq174510447/goframe/lib/frame/proxy/proxyclass"
	"unsafe"
)

type UsersService struct {
	UsersDaoImpl    *dao.UsersDao          `FrameAutowired:""`
	DefaultDbConfig *dbcore.DatabaseConfig `FrameValue:"${platform.datasource.config.default}"`
	Proxy_          *proxyclass.ProxyClass
	Save_           func(local *context.LocalStack, data *po.Users, self *UsersService) (*vo.UsersVo, error)
	Update_         func(local *context.LocalStack, data *po.Users, self *UsersService) (*vo.UsersVo, error)
	Delete_         func(local *context.LocalStack, id int, self *UsersService) (int, error)
	Get_            func(local *context.LocalStack, id int, self *UsersService) (*vo.UsersVo, error)
	ChangeStatus_   func(local *context.LocalStack, id int, status int, self *UsersService) (int, error)
	List_           func(local *context.LocalStack, param *vo.UsersParam, self *UsersService) (*vo.UsersPage, error)
}

func (c *UsersService) Save(local *context.LocalStack, data *po.Users) (*vo.UsersVo, error) {
	s2, _ := json.Marshal(c.DefaultDbConfig)
	fmt.Println("DefaultDbConfig-->", string(s2), uintptr(unsafe.Pointer(c.DefaultDbConfig)))

	return c.Save_(local, data, c)
}

func (c *UsersService) Update(local *context.LocalStack, data *po.Users) (*vo.UsersVo, error) {
	return c.Update_(local, data, c)
}

func (c *UsersService) Delete(local *context.LocalStack, id int) (int, error) {
	return c.Delete_(local, id, c)
}

func (c *UsersService) Get(local *context.LocalStack, id int) (*vo.UsersVo, error) {
	return c.Get_(local, id, c)
}

func (c *UsersService) ChangeStatus(local *context.LocalStack, id int, status int) error {
	c.ChangeStatus_(local, id, status, c)

	return nil
}

func (c *UsersService) List(local *context.LocalStack, param *vo.UsersParam) (*vo.UsersPage, error) {
	return c.List_(local, param, c)
}

func (c *UsersService) ProxyTarget() *proxyclass.ProxyClass {
	return c.Proxy_
}

var usersService UsersService = UsersService{
	Proxy_: &proxyclass.ProxyClass{
		Annotations: []*proxyclass.AnnotationClass{
			proxyclass.NewSingleAnnotation(proxyclass.AnnotationService, nil),
		},
		Methods: []*proxyclass.ProxyMethod{
			&proxyclass.ProxyMethod{
				Name: "Save",
				Annotations: []*proxyclass.AnnotationClass{
					proxyclass.NewSingleAnnotation(dbcore.TransactionRequire, nil),
				},
			},
			&proxyclass.ProxyMethod{
				Name: "Update",
				Annotations: []*proxyclass.AnnotationClass{
					proxyclass.NewSingleAnnotation(dbcore.TransactionRequire, nil),
				},
			},
			&proxyclass.ProxyMethod{
				Name: "Delete",
				Annotations: []*proxyclass.AnnotationClass{
					proxyclass.NewSingleAnnotation(dbcore.TransactionRequire, nil),
				},
			},
			&proxyclass.ProxyMethod{
				Name: "ChangeStatus",
				Annotations: []*proxyclass.AnnotationClass{
					proxyclass.NewSingleAnnotation(dbcore.TransactionRequire, nil),
				},
			},
		},
	},
	Get_: func(local *context.LocalStack, id int, self *UsersService) (*vo.UsersVo, error) {
		return self.UsersDaoImpl.Get1(local, id)
	},
	List_: func(local *context.LocalStack, param *vo.UsersParam, self *UsersService) (*vo.UsersPage, error) {
		return self.UsersDaoImpl.List1(local, param)
	},
	Update_: func(local *context.LocalStack, data *po.Users, self *UsersService) (*vo.UsersVo, error) {
		total, _ := self.UsersDaoImpl.FindByNameExcludeId1(local, data.Name, data.Id)
		if total > 0 {
			panic(exception.NewException(502, "名称重复"))
		}

		self.UsersDaoImpl.Update1(local, data)
		return self.Get(local, data.Id)
	},
	Save_: func(local *context.LocalStack, data *po.Users, self *UsersService) (*vo.UsersVo, error) {
		total, _ := self.UsersDaoImpl.FindByName1(local, data.Name)
		if total > 0 {
			panic(exception.NewException(502, "名称重复"))
		}

		self.UsersDaoImpl.Save(local, data)

		if data.Status == -1 {
			panic(exception.NewException(502, "状态不正确"))
		}

		return self.Get(local, data.Id)
	},
	Delete_: func(local *context.LocalStack, id int, self *UsersService) (int, error) {
		self.UsersDaoImpl.Delete1(local, id)
		panic("123123")
		return 1, exception.NewException(500, "12312")
	},
	ChangeStatus_: func(local *context.LocalStack, id int, status int, self *UsersService) (int, error) {
		self.UsersDaoImpl.ChangeStatus1(local, id, status)
		self.Delete(local, id)
		return 0, nil
	},
}

func init() {
	application.AddProxyInstance("", proxyclass.ProxyTarger(&usersService))
}
