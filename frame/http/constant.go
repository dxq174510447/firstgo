package http

import (
	"bytes"
	"firstgo/frame/context"
	"firstgo/frame/proxy"
	"fmt"
	"runtime"
)

const (
	AnnotationRestController = "AnnotationRestController_"

	AnnotationValueRestKey = "AnnotationValueRestKey_"

	FilterIndexWaitToExecute = "FilterIndexWaitToExecute_"

	CurrentControllerInvoker = "CurrentControllerInvoker_"
)

func SetCurrentControllerInvoker(local *context.LocalStack, invoker1 *ControllerVar) {
	local.Set(CurrentControllerInvoker, invoker1)
}
func GetCurrentControllerInvoker(local *context.LocalStack) *ControllerVar {
	invoker := local.Get(CurrentControllerInvoker)
	return invoker.(*ControllerVar)
}

func SetCurrentFilterIndex(local *context.LocalStack, index int) {
	local.Set(FilterIndexWaitToExecute, index)
}

func GetCurrentFilterIndex(local *context.LocalStack) int {
	index := local.Get(FilterIndexWaitToExecute)
	if index == nil {
		return 0
	}
	return index.(int)
}

func GetRequestAnnotationSetting(annotations []*proxy.AnnotationClass) *RestAnnotationSetting {
	for _, annotation := range annotations {
		if annotation.Name == AnnotationRestController {
			r, _ := annotation.Value[AnnotationValueRestKey]
			return r.(*RestAnnotationSetting)
		}
	}
	return nil
}

func NewRestAnnotation(httpPath string,
	httpMethod string,
	methodParamter string,
	methodRender string) *proxy.AnnotationClass {
	return &proxy.AnnotationClass{
		Name: AnnotationRestController,
		Value: map[string]interface{}{
			AnnotationValueRestKey: &RestAnnotationSetting{
				HttpPath:       httpPath,
				HttpMethod:     httpMethod,
				MethodParamter: methodParamter,
				MethodRender:   methodRender,
			},
		},
	}
}

func PrintStackTrace(err interface{}) string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%v\n", err)
	for i := 1; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
	}
	return buf.String()
}

func GetControllerPathPrefix(dispatchServlet *DispatchServlet, target proxy.ProxyTarger) string {
	//context-path
	var sp string = dispatchServlet.ContextPath
	if sp == "/" {
		sp = ""
	}

	//controller-path
	var classRestSetting *RestAnnotationSetting = GetRequestAnnotationSetting(target.ProxyTarget().Annotations)
	var cp string = classRestSetting.HttpPath
	if cp == "/" {
		cp = ""
	}
	return fmt.Sprintf("%s%s", sp, cp)
}
