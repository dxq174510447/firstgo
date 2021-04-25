package httpfilter

import (
	"firstgo/frame/context"
	http2 "firstgo/frame/http"
	"firstgo/povo/po"
	"firstgo/util"
	"fmt"
	"net/http"
)

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

var bindUserFilter BindUserFilter = BindUserFilter{}

func GetBindUserFilter() *BindUserFilter {
	return &bindUserFilter
}

func getUsersByToken(token string) *po.Users {
	return &po.Users{
		Id:   1,
		Name: "44444",
	}
}

func init() {
	http2.AddFilter(http2.Filter(&bindUserFilter))
}
