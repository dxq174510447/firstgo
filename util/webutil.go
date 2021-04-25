package util

import (
	"firstgo/frame/context"
	"firstgo/povo/po"
)

type webUtil struct {
}

func (w *webUtil) SetThreadUser(local *context.LocalStack, user *po.Users) {
	local.Set("SetThreadUser_", user)
}
func (w *webUtil) GetThreadUser(local *context.LocalStack) *po.Users {
	v := local.Get("SetThreadUser_")
	if v == nil {
		return nil
	}
	return v.(*po.Users)
}

var WebUtil webUtil = webUtil{}
