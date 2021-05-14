package http

import (
	"encoding/json"
	"firstgo/frame/context"
	"firstgo/frame/exception"
	"firstgo/frame/proxy"
	"firstgo/frame/vo"
	"firstgo/util"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type RestAnnotationSetting struct {

	//对应path路径 以/开头
	HttpPath string

	//http method get,post,put,delete,*
	HttpMethod string

	//方法对应的request参数名
	MethodParamter string

	//默认的渲染类型 json html 默认是json
	MethodRender string
}

type ControllerVar struct {
	Target             proxy.ProxyTarger
	PrefixPath         string
	AbsoluteMethodPath map[string]*proxy.ProxyMethod
}

type DispatchServlet struct {
	ContextPath string
}

func (d *DispatchServlet) Dispatch(local *context.LocalStack, request *http.Request, response http.ResponseWriter) {
	controller := GetCurrentControllerInvoker(local)

	var proxyMethod *proxy.ProxyMethod
	var methodRequestSetting *RestAnnotationSetting

	defer func() {
		if err := recover(); err != nil {
			s := PrintStackTrace(err)
			fmt.Println(s)
			if methodRequestSetting.MethodRender == "" || methodRequestSetting.MethodRender == "json" {
				d.renderExceptionJson(response, request, err)
			}

		}
	}()

	// 去除?之后的
	url := util.ConfigUtil.ClearHttpPath(request.URL.Path)
	url = util.ConfigUtil.RemovePrefix(url, controller.PrefixPath)

	httpMethod := strings.ToLower(request.Method)
	mk := fmt.Sprintf("%s-%s", httpMethod, url)

	if _, ok := controller.AbsoluteMethodPath[mk]; ok {
		proxyMethod = controller.AbsoluteMethodPath[mk]
	} else {
		mk = fmt.Sprintf("%s-%s", "*", url)
		proxyMethod = controller.AbsoluteMethodPath[mk]
	}

	// proxyMethod== nil 404 TODO

	methodRequestSetting = GetRequestAnnotationSetting(proxyMethod.Annotations)

	methodInvoker := reflect.ValueOf(controller.Target).MethodByName(proxyMethod.Name)
	paramlen := methodInvoker.Type().NumIn()

	var result []reflect.Value
	if paramlen == 0 {
		result = methodInvoker.Call([]reflect.Value{})
	} else {
		var paramter []string
		if methodRequestSetting.MethodParamter != "" {
			paramter = strings.Split(methodRequestSetting.MethodParamter, ",")
		}

		param := make([]reflect.Value, paramlen)
		for i := 0; i < paramlen; i++ {
			pt := methodInvoker.Type().In(i)
			switch pt.Kind() {
			case reflect.Ptr:
				if pt.Elem() == reflect.TypeOf(*request) {
					param[i] = reflect.ValueOf(request)
				} else if pt.Elem() == reflect.TypeOf(*local) {
					param[i] = reflect.ValueOf(local)
				} else {
					nt := reflect.New(pt.Elem())
					ntpr := nt.Interface()

					body, err := ioutil.ReadAll(request.Body)
					if err != nil {
						panic(fmt.Errorf("read requestbody error"))
					}
					if len(body) == 0 {
						panic(fmt.Errorf("read requestbody empty"))
					}
					json.Unmarshal(body, ntpr)
					param[i] = reflect.ValueOf(ntpr)
				}
			case reflect.Interface:
				if reflect.TypeOf(response).Implements(pt) {
					param[i] = reflect.ValueOf(response)
				}
			case reflect.String:
				var pk string = paramter[i]
				var pv string = request.FormValue(pk)
				if pv == "" {
					pv = request.URL.Query().Get(pk)
				}
				param[i] = reflect.ValueOf(pv)
			case reflect.Int:
				var pk string = paramter[i]
				var pv string = request.FormValue(pk)
				if pv == "" {
					pv = request.URL.Query().Get(pk)
				}
				var pvi int = 0
				if pv != "" {
					var err error
					pvi, err = strconv.Atoi(pv)
					if err != nil {
						panic(fmt.Errorf("string2int error"))
					}
				} else {
					panic(fmt.Errorf("paramter %s get error", pk))
				}
				param[i] = reflect.ValueOf(pvi)
			case reflect.Struct:
				panic(fmt.Errorf("struct only ptr"))
			}
		}
		result = methodInvoker.Call(param)
	}
	if len(result) == 1 && methodRequestSetting.MethodRender == "" || methodRequestSetting.MethodRender == "json" {
		d.renderJson(response, request, result[0].Interface())
	}

}

func (d *DispatchServlet) renderJson(response http.ResponseWriter, request *http.Request, result interface{}) {
	response.Header().Add("Content-Type", "application/json;charset=UTF-8")
	a, _ := json.Marshal(result)
	response.Write(a)
}
func (d *DispatchServlet) renderExceptionJson(response http.ResponseWriter, request *http.Request, throwable interface{}) {

	var errJson *vo.JsonResult
	switch throwable.(type) {
	case *exception.FrameException:
		value, _ := throwable.(*exception.FrameException)
		errJson = util.JsonUtil.BuildJsonFailure(value.Code, value.Message, nil)
	default:
		tip := fmt.Sprintln(throwable)
		errJson = util.JsonUtil.BuildJsonFailure1(tip, nil)
	}

	response.Header().Add("Content-Type", "application/json;charset=UTF-8")
	a, _ := json.Marshal(errJson)
	response.Write(a)
}

var dispatchServlet DispatchServlet = DispatchServlet{
	ContextPath: util.ConfigUtil.Get("contextPath", "/api"),
}

func GetDispatchServlet() *DispatchServlet {
	return &dispatchServlet
}

// AddControllerProxyTarget 思路是根据path前缀匹配到controller，在根据path和method去匹配controller具体的method
func AddControllerProxyTarget(target1 proxy.ProxyTarger) {
	proxy.AddClassProxy(target1)

	var methodRef = make(map[string]*proxy.ProxyMethod)
	for _, method := range target1.ProxyTarget().Methods {
		methodRestSetting := GetRequestAnnotationSetting(method.Annotations)
		if methodRestSetting == nil {
			continue
		}

		//http method
		var hm = strings.ToLower(methodRestSetting.HttpMethod)
		if hm == "" {
			hm = "*"
		}

		var hp = methodRestSetting.HttpPath
		if hp == "/" {
			hp = ""
		}

		mkey := fmt.Sprintf("%s-%s", hm, hp)
		methodRef[mkey] = method
	}
	var prefix = GetControllerPathPrefix(&dispatchServlet, target1)
	invoker := &ControllerVar{
		Target:             target1,
		PrefixPath:         prefix,
		AbsoluteMethodPath: methodRef,
	}

	f := func(invoker1 *ControllerVar) func(http.ResponseWriter, *http.Request) {
		return func(response http.ResponseWriter, request *http.Request) {
			local := context.NewLocalStack()

			SetCurrentHttpRequest(local, request)
			SetCurrentHttpResponse(local, response)

			defer local.Destroy()

			SetCurrentControllerInvoker(local, invoker1)

			GetDefaultFilterChain().DoFilter(local, request, response)
		}
	}(invoker)
	http.HandleFunc(prefix+"/", f) //前缀匹配
	http.HandleFunc(prefix, f)     //绝对匹配
}
