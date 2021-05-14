package kllib

import (
	"firstgo/frame/context"
	"firstgo/service/kllib/vo"
)

// 常量
const (
	KlRequestHeader = "KlRequestHeader_"
	KlRequestIdKey  = "X-Klook-Request-Id"
)

func SetCurrentKlHeader(local *context.LocalStack, klheader *vo.KlHeaderVo) {
	local.Set(KlRequestHeader, klheader)
}
func GetCurrentKlHeader(local *context.LocalStack) *vo.KlHeaderVo {
	invoker := local.Get(KlRequestHeader)
	return invoker.(*vo.KlHeaderVo)
}
