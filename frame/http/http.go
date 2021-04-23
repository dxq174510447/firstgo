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

func NewRestAnnotation(httpPath string,
	httpMethod string,
	methodParamter string,
	methodRender string) []*proxy.AnnotationClass {
	return []*proxy.AnnotationClass{
		&proxy.AnnotationClass{
			Name: AnnotationRestController,
			Value: map[string]interface{}{
				AnnotationValueRestKey: &RestAnnotationSetting{
					HttpPath:       httpPath,
					HttpMethod:     httpMethod,
					MethodParamter: methodParamter,
					MethodRender:   methodRender,
				},
			},
		},
	}
}

type dispatchServlet struct {
	contextPath string
}

func (d *dispatchServlet) renderJson(response http.ResponseWriter, request *http.Request, result interface{}) {
	response.Header().Add("Content-Type", "application/json;charset=UTF-8")
	a, _ := json.Marshal(result)
	response.Write(a)
}
func (d *dispatchServlet) renderExceptionJson(response http.ResponseWriter, request *http.Request, throwable interface{}) {

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

// AddRequestMapping 思路是根据path前缀匹配到controller，在根据path和method去匹配controller具体的method
func (d *dispatchServlet) AddRequestMapping(mapping *RequestController) {
	var sp string = d.contextPath
	if sp == "/" {
		sp = ""
	}
	var cp = mapping.HttpPath
	if cp == "/" {
		cp = ""
	}
	var prefix = fmt.Sprintf("%s%s", sp, cp)
	f := func() func(http.ResponseWriter, *http.Request) {
		var pf string = prefix
		var target interface{} = mapping.Target
		var methodRef = make(map[string]RequestMethod)
		for _, method := range mapping.Methods {
			var hm = strings.ToLower(method.HttpMethod)
			var hp = method.HttpPath
			if hp == "/" {
				hp = ""
			}
			if hm == "" {
				hm = "*"
			}
			mk := fmt.Sprintf("%s-%s", hm, hp)
			methodRef[mk] = method
		}

		return func(response http.ResponseWriter, request *http.Request) {
			//fmt.Println(request.URL.Path)
			var requestMethod RequestMethod
			local := context.NewLocalStack()
			defer func() {
				local.Pop()
				local = nil

				if err := recover(); err != nil {

					if requestMethod.MethodRender == "" || requestMethod.MethodRender == "json" {
						d.renderExceptionJson(response, request, err)
					}

				}
			}()

			url := util.ConfigUtil.ClearHttpPath(request.URL.Path)
			url = util.ConfigUtil.RemovePrefix(url, pf)
			httpMethod := strings.ToLower(request.Method)
			mk := fmt.Sprintf("%s-%s", httpMethod, url)
			//		fmt.Println(mk)

			if _, ok := methodRef[mk]; ok {
				requestMethod = methodRef[mk]
			} else {
				mk = fmt.Sprintf("%s-%s", "*", url)
				requestMethod = methodRef[mk]
			}
			methodType := reflect.ValueOf(target).MethodByName(requestMethod.MethodName)
			paramlen := methodType.Type().NumIn()

			var result []reflect.Value
			if paramlen == 0 {
				result = reflect.ValueOf(target).MethodByName(requestMethod.MethodName).Call([]reflect.Value{})
			} else {
				var paramter []string
				if requestMethod.MethodParamter != "" {
					paramter = strings.Split(requestMethod.MethodParamter, ",")
				}
				param := make([]reflect.Value, paramlen)
				for i := 0; i < paramlen; i++ {
					pt := methodType.Type().In(i)
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
				result = reflect.ValueOf(target).MethodByName(requestMethod.MethodName).Call(param)
			}
			if len(result) == 1 && requestMethod.MethodRender == "" || requestMethod.MethodRender == "json" {
				d.renderJson(response, request, result[0].Interface())
			}
		}
	}()
	http.HandleFunc(prefix+"/", f)
	http.HandleFunc(prefix, f)
}

var DispatchServlet dispatchServlet = dispatchServlet{contextPath: util.ConfigUtil.Get("contextPath", "/api")}
