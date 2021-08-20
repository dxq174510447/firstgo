package impl

import (
	"context"
	"firstgo/src/main/golang/com/firstgo/povo/po"
	"fmt"
)

type RoleServicerProxy struct {
}

func (r *RoleServicerProxy) AddRole(local context.Context, user *po.User) (*po.User, error) {
	fmt.Println("RoleServicerProxy.AddRole")
	return nil, nil
}

func init() {
	tmp := RoleServicerProxy{}
	_ = RoleServicer(&tmp)
}
