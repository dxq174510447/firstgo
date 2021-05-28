package httpfilter

import (
	"firstgo/povo/po"
	"firstgo/util"
	"fmt"
	"github.com/dxq174510447/goframe/lib/frame/application"
	context "github.com/dxq174510447/goframe/lib/frame/context"
	http2 "github.com/dxq174510447/goframe/lib/frame/http"
	"github.com/dxq174510447/goframe/lib/frame/proxy"
	"net/http"
)

// BindUserFilter test localstack
type BindUserFilter struct {
}

func (b *BindUserFilter) DoFilter(local *context.LocalStack,
	request *http.Request, response http.ResponseWriter, chain http2.FilterChain) {
	fmt.Println("BindUserFilter begin")
	defer func() {
		fmt.Println("BindUserFilter end")
	}()

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

func (b *BindUserFilter) ProxyTarget() *proxy.ProxyClass {
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
	application.AddProxyInstance("", proxy.ProxyTarger(&bindUserFilter))
}
