package kllib

import (
	"github.com/dxq174510447/goframe/lib/frame/context"
	"github.com/dxq174510447/goframe/lib/frame/http"
	"github.com/dxq174510447/goframe/lib/frame/proxy"
	//"fmt"
	//"github.com/urfave/negroni"
	//"klook.libs/comm"
	//"klook.libs/statsd"
	"reflect"
	//"time"
)

// KlMonitorProxyFilter aop 拦截controller 监控方法访问
type KlMonitorProxyFilter struct {
	Next proxy.ProxyFilter
}

func (d *KlMonitorProxyFilter) Execute(context *context.LocalStack,
	classInfo *proxy.ProxyClass,
	methodInfo *proxy.ProxyMethod,
	invoker *reflect.Value,
	arg []reflect.Value) []reflect.Value {
	return nil
	//fmt.Printf("KlMonitorProxyFilter begin \n")
	//request := http.GetCurrentHttpRequest(context)
	//response := http.GetCurrentHttpResponse(context)
	//startTime := time.Now()
	//funcName := fmt.Sprintf("%s-%s", classInfo.Name, methodInfo.Name)
	//httpMethod := request.Method
	//klheader := GetCurrentKlHeader(context)
	//
	//defer func() {
	//	fmt.Printf("KlMonitorProxyFilter end \n")
	//
	//	duration := time.Since(startTime)
	//
	//	statsd.HttpHandlerStatsCollector.Collect(negroni.NewResponseWriter(response), request, funcName, klheader.RequestID,
	//		startTime, duration, comm.GetPath(request))
	//
	//	if err := recover(); err != nil {
	//		panic(err)
	//	}
	//}()
	//statsd.Metrics.Inc(httpMethod, funcName, "")
	//
	//result := d.Next.Execute(context, classInfo, methodInfo, invoker, arg)
	//
	//statsd.Metrics.Observe(httpMethod, funcName, "", startTime)
	//
	//return result
}

func (d *KlMonitorProxyFilter) SetNext(next proxy.ProxyFilter) {
	d.Next = next
}

func (d *KlMonitorProxyFilter) Order() int {
	return 1
}

var klMonitorProxyFilter KlMonitorProxyFilter = KlMonitorProxyFilter{}

type KlMonitorProxyFilterFactory struct {
}

func (d *KlMonitorProxyFilterFactory) GetInstance(m map[string]interface{}) proxy.ProxyFilter {
	return proxy.ProxyFilter(&klMonitorProxyFilter)
}

func (d *KlMonitorProxyFilterFactory) AnnotationMatch() []string {
	return []string{http.AnnotationRestController, http.AnnotationController}
}

var klMonitorProxyFilterFactory KlMonitorProxyFilterFactory = KlMonitorProxyFilterFactory{}

func init() {
	//proxy.AddProxyFilterFactory(proxy.ProxyFilterFactory(&klMonitorProxyFilterFactory))
}
