package db

import (
	"encoding/xml"
	"firstgo/frame/proxy"
)

type MapperElementXml struct {
	Id  string `xml:"id,attr"`
	Sql string `xml:",innerxml"`
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

func (m *MapperFactory) ParseXml(target proxy.ProxyTarger, content string) *MapperXml {
	mapper := &MapperXml{}
	err := xml.Unmarshal([]byte(content), mapper)
	if err != nil {
		panic(err)
	}
	return mapper
}

var mapperFactory MapperFactory = MapperFactory{}

func GetMapperFactory() *MapperFactory {
	return &mapperFactory
}

func AddMapperProxyTarget(target1 proxy.ProxyTarger) {
	proxy.AddClassProxy(target1)
}
