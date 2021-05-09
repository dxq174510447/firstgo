package dbcore

import (
	"bytes"
	"database/sql"
	"encoding/xml"
	"firstgo/frame/context"
	"firstgo/frame/exception"
	"firstgo/frame/proxy"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"text/template"
	"time"
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

type sqlColumnType struct {
	column      *sql.ColumnType
	field       *reflect.StructField
	defaultType reflect.Type
}

type sqlInvoke struct {
	target         interface{}
	clazz          *proxy.ProxyClass
	method         *proxy.ProxyMethod
	mapper         map[string]*MapperElementXml
	providerConfig *SqlProviderConfig
	//slice int string ptr float64
	returnSqlType reflect.Type

	//具体返回的类型 如果返回的是指针 就对应的是结构体 如果返回的是slice 就对应的里面的元素类型，如果元素是指针就是对应的结构体 否则就是int,string等
	returnSqlElementType reflect.Type
	defaultReturnValue   *reflect.Value
	//如果返回的是结构体类型 字段field
	structFieldMap map[string]reflect.StructField
	sqlFieldMap    []*sqlColumnType
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
	var nilError = reflect.Zero(reflect.TypeOf((*error)(nil)).Elem())

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

	if s.returnSqlType != nil {
		return []reflect.Value{*queryResult, nilError}
	} else {
		return []reflect.Value{nilError}
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

	//test
	//r1, e1 := result.ColumnTypes()
	//if e1 == nil {
	//	for _, r11 := range r1 {
	//		fmt.Println(r11)
	//	}
	//}
	//
	//r2, e2 := result.Columns()
	//if e2 == nil {
	//	for _, r11 := range r2 {
	//		fmt.Println(r11)
	//	}
	//}

	if s.sqlFieldMap == nil {
		r1, e1 := result.ColumnTypes()
		if e1 == nil {
			sts := make([]*sqlColumnType, len(r1), len(r1))
			for k, ct := range r1 {
				var m1 *sqlColumnType = s.coverToGoType(ct)
				sts[k] = m1
			}
			s.sqlFieldMap = sts
		}
	}

	for result.Next() {
		if queryCount != 0 && currentCount >= pageSize {
			total = reflect.AppendSlice(total, current)
			current = reflect.MakeSlice(s.returnSqlType, 0, pageSize)
			currentCount = 0
		}

		r1, err2 := s.scanRow(result)
		if err2 != nil {
			if errorFlag == 0 {
				panic(err2)
			} else {
				return nil, err2
			}
		}

		current = reflect.Append(current, *r1)

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

func (s *sqlInvoke) scanRow(result *sql.Rows) (*reflect.Value, error) {
	valueptr := make([]interface{}, len(s.sqlFieldMap), len(s.sqlFieldMap))
	for k, v := range s.sqlFieldMap {
		d1 := GetSqlFieldReturnDefaultValue(v.defaultType)
		fmt.Println(v.field.Name, v.defaultType.Kind(), reflect.ValueOf(d1).Elem().Kind())
		valueptr[k] = d1
	}
	e1 := result.Scan(valueptr...) //不scan会导致连接不释放
	if e1 != nil {
		return nil, e1
	}
	var result1 *reflect.Value
	if s.returnSqlElementType.Kind() == reflect.Struct {
		hp := reflect.New(s.returnSqlElementType)
		hv := hp.Elem()
		for k, v := range s.sqlFieldMap {
			if v.field != nil {
				SetEntityFieldValue(&hv, v.field, valueptr[k])
			}
		}
		result1 = &hp
	}
	return result1, nil
}

func (s *sqlInvoke) coverToGoType(ct *sql.ColumnType) *sqlColumnType {
	result := sqlColumnType{column: ct}
	addDefaultType := true

	if s.structFieldMap != nil {
		columnName := ct.Name()
		fieldName := proxy.GetCamelCaseName(columnName)
		if field, ok := s.structFieldMap[fieldName]; ok {
			addDefaultType = false
			result.field = &field
			result.defaultType = field.Type
		}
	}
	fmt.Println(result.column.Name(), result.field.Type)
	if addDefaultType {
		databasetype := strings.ToLower(ct.DatabaseTypeName())
		if strings.Index(databasetype, "int") >= 0 {
			result.defaultType = reflect.TypeOf(int(1))
		} else if strings.Index(databasetype, "decimal") >= 0 {
			result.defaultType = reflect.TypeOf(float64(1.0))
		} else if strings.Index(databasetype, "char") >= 0 {
			result.defaultType = reflect.TypeOf(string(""))
		} else if strings.Index(databasetype, "date") >= 0 {
			n := time.Now()
			result.defaultType = reflect.TypeOf(&n)
		} else {
			result.defaultType = reflect.TypeOf(string(""))
		}
	}
	return &result
}

func newSqlInvoke(
	target interface{}, //对象
	clazz *proxy.ProxyClass, //代理信息
	method *proxy.ProxyMethod, //当前方法代理信息
	mapper map[string]*MapperElementXml, //对应sql节点
	returnSqlType reflect.Type, //返回的类型 不是error 如果没有就nil
	providerConfig *SqlProviderConfig,
	defaultReturnValue *reflect.Value, //默认返回类型值
	structFieldMap map[string]reflect.StructField,
) *sqlInvoke {

	var returnSqlElementType reflect.Type = nil
	if returnSqlType != nil {
		switch returnSqlType.Kind() {
		case reflect.Slice:
			if returnSqlType.Elem().Kind() == reflect.Ptr {
				returnSqlElementType = returnSqlType.Elem().Elem()
			} else {
				returnSqlElementType = returnSqlType.Elem()
			}
		case reflect.Ptr:
			if returnSqlType.Elem().Kind() == reflect.Struct {
				returnSqlElementType = returnSqlType.Elem()
			} else {
				returnSqlElementType = returnSqlType
			}
		default:
			returnSqlElementType = returnSqlType
		}
	}

	return &sqlInvoke{
		target:               target,
		clazz:                clazz,
		method:               method,
		mapper:               mapper,
		providerConfig:       providerConfig,
		returnSqlType:        returnSqlType,
		defaultReturnValue:   defaultReturnValue,
		structFieldMap:       structFieldMap,
		returnSqlElementType: returnSqlElementType,
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
					defaultReturnValue := proxy.GetMethodReturnDefaultValue(field.Type.Out(0))
					structFields := proxy.GetStructField(field.Type.Out(0))
					invoker = newSqlInvoke(target1, target1.ProxyTarget(),
						methodSetting,
						xmlele, field.Type.Out(0),
						providerConfig, defaultReturnValue, structFields)
				} else {
					invoker = newSqlInvoke(target1, target1.ProxyTarget(), methodSetting,
						xmlele, nil,
						providerConfig, nil, nil)
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

// GetSqlFieldReturnDefaultValue  用来接受sql column 返回的值 类型 int int64 float64 string *Time
// 当有错误的时候 返回这个默认结果的指针 和 错误
func GetSqlFieldReturnDefaultValue(rtType reflect.Type) interface{} {
	switch rtType.Kind() {
	case reflect.String:
		v := sql.NullString{}
		return &v
	case reflect.Int64:
		v := sql.NullInt64{}
		return &v
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
		v := sql.NullInt32{}
		return &v
	case reflect.Float32, reflect.Float64:
		v := sql.NullFloat64{}
		return &v
	case reflect.Ptr:
		eleType := rtType.Elem()
		if eleType == GoTimeType {
			v := sql.NullTime{}
			return &v
		} else if eleType == SqlNullStringType {
			v := sql.NullString{}
			return &v
		} else if eleType == SqlNullTimeType {
			v := sql.NullTime{}
			return &v
		} else if eleType == SqlNullInt64Type {
			v := sql.NullInt64{}
			return &v
		} else if eleType == SqlNullInt32Type {
			v := sql.NullInt32{}
			return &v
		} else if eleType == SqlNullFloat64Type {
			v := sql.NullFloat64{}
			return &v
		} else if eleType == SqlNullBoolType {
			v := sql.NullBool{}
			return &v
		}
		return nil
	default:
		panic(fmt.Sprintf("%s找不到对应默认值", rtType.String()))
	}
}

// SetEntityFieldValue target 目标对象  name fieldname value sql.null* 的指针
// 如果value指针对应的是struct 根据ptr来判断是否是指针类型
func SetEntityFieldValue(target *reflect.Value, field *reflect.StructField, value interface{}) {
	switch field.Type.Kind() {
	case reflect.String:
		s := value.(*sql.NullString)
		if s.Valid {
			(*target).FieldByName(field.Name).Set(reflect.ValueOf(s.String))
		}
	case reflect.Int64:
		s := value.(*sql.NullInt64)
		if s.Valid {
			(*target).FieldByName(field.Name).Set(reflect.ValueOf(s.Int64))
		}
	case reflect.Int:
		s := value.(*sql.NullInt32)
		if s.Valid {
			(*target).FieldByName(field.Name).Set(reflect.ValueOf(int(s.Int32)))
		}
	case reflect.Float64:
		s := value.(*sql.NullFloat64)
		if s.Valid {
			(*target).FieldByName(field.Name).Set(reflect.ValueOf(s.Float64))
		}
	case reflect.Ptr:
		//fix field is *Time ,same as ptr
		eleType := field.Type.Elem()
		if eleType == GoTimeType {
			s := value.(*sql.NullTime)
			if s.Valid {
				(*target).FieldByName(field.Name).Set(reflect.ValueOf(&(s.Time)))
			}
		} else if eleType == SqlNullStringType {
			s := value.(*sql.NullString)
			if s.Valid {
				(*target).FieldByName(field.Name).Set(reflect.ValueOf(s))
			}
		} else if eleType == SqlNullTimeType {
			s := value.(*sql.NullTime)
			if s.Valid {
				(*target).FieldByName(field.Name).Set(reflect.ValueOf(s))
			}
		} else if eleType == SqlNullInt64Type {
			s := value.(*sql.NullInt64)
			if s.Valid {
				(*target).FieldByName(field.Name).Set(reflect.ValueOf(s))
			}
		} else if eleType == SqlNullInt32Type {
			s := value.(*sql.NullInt32)
			if s.Valid {
				(*target).FieldByName(field.Name).Set(reflect.ValueOf(s))
			}
		} else if eleType == SqlNullFloat64Type {
			s := value.(*sql.NullFloat64)
			if s.Valid {
				(*target).FieldByName(field.Name).Set(reflect.ValueOf(s))
			}
		}
	}

}

func GetSqlNullTypeValue(p interface{}) interface{} {
	if p == nil {
		return nil
	}
	if reflect.ValueOf(p).IsZero() {
		return nil
	}
	switch p1 := p.(type) {
	case *sql.NullString:
		return p1.String
	case *sql.NullInt64:
		return p1.Int64
	case *sql.NullInt32:
		return int(p1.Int32)
	case *sql.NullFloat64:
		return p1.Float64
	case *sql.NullTime:
		return p1.Time
	case *sql.NullBool:
		return p1.Bool
	}
	return nil
}
