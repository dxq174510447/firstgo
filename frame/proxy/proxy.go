package proxy

import (
	"firstgo/frame/context"
	"fmt"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"sync"
)

type methodInvoke struct {
	target   interface{}
	clazz    *ProxyClass
	method   *ProxyMethod
	invoker  *reflect.Value
	filters  []ProxyFilter
	initLock sync.Mutex
}

func (m *methodInvoke) initFilter() {
	if m.filters != nil && len(m.filters) > 0 {
		return
	}
	m.initLock.Lock()
	defer m.initLock.Unlock()

	if m.filters != nil && len(m.filters) > 0 {
		return
	}

	fmt.Println("初始化")
	executor := defaultProxyFilterFactory.GetInstance(nil)

	var fs []ProxyFilter
	var hasCreate map[string]string = make(map[string]string)

	// method level
	if m.method != nil && len(m.method.Annotations) > 0 {
		for _, annotation := range m.method.Annotations {
			if factorys, ok := methodFilter[annotation.Name]; ok {
				for _, factory := range factorys {
					instance := factory.GetInstance(annotation.Value)
					if instance != nil {
						fs = append(fs, instance)
						hasCreate[annotation.Name] = "1"
					}
				}
			}
		}
	}

	if m.clazz != nil && len(m.clazz.Annotations) > 0 {
		for _, annotation := range m.clazz.Annotations {
			if _, ok := hasCreate[annotation.Name]; ok {
				continue
			}
			if factorys, ok := methodFilter[annotation.Name]; ok {
				for _, factory := range factorys {
					instance := factory.GetInstance(annotation.Value)
					if instance != nil {
						fs = append(fs, instance)
					}
				}
			}
		}
	}

	if len(fs) > 0 {
		if len(fs) > 1 {
			sort.Slice(fs, func(i, j int) bool {
				return fs[i].Order() < fs[j].Order()
			})
		}
		m.filters = append(fs, executor)
		l := len(m.filters)
		for i, f := range m.filters {
			if i == (l - 1) {
				break
			}
			f.SetNext(m.filters[i+1])
		}
	} else {
		m.filters = []ProxyFilter{executor}
	}
}
func (m *methodInvoke) invoke(context *context.LocalStack, args []reflect.Value) []reflect.Value {
	m.initFilter()
	for _, r := range m.filters {
		fmt.Println(reflect.ValueOf(r).Elem().Type().Name())
	}
	return m.filters[0].Execute(context,
		m.clazz,
		m.method,
		m.invoker,
		args,
	)
}

func newMethodInvoke(
	target interface{},
	clazz *ProxyClass,
	method *ProxyMethod,
	invoker *reflect.Value) *methodInvoke {
	return &methodInvoke{
		target:  target,
		clazz:   clazz,
		method:  method,
		invoker: invoker,
	}
}

// classProxy 好像没地方用到 key是全路径 GetClassName
var classProxy map[string]*ProxyClass = make(map[string]*ProxyClass)

// ProxyFilterFactory key annotation 名字 可以生产filter的factory实例
var methodFilter map[string][]ProxyFilterFactory = make(map[string][]ProxyFilterFactory)

var defaultExecuteFilter ProxyFilterFactory

// AddDefaultInvokerFilterFactory 默认filter执行器
func AddDefaultInvokerFilterFactory(target ProxyFilterFactory) {
	defaultExecuteFilter = target
}

// AddProxyFilterFactory 添加拦截器
func AddProxyFilterFactory(target ProxyFilterFactory) {

	match := target.AnnotationMatch()
	if len(match) == 0 {
		return
	}

	for _, annotation := range match {
		if v, ok := methodFilter[annotation]; ok {
			methodFilter[annotation] = append(v, target)
		} else {
			methodFilter[annotation] = []ProxyFilterFactory{target}
		}
	}

}

