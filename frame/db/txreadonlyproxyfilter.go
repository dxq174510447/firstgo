package db

import (
	"firstgo/frame/context"
	"firstgo/frame/proxy"
	"fmt"
	"reflect"
)

type TxReadOnlyProxyFilter struct {
	Next proxy.ProxyFilter
}

func (d *TxReadOnlyProxyFilter) Execute(context *context.LocalStack,
	classInfo *proxy.ProxyClass,
	methodInfo *proxy.ProxyMethod,
	invoker *reflect.Value,
	arg []reflect.Value) []reflect.Value {
	fmt.Println("TxReadOnlyProxyFilter begin")

	defer fmt.Println("TxReadOnlyProxyFilter end")
	return d.Next.Execute(context, classInfo, methodInfo, invoker, arg)
}

func (d *TxReadOnlyProxyFilter) SetNext(next proxy.ProxyFilter) {
	d.Next = next
}

func (d *TxReadOnlyProxyFilter) Order() int {
	return 3
}

var txReadOnlyProxyFilter TxReadOnlyProxyFilter = TxReadOnlyProxyFilter{}

type TxReadOnlyProxyFilterFactory struct {
}

func (d *TxReadOnlyProxyFilterFactory) GetInstance(m map[string]interface{}) proxy.ProxyFilter {
	return proxy.ProxyFilter(&txReadOnlyProxyFilter)
}

func (d *TxReadOnlyProxyFilterFactory) AnnotationMatch() []string {
	return []string{TransactionReadOnly}
}

var txReadOnlyProxyFilterFactory TxReadOnlyProxyFilterFactory = TxReadOnlyProxyFilterFactory{}

func init() {
	proxy.AddProxyFilterFactory(proxy.ProxyFilterFactory(&txReadOnlyProxyFilterFactory))
}
