package main

import (
	"firstgo/src/main/golang/com/firstgo/povo/po"
	"github.com/dxq174510447/goframe/lib/frame/context"
)

// Query
type IUserService interface {
	// Save path(sdfsfs)
	Save(local /*swagger2*/ *context.LocalStack, user /*swagger1*/ []*po.User, m *int, d1 /*swagger0*/ map[*po.User][]*po.User, a string) (int, *po.User, map[*po.User][]*po.User)
}