func AddClassProxy(target ProxyTarger) {
	clazz := target.ProxyTarget()
	clazzName := GetClassName(target)

	//添加到映射中
	clazz.Target = target
	clazz.Name = clazzName
	classProxy[clazzName] = clazz

	//获取对象方法设置
	methodRef := make(map[string]*ProxyMethod)
	if len(clazz.Methods) != 0 {
		for _, md := range clazz.Methods {
			methodRef[md.Name] = md
		}
	}

	//解析字段方法 包裹一层
	rv := reflect.ValueOf(target)
	rt := rv.Elem().Type()
	if m1 := rt.NumField(); m1 > 0 {
		for i := 0; i < m1; i++ {
			field := rt.Field(i)
			if field.Type.Kind() == reflect.Func {
				call := rv.Elem().FieldByName(field.Name)
				oldCall := reflect.ValueOf(call.Interface())

				methodName := strings.ReplaceAll(field.Name, "_", "")
				methodSetting, ok := methodRef[methodName]
				if !ok {
					methodSetting = &ProxyMethod{Name: methodName}
				}

				invoker := newMethodInvoke(target, clazz, methodSetting, &oldCall)

				proxyCall := func(command *methodInvoke) reflect.Value {
					newCall := reflect.MakeFunc(rt.Field(i).Type, func(in []reflect.Value) []reflect.Value {
						fmt.Printf(" %s %s agent begin \n", invoker.clazz.Name, invoker.method.Name)
						defer fmt.Printf(" %s %s agent end \n", invoker.clazz.Name, invoker.method.Name)
						return command.invoke(in[0].Interface().(*context.LocalStack), in)
					})
					return newCall
				}(invoker)
				call.Set(proxyCall)
			}
		}
	}
}

//GetClassName 用来获取struct的全路径 传递指针
func GetClassName(target interface{}) string {
	t := reflect.ValueOf(target).Elem().Type()
	return fmt.Sprintf("%s/%s", t.PkgPath(), t.Name())
}

func NewSingleAnnotation(annotationName string, value map[string]interface{}) []*AnnotationClass {
	return []*AnnotationClass{
		&AnnotationClass{
			Name:  annotationName,
			Value: value,
		},
	}
}

func getTargetValue(target interface{}, name string) interface{} {
	v := reflect.ValueOf(target)
	switch v.Kind() {
	case reflect.String, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64:
		return target
	case reflect.Map:
		m := target.(map[string]interface{})
		return m[name]
	case reflect.Ptr:
		if v.Elem().Kind() == reflect.Struct {
			return v.Elem().FieldByName(name).Interface()
		}
	}
	panic(fmt.Sprintf("%s找不到对应属性", name))
}

// GetVariableValue target 可能map接口 基础类型 指针结构体类型
func GetVariableValue(target interface{}, name string) interface{} {

	keys := strings.Split(name, ".")

	l := len(keys)
	if l == 1 {
		return getTargetValue(target, name)
	} else {
		nt := target
		for i := 0; i < l; i++ {
			nt = getTargetValue(nt, keys[i])
			//中间值为nil 就panic
			if i < (l-1) && reflect.ValueOf(nt).IsZero() {
				panic(fmt.Sprintf("sql %s is nil value", name))
			}
		}
		return nt
	}

}

func GetStructField(rtType reflect.Type) map[string]reflect.StructField {
	ref := make(map[string]reflect.StructField)
	switch rtType.Kind() {
	case reflect.Slice:
		if rtType.Elem().Kind() == reflect.Struct {
			n := rtType.Elem().NumField()
			for i := 0; i < n; i++ {
				sf := rtType.Elem().Field(i)
				ref[sf.Name] = sf
			}
			return ref
		} else if rtType.Elem().Kind() == reflect.Ptr && rtType.Elem().Elem().Kind() == reflect.Struct {
			n := rtType.Elem().Elem().NumField()
			for i := 0; i < n; i++ {
				sf := rtType.Elem().Elem().Field(i)
				ref[sf.Name] = sf
			}
			return ref
		} else {
			return nil
		}
	case reflect.Ptr:
		if rtType.Elem().Kind() != reflect.Struct {
			return nil
		} else {
			n := rtType.Elem().NumField()
			for i := 0; i < n; i++ {
				sf := rtType.Elem().Field(i)
				ref[sf.Name] = sf
			}
			return ref
		}
	default:
		return nil
	}

}

func GetTypeDefaultValue(rtType reflect.Type) *reflect.Value {
	var result reflect.Value
	switch rtType.Kind() {
	case reflect.String:
		result = reflect.ValueOf("")
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		result = reflect.ValueOf(0)
	case reflect.Float32, reflect.Float64:
		result = reflect.ValueOf(0.0)
	case reflect.Map:
		v := reflect.MakeMap(rtType)
		result = v
	case reflect.Slice:
		result = reflect.MakeSlice(rtType, 0, 0)
	case reflect.Ptr:
		result = reflect.New(rtType)
	case reflect.Struct:
		result = reflect.New(rtType)
	default:
		panic(fmt.Sprintf("%s找不到对应默认值", rtType.String()))
	}
	return &result
}

var matchAllCap = regexp.MustCompile(`[^A-Za-z0-9]+`)

func GetCamelCaseName(str string) string {
	st := matchAllCap.Split(str, -1)
	for k, s := range st {
		st[k] = strings.Title(strings.ToLower(s))
	}
	return strings.Join(st, "")
}
