package dbcore

import (
	"bytes"
	"encoding/xml"
	"firstgo/frame/context"
	"firstgo/frame/proxy"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"text/template"
)

type SqlProviderConfig struct {
	Param string
}

type MapperElementXml struct {
	Id      string `xml:"id,attr"`
	Sql     string `xml:",innerxml"`
	SqlType string
	Tpl     *template.Template
}

type MapperXml struct {
	Sql []*MapperElementXml `xml:"sql"`

	SelectSql []*MapperElementXml `xml:"select"`

	UpdateSql []*MapperElementXml `xml:"update"`

	DeleteSql []*MapperElementXml `xml:"delete"`

	InsertSql []*MapperElementXml `xml:"insert"`
}

type MapperFactory struct {
	importTagRegexp *regexp.Regexp
}

func (m *MapperFactory) ReplaceImportTag(sql string, refs map[string]*MapperElementXml) string {
	ns := m.importTagRegexp.ReplaceAllStringFunc(sql, func(str string) string {
		str1 := m.importTagRegexp.FindStringSubmatch(str)
		if s, ok := refs[str1[1]]; ok {
			return s.Sql
		}
		return ""
	})
	return ns
}

func (m *MapperFactory) ParseXml(target proxy.ProxyTarger, content string) map[string]*MapperElementXml {
	mapper := &MapperXml{}
	err := xml.Unmarshal([]byte(content), mapper)

	if err != nil {
		panic(err)
	}
	refs := make(map[string]*MapperElementXml)

	if len(mapper.Sql) > 0 {
		for _, ele := range mapper.Sql {
			ele.SqlType = SqlTypeSql
			refs[ele.Id] = ele
		}
	}

	if len(mapper.UpdateSql) > 0 {
		for _, ele := range mapper.UpdateSql {
			ele.SqlType = SqlTypeUpdate
			ele.Sql = m.ReplaceImportTag(ele.Sql, refs)
			ele.Tpl = template.Must(template.New(fmt.Sprintf("%s-%s", proxy.GetClassName(target), ele.Id)).Parse(ele.Sql))
			refs[ele.Id] = ele
		}
	}

	if len(mapper.InsertSql) > 0 {
		for _, ele := range mapper.InsertSql {
			ele.SqlType = SqlTypeInsert
			ele.Sql = m.ReplaceImportTag(ele.Sql, refs)
			ele.Tpl = template.Must(template.New(fmt.Sprintf("%s-%s", proxy.GetClassName(target), ele.Id)).Parse(ele.Sql))
			refs[ele.Id] = ele
		}
	}

	if len(mapper.SelectSql) > 0 {
		for _, ele := range mapper.SelectSql {
			ele.SqlType = SqlTypeSelect
			ele.Sql = m.ReplaceImportTag(ele.Sql, refs)
			ele.Tpl = template.Must(template.New(fmt.Sprintf("%s-%s", proxy.GetClassName(target), ele.Id)).Parse(ele.Sql))
			refs[ele.Id] = ele
		}
	}

	if len(mapper.DeleteSql) > 0 {
		for _, ele := range mapper.DeleteSql {
			ele.SqlType = SqlTypeDelete
			ele.Sql = m.ReplaceImportTag(ele.Sql, refs)
			ele.Tpl = template.Must(template.New(fmt.Sprintf("%s-%s", proxy.GetClassName(target), ele.Id)).Parse(ele.Sql))
			refs[ele.Id] = ele
		}
	}

	return refs
}

var mapperFactory MapperFactory = MapperFactory{
	importTagRegexp: regexp.MustCompile(`<include refid="(\S+)">\s*</include>`),
}

func GetMapperFactory() *MapperFactory {
	return &mapperFactory
}

type sqlInvoke struct {
	target         interface{}
	clazz          *proxy.ProxyClass
	method         *proxy.ProxyMethod
	mapper         map[string]*MapperElementXml
	providerConfig *SqlProviderConfig
}

func (s *sqlInvoke) invoke(context *context.LocalStack, args []reflect.Value) []reflect.Value {
	methodName := s.method.Name
	ele := s.mapper[methodName]
	switch ele.SqlType {
	case SqlTypeSelect:
		return s.invokeSelect(context, args, ele)
	case SqlTypeUpdate:
		return s.invokeUpdate(context, args, ele)
	case SqlTypeInsert:
		return s.invokeInsert(context, args, ele)
	case SqlTypeDelete:
		return s.invokeDelete(context, args, ele)
	}
	return nil
}

