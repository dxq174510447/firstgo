package dbcore

import (
	"encoding/xml"
	"firstgo/frame/context"
	"firstgo/frame/proxy"
	"fmt"
	"reflect"
	"strings"
)

type MapperElementXml struct {
	Id      string `xml:"id,attr"`
	Sql     string `xml:",innerxml"`
	SqlType string
}

type MapperXml struct {
	Sql []MapperElementXml `xml:"sql"`

	SelectSql []MapperElementXml `xml:"select"`

	UpdateSql []MapperElementXml `xml:"update"`

	DeleteSql []MapperElementXml `xml:"delete"`

	InsertSql []MapperElementXml `xml:"insert"`
}

type MapperFactory struct {
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
			refs[ele.Id] = &ele
		}
	}

	if len(mapper.UpdateSql) > 0 {
		for _, ele := range mapper.UpdateSql {
			ele.SqlType = SqlTypeUpdate
			refs[ele.Id] = &ele
		}
	}

	if len(mapper.InsertSql) > 0 {
		for _, ele := range mapper.InsertSql {
			ele.SqlType = SqlTypeInsert
			refs[ele.Id] = &ele
		}
	}

	if len(mapper.SelectSql) > 0 {
		for _, ele := range mapper.SelectSql {
			ele.SqlType = SqlTypeSelect
			refs[ele.Id] = &ele
		}
	}

	if len(mapper.DeleteSql) > 0 {
		for _, ele := range mapper.DeleteSql {
			ele.SqlType = SqlTypeDelete
			refs[ele.Id] = &ele
		}
	}
	fmt.Println(refs)

	return refs
}

var mapperFactory MapperFactory = MapperFactory{}

func GetMapperFactory() *MapperFactory {
	return &mapperFactory
}

type sqlInvoke struct {
	target interface{}
	clazz  *proxy.ProxyClass
	method *proxy.ProxyMethod
	mapper map[string]*MapperElementXml
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

func (s *sqlInvoke) invokeSelect(context *context.LocalStack, args []reflect.Value, sql *MapperElementXml) []reflect.Value {
	fmt.Println(sql.Id, sql.Sql)
	return nil
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

func newSqlInvoke(
	target interface{},
	clazz *proxy.ProxyClass,
	method *proxy.ProxyMethod, mapper map[string]*MapperElementXml) *sqlInvoke {
	return &sqlInvoke{
		target: target,
		clazz:  clazz,
		method: method,
		mapper: mapper,
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

				invoker := newSqlInvoke(target1, target1.ProxyTarget(), methodSetting, xmlele)
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
