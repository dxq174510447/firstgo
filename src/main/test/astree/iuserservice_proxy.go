package main

import (
	"firstgo/src/main/golang/com/firstgo/povo/po"
	"github.com/dxq174510447/goframe/lib/frame/context"
	"reflect"
)

type IUserServiceProxy struct {
	ProxyImpl IUserService `FrameAutowired:""`
}

func (u *IUserServiceProxy) GetProxyType(local *context.LocalStack) reflect.Type {
	return nil
}

func (u *IUserServiceProxy) GetProxyTarget(local *context.LocalStack) IUserService {
	return nil
}

func (u *IUserServiceProxy) Save(local *context.LocalStack, user []*po.User, m *int, d1 map[*po.User][]*po.User, a string) (int, *po.User, map[*po.User][]*po.User) {
	var impl interface{} = ApplicationContext.GetByInterface(local, reflect.TypeOf(IUserService))
	impl1 := impl.(IUserService)
	reflect.MakeFunc()
}

var iUserServiceProxy IUserServiceProxy = IUserServiceProxy{}

func init() {
	_ = IUserService(&iUserServiceProxy)
}
