package httpfilter

import (
	"firstgo/src/main/golang/com/firstgo/povo/po"
	"firstgo/src/main/golang/com/firstgo/util"
	"github.com/dxq174510447/goframe/lib/frame/application"
	context "github.com/dxq174510447/goframe/lib/frame/context"
	http2 "github.com/dxq174510447/goframe/lib/frame/http"
	"github.com/dxq174510447/goframe/lib/frame/proxy/proxyclass"
	"net/http"
)

// BindUserFilter test localstack
type BindUserFilter struct {
	Logger application.AppLoger `FrameAutowired:""`
}

func (b *BindUserFilter) DoFilter(local *context.LocalStack,
	request *http.Request, response http.ResponseWriter, chain http2.FilterChain) {

	b.Logger.Debug(local, "%s", "BindUserFilter begin")

	token := request.Header.Get("token")
	if token != "" {
		current := getUsersByToken(token)
		if current != nil {
			util.WebUtil.SetThreadUser(local, current)
		}
	}
	chain.DoFilter(local, request, response)
}

func (b *BindUserFilter) Order() int {
	return 10
}

func (b *BindUserFilter) ProxyTarget() *proxyclass.ProxyClass {
	return nil
}

var bindUserFilter BindUserFilter = BindUserFilter{}

func getUsersByToken(token string) *po.Users {
	return &po.Users{
		Id:   1,
		Name: "44444",
	}
}

func init() {
	application.AddProxyInstance("", proxyclass.ProxyTarger(&bindUserFilter))
}
