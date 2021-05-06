package dbcore

import (
	"bytes"
	"database/sql"
	"encoding/xml"
	"firstgo/frame/context"
	"firstgo/frame/exception"
	"firstgo/frame/proxy"
	"firstgo/povo/po"
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
	reg := regexp.MustCompile(`(?m)(^\s+|\s+$)`)
	ns = reg.ReplaceAllString(ns, " ")
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
	target             interface{}
	clazz              *proxy.ProxyClass
	method             *proxy.ProxyMethod
	mapper             map[string]*MapperElementXml
	providerConfig     *SqlProviderConfig
	returnSqlType      reflect.Type
	defaultReturnValue *reflect.Value
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

// 必须返回1-2个参数，其他一个必须是error并且放在最后一个返回值，至于sql返回有没有都行
// 默认只返回3种类型 slice，单个结构体，int float64 string
func (s *sqlInvoke) invokeSelect(local *context.LocalStack, args []reflect.Value, sqlEle *MapperElementXml) []reflect.Value {
	var nilError *DaoException

	errorFlag := GetErrorHandleFlag(local) //0 panic 1 return
	con := GetDbConnection(local)
	if con == nil {
		var errortip string = "上下文中找不到数据库链接"
		if errorFlag == 0 {
			panic(errortip)
		} else {
			var defaultError *DaoException = &DaoException{exception.FrameException{Code: 505, Message: errortip}}
			if s.returnSqlType != nil {
				return []reflect.Value{*s.defaultReturnValue, reflect.ValueOf(defaultError)}
			} else {
				return []reflect.Value{reflect.ValueOf(defaultError)}
			}
		}
	}

	sql, err1 := s.getSqlFromTpl(local, args, sqlEle)
	if err1 != nil {
		if errorFlag == 0 {
			panic(err1)
		} else {
			if s.returnSqlType != nil {
				return []reflect.Value{*s.defaultReturnValue, reflect.ValueOf(err1)}
			} else {
				return []reflect.Value{reflect.ValueOf(err1)}
			}
		}
	}

	sqlParam, newsql, err2 := s.getArgumentsFromSql(local, args, sql)
	if err2 != nil {
		if errorFlag == 0 {
			panic(err2)
		} else {
			if s.returnSqlType != nil {
				return []reflect.Value{*s.defaultReturnValue, reflect.ValueOf(err2)}
			} else {
				return []reflect.Value{reflect.ValueOf(err2)}
			}
		}
	}
	if newsql != "" {
		sql = newsql
	}
	fmt.Printf("Sql[%s]: %s \n", sqlEle.Id, sql)
	fmt.Printf("Paramters[%s]: %s \n", sqlEle.Id, sqlParam)
	stmt, err := con.Con.PrepareContext(con.Ctx, sql)
	if err != nil {
		if errorFlag == 0 {
			panic(err)
		} else {
			if s.returnSqlType != nil {
				return []reflect.Value{*s.defaultReturnValue, reflect.ValueOf(err)}
			} else {
				return []reflect.Value{reflect.ValueOf(err)}
			}
		}
	}
	defer stmt.Close()

	var queryResult *reflect.Value
	var queryError error
	switch s.returnSqlType.Kind() {
	case reflect.Slice:
		queryResult, queryError = s.selectList(stmt, sqlParam, errorFlag)
	case reflect.Ptr:
		//queryResult,queryError = s.selectRow(stmt,sqlParam,errorFlag)
		fmt.Println("1")
	default:
		fmt.Println("2")
	}

	if queryError != nil {
		if errorFlag == 0 {
			panic(queryError)
		} else {
			if s.returnSqlType != nil {
				return []reflect.Value{*s.defaultReturnValue, reflect.ValueOf(queryError)}
			} else {
				return []reflect.Value{reflect.ValueOf(queryError)}
			}
		}
	}

	//data := vo.UsersVo{}
	//if err := result.Scan(&data.Id, &data.Name, &data.Status); err != nil {
	//	return nil
	//}
	//return &data
	//
	//return []reflect.Value{reflect.ValueOf(1), reflect.ValueOf(dberr)}
	if s.returnSqlType != nil {
		return []reflect.Value{*queryResult, reflect.ValueOf(nilError)}
	} else {
		return []reflect.Value{reflect.ValueOf(nilError)}
	}

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

func (s *sqlInvoke) getArgumentsFromSql(local *context.LocalStack, args []reflect.Value, sql string) ([]interface{}, string, error) {
	// 去除局部变量参数
	if len(args) <= 1 {
		return nil, "", nil
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

	variables, nsql := parseAndGetSqlVariables(sql)
	if len(variables) == 0 {
		return nil, "", nil
	}
	var result []interface{}
	for _, v := range variables {
		m := proxy.GetVariableValue(root, v)
		result = append(result, m)
	}
	return result, nsql, nil
}

func (s *sqlInvoke) selectList(stmt *sql.Stmt, param []interface{}, errorFlag int) (*reflect.Value, error) {
	result, err1 := stmt.Query(param...)
	if err1 != nil {
		if errorFlag == 0 {
			panic(err1)
		} else {
			return nil, err1
		}
	}
	defer func() {
		if result != nil {
			result.Close() //可以关闭掉未scan连接一直占用
		}
	}()

	pageSize := 50
	queryCount := 0
	total := reflect.MakeSlice(s.returnSqlType, 0, 0)
	current := reflect.MakeSlice(s.returnSqlType, 0, pageSize)
	currentCount := 0
	for result.Next() {
		if queryCount != 0 && currentCount >= pageSize {
			total = reflect.AppendSlice(total, current)
			current = reflect.MakeSlice(s.returnSqlType, 0, pageSize)
			currentCount = 0
		}
		data := po.Users{}
		result.Scan(&data.Id, &data.Name, &data.Status) //不scan会导致连接不释放

		current = reflect.Append(current, reflect.ValueOf(&data))

		queryCount++
		currentCount++
	}
	fmt.Printf("queryCount %d \n", queryCount)
	if queryCount > 0 {
		total = reflect.AppendSlice(total, current)
		total = total.Slice(0, queryCount)
	}
	return &total, nil
}

func (s *sqlInvoke) selectRow(stmt *sql.Stmt, param []interface{}, errorFlag int) *reflect.Value {
	return nil
}

func newSqlInvoke(
	target interface{}, //对象
	clazz *proxy.ProxyClass, //代理信息
	method *proxy.ProxyMethod, //当前方法代理信息
	mapper map[string]*MapperElementXml, //对应sql节点
	returnSqlType reflect.Type, //返回的类型 不是error 如果没有就nil
	providerConfig *SqlProviderConfig,
	defaultReturnValue *reflect.Value, //默认返回类型值
) *sqlInvoke {
	return &sqlInvoke{
		target:             target,
		clazz:              clazz,
		method:             method,
		mapper:             mapper,
		providerConfig:     providerConfig,
		returnSqlType:      returnSqlType,
		defaultReturnValue: defaultReturnValue,
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

				var invoker *sqlInvoke
				fo := field.Type.NumOut()
				if fo >= 2 {
					defaultReturnValue := proxy.GetTypeDefaultValue(field.Type.Out(0))
					invoker = newSqlInvoke(target1, target1.ProxyTarget(), methodSetting, xmlele, field.Type.Out(0), providerConfig, defaultReturnValue)
				} else {
					invoker = newSqlInvoke(target1, target1.ProxyTarget(), methodSetting, xmlele, nil, providerConfig, nil)
				}

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
func parseAndGetSqlVariables(sql string) ([]string, string) {
	reg := regexp.MustCompile(`(?m)#\{(\S+?)\}`)
	result := reg.FindAllStringSubmatch(sql, -1)
	if result != nil {
		var r1 []string
		for _, k := range result {
			r1 = append(r1, k[1])
		}
		return r1, reg.ReplaceAllString(sql, "?")
	} else {
		return nil, ""
	}
}
