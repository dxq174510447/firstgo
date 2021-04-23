package db

import (
	"firstgo/frame/context"
	"firstgo/frame/proxy"
	"fmt"
	"reflect"
)

type DaoConnectProxyFilter struct {
	Next proxy.ProxyFilter
}

func (d *DaoConnectProxyFilter) Execute(context *context.LocalStack,
	classInfo *proxy.ProxyClass,
	methodInfo *proxy.ProxyMethod,
	invoker *reflect.Value,
	arg []reflect.Value) []reflect.Value {
	fmt.Println("DaoConnectProxyFilter begin")

	defer fmt.Println("DaoConnectProxyFilter end")
	return d.Next.Execute(context, classInfo, methodInfo, invoker, arg)
}

func (d *DaoConnectProxyFilter) SetNext(next proxy.ProxyFilter) {
	d.Next = next
}

func (d *DaoConnectProxyFilter) Order() int {
	return 15
}

var daoConnectProxyFilter DaoConnectProxyFilter = DaoConnectProxyFilter{}

type DaoConnectProxyFilterFactory struct {
}

func (d *DaoConnectProxyFilterFactory) GetInstance(m map[string]interface{}) proxy.ProxyFilter {
	return proxy.ProxyFilter(&daoConnectProxyFilter)
}

func (d *DaoConnectProxyFilterFactory) AnnotationMatch() []string {
	return []string{proxy.AnnotationDao}
}

var daoConnectProxyFilterFactory DaoConnectProxyFilterFactory = DaoConnectProxyFilterFactory{}

func init() {
	proxy.AddProxyFilterFactory(proxy.ProxyFilterFactory(&daoConnectProxyFilterFactory))
}
