package util

import (
	"firstgo/povo/po"
	"fmt"
	"github.com/dxq174510447/goframe/lib/frame/context"
	uuid "github.com/nu7hatch/gouuid"
	"time"
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
func (w *webUtil) GenKlRequestId() string {
	id, _ := uuid.NewV4()

	rid := id.String()
	if rid == "" || len(rid) < 8 { // 取毫秒
		return fmt.Sprintf("%d", time.Now().UTC().UnixNano()/1000)
	}
	return rid[0:7] // 取uuid的前8位
}

var WebUtil webUtil = webUtil{}
