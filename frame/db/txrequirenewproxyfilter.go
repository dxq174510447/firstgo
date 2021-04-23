package db

import (
	"firstgo/frame/context"
	"firstgo/frame/proxy"
	"fmt"
	"reflect"
)

type TxRequireNewProxyFilter struct {
	Next proxy.ProxyFilter
}

func (d *TxRequireNewProxyFilter) Execute(context *context.LocalStack,
	classInfo *proxy.ProxyClass,
	methodInfo *proxy.ProxyMethod,
	invoker *reflect.Value,
	arg []reflect.Value) []reflect.Value {
	fmt.Println("TxRequireNewProxyFilter begin")

	defer fmt.Println("TxRequireNewProxyFilter end")
	return d.Next.Execute(context, classInfo, methodInfo, invoker, arg)
}

func (d *TxRequireNewProxyFilter) SetNext(next proxy.ProxyFilter) {
	d.Next = next
}

func (d *TxRequireNewProxyFilter) Order() int {
	return 4
}

var txRequireNewProxyFilter TxRequireNewProxyFilter = TxRequireNewProxyFilter{}

type TxRequireNewProxyFilterFactory struct {
}

func (d *TxRequireNewProxyFilterFactory) GetInstance(m map[string]interface{}) proxy.ProxyFilter {
	return proxy.ProxyFilter(&txRequireNewProxyFilter)
}

func (d *TxRequireNewProxyFilterFactory) AnnotationMatch() []string {
	return []string{TransactionRequireNew}
}

var txRequireNewProxyFilterFactory TxRequireNewProxyFilterFactory = TxRequireNewProxyFilterFactory{}

func init() {
	proxy.AddProxyFilterFactory(proxy.ProxyFilterFactory(&txRequireNewProxyFilterFactory))
}