func (s *sqlInvoke) getSqlFromTpl(context *context.LocalStack, args []reflect.Value, sql *MapperElementXml) (string, error) {

	// 去除局部变量参数
	if len(args) <= 1 {
		return sql.Sql, nil
	}

	// 只有一个参数 结构体 基础类型 string
	var params []string
	if s.providerConfig != nil && s.providerConfig.Param != "" {
		params = strings.Split(s.providerConfig.Param, ",")
	}
	if len(args) == 2 {
		if len(params) >= 2 && params[1] != "_" && params[1] != "" {
			root := make(map[string]interface{})
			root[params[1]] = args[1].Interface()
			buf := &bytes.Buffer{}
			err := sql.Tpl.Execute(buf, root)
			if err != nil {
				return "", err
			}
			return buf.String(), nil
		} else {
			buf := &bytes.Buffer{}
			err := sql.Tpl.Execute(buf, args[1].Interface())
			if err != nil {
				return "", err
			}
			return buf.String(), nil
		}
	}
	// > 2
	root := make(map[string]interface{})
	for i := 1; i < len(params); i++ {
		if params[i] != "_" && params[i] != "" {
			root[params[i]] = args[i].Interface()
		}
	}
	buf := &bytes.Buffer{}
	err := sql.Tpl.Execute(buf, root)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

// 必须返回两个值  一个sql返回的 一个error
// 默认只返回3个类型 slice，单个结构体，int
func (s *sqlInvoke) invokeSelect(local *context.LocalStack, args []reflect.Value, sqlEle *MapperElementXml) []reflect.Value {

	var defaultValue reflect.Value = reflect.ValueOf(1)
	//var defaultError *DaoException = &DaoException{exception.FrameException{Code: 505,Message: "数据库操作错误"}}
	var nilError error

	errorFlag := GetErrorHandleFlag(local) //0 panic 1 return
	//con := GetDbConnection(local)
	//if con == nil {
	//	var errortip string = "上下文中找不到数据库链接"
	//	if errorFlag == 0 {
	//		panic(errortip)
	//	}else{
	//		defaultError.Message=errortip
	//		return []reflect.Value{defaultValue, reflect.ValueOf(defaultError)}
	//	}
	//}

	sql, err1 := s.getSqlFromTpl(local, args, sqlEle)
	if err1 != nil {
		if errorFlag == 0 {
			panic(err1)
		} else {
			return []reflect.Value{defaultValue, reflect.ValueOf(err1)}
		}
	}

	sqlParam, err2 := s.getArgumentsFromSql(local, args, sql)
	if err2 != nil {
		if errorFlag == 0 {
			panic(err2)
		} else {
			return []reflect.Value{defaultValue, reflect.ValueOf(err2)}
		}
	}
	fmt.Println(sql)
	fmt.Println(sqlParam)
	//stmt, err := con.Con.PrepareContext(con.Ctx, sql)
	//if err != nil {
	//	if errorFlag == 0 {
	//		panic(err)
	//	}else{
	//		return []reflect.Value{defaultValue, reflect.ValueOf(err)}
	//	}
	//}
	//defer stmt.Close()
	//
	//
	//result := stmt.QueryRow(sqlParam...)

	//data := vo.UsersVo{}
	//if err := result.Scan(&data.Id, &data.Name, &data.Status); err != nil {
	//	return nil
	//}
	//return &data
	//
	//return []reflect.Value{reflect.ValueOf(1), reflect.ValueOf(dberr)}
	return []reflect.Value{defaultValue, reflect.ValueOf(nilError)}
}

func (s *sqlInvoke) invokeUpdate(context *context.LocalStack, args []reflect.Value, sql *MapperElementXml) []reflect.Value {
	return nil
}

func (s *sqlInvoke) invokeDelete(context *context.LocalStack, args []reflect.Value, sql *MapperElementXml) []reflect.Value {
	return nil
}

func (s *sqlInvoke) invokeInsert(context *context.LocalStack, args []reflect.Value, sql *MapperElementXml) []reflect.Value {
	return nil
}

func (s *sqlInvoke) getArgumentsFromSql(local *context.LocalStack, args []reflect.Value, sql string) ([]interface{}, error) {
	// 去除局部变量参数
	if len(args) <= 1 {
		return nil, nil
	}

	// 只有一个参数 结构体 基础类型 string
	var params []string
	if s.providerConfig != nil && s.providerConfig.Param != "" {
		params = strings.Split(s.providerConfig.Param, ",")
	}
	var root interface{}
	if len(args) == 2 {
		if len(params) >= 2 && params[1] != "_" && params[1] != "" {
			root1 := make(map[string]interface{})
			root1[params[1]] = args[1].Interface()
			root = root1
		} else {
			root = args[1].Interface()
		}
	} else {
		root1 := make(map[string]interface{})
		for i := 1; i < len(params); i++ {
			if params[i] != "_" && params[i] != "" {
				root1[params[i]] = args[i].Interface()
			}
		}
		root = root1
	}

	variables := parseAndGetSqlVariables(sql)
	if len(variables) == 0 {
		return nil, nil
	}
	var result []interface{}
	for _, v := range variables {
		m := proxy.GetVariableValue(root, v)
		result = append(result, m)
	}
	return result, nil
}

func newSqlInvoke(
	target interface{},
	clazz *proxy.ProxyClass,
	method *proxy.ProxyMethod,
	mapper map[string]*MapperElementXml,
	providerConfig *SqlProviderConfig) *sqlInvoke {
	return &sqlInvoke{
		target:         target,
		clazz:          clazz,
		method:         method,
		mapper:         mapper,
		providerConfig: providerConfig,
	}
}

func AddMapperProxyTarget(target1 proxy.ProxyTarger, xml string) {

	//解析字段方法 包裹一层
	rv := reflect.ValueOf(target1)
	rt := rv.Elem().Type()

	xmlele := mapperFactory.ParseXml(target1, xml)

	methodRef := make(map[string]*proxy.ProxyMethod)
	if len(target1.ProxyTarget().Methods) != 0 {
		for _, md := range target1.ProxyTarget().Methods {
			methodRef[md.Name] = md
		}
	}

	if m1 := rt.NumField(); m1 > 0 {
		for i := 0; i < m1; i++ {
			field := rt.Field(i)
			if field.Type.Kind() == reflect.Func && rv.Elem().FieldByName(field.Name).IsNil() {
				call := rv.Elem().FieldByName(field.Name)

				methodName := strings.ReplaceAll(field.Name, "_", "")
				methodSetting, ok := methodRef[methodName]
				if !ok {
					methodSetting = &proxy.ProxyMethod{Name: methodName}
				}

				var providerConfig *SqlProviderConfig = nil
				if methodSetting != nil && len(methodSetting.Annotations) > 0 {
					for _, anno := range methodSetting.Annotations {
						if anno.Name == AnnotationSqlProviderConfig {
							if provider, f := anno.Value[AnnotationSqlProviderConfigValueKey]; f {
								providerConfig = provider.(*SqlProviderConfig)
							}
						}
					}
				}

				invoker := newSqlInvoke(target1, target1.ProxyTarget(), methodSetting, xmlele, providerConfig)
				proxyCall := func(command *sqlInvoke) reflect.Value {
					newCall := reflect.MakeFunc(field.Type, func(in []reflect.Value) []reflect.Value {
						return command.invoke(in[0].Interface().(*context.LocalStack), in)
					})
					return newCall
				}(invoker)
				call.Set(proxyCall)
			}
		}
	}

	proxy.AddClassProxy(target1)
}

func NewSqlProvierConfigAnnotation(param string) []*proxy.AnnotationClass {
	return []*proxy.AnnotationClass{
		&proxy.AnnotationClass{
			Name: AnnotationSqlProviderConfig,
			Value: map[string]interface{}{
				AnnotationSqlProviderConfigValueKey: &SqlProviderConfig{
					Param: param,
				},
			},
		},
	}
}

// parseAndGetSqlVariables #{ada} 获取ada
func parseAndGetSqlVariables(sql string) []string {
	reg := regexp.MustCompile(`(?m)#\{(\S+?)\}`)
	result := reg.FindAllStringSubmatch(sql, -1)
	if result != nil {
		var r1 []string
		for _, k := range result {
			r1 = append(r1, k[1])
		}
		return r1
	} else {
		return nil
	}
}
