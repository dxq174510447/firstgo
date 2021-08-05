package main

import (
	"firstgo/src/main/golang/com/firstgo/povo/po"
	"github.com/dxq174510447/goframe/lib/frame/context"
)

type IUserServiceProxy struct {
	ProxyImpl IUserService `FrameAutowired:""`
}

func (u *IUserServiceProxy) Save(local *context.LocalStack, user []*po.User, m *int, d1 map[*po.User][]*po.User, a string) (int, *po.User, map[*po.User][]*po.User) {
	var r1 int
	var r2 *po.User
	var r3 map[*po.User][]*po.User
	ProxyInteter.NewInvoker.Param(local, user, m, d1, a).Invoker(func() (int, *po.User, map[*po.User][]*po.User) {
		r1, r2, r3 := u.ProxyImpl.Save(local, user, m, d1, a)
		return r1, r2, r3
	})
	return r1, r2, r3
}

var iUserServiceProxy IUserServiceProxy = IUserServiceProxy{}

func init() {
	_ = IUserService(&iUserServiceProxy)
}
