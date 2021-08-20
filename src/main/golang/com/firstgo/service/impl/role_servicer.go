package impl

import (
	"context"
	"firstgo/src/main/golang/com/firstgo/povo/po"
)

type RoleServicer interface {
	AddRole(local context.Context, user *po.User) (*po.User, error)
}
