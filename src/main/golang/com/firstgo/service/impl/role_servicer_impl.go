package impl

import (
	"context"
	"firstgo/src/main/golang/com/firstgo/povo/po"
	"fmt"
)

type RoleServicerImpl struct {
}

func (r *RoleServicerImpl) AddRole(local context.Context, user *po.User) (*po.User, error) {
	fmt.Println("RoleServicerImpl.AddRole")
	return nil, nil
}

func init() {
	tmp := RoleServicerImpl{}
	_ = RoleServicer(&tmp)
}
