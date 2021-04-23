package db

import (
	"firstgo/frame/context"
	"firstgo/frame/proxy"
	"fmt"
	"reflect"
)

type TxRequireProxyFilter struct {
	Next proxy.ProxyFilter
}

func (d *TxRequireProxyFilter) Execute(context *context.LocalStack,
	classInfo *proxy.ProxyClass,
	methodInfo *proxy.ProxyMethod,
	invoker *reflect.Value,
	arg []reflect.Value) []reflect.Value {
	fmt.Println("TxRequireProxyFilter begin")

	defer fmt.Println("TxRequireProxyFilter end")
	return d.Next.Execute(context, classInfo, methodInfo, invoker, arg)
}

func (d *TxRequireProxyFilter) SetNext(next proxy.ProxyFilter) {
	d.Next = next
}

func (d *TxRequireProxyFilter) Order() int {
	return 5
}

var txRequireProxyFilter TxRequireProxyFilter = TxRequireProxyFilter{}

type TxRequireProxyFilterFactory struct {
}

func (d *TxRequireProxyFilterFactory) GetInstance(m map[string]interface{}) proxy.ProxyFilter {
	return proxy.ProxyFilter(&txRequireProxyFilter)
}

func (d *TxRequireProxyFilterFactory) AnnotationMatch() []string {
	return []string{TransactionRequire}
}

var txRequireProxyFilterFactory TxRequireProxyFilterFactory = TxRequireProxyFilterFactory{}

func init() {
	proxy.AddProxyFilterFactory(proxy.ProxyFilterFactory(&txRequireProxyFilterFactory))
}
