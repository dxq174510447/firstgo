package test

import (
	"encoding/json"
	"firstgo/src/main/golang/com/firstgo/povo/po"
	"fmt"
	"reflect"
	"testing"
)

func TestName(t *testing.T) {
	var m map[string]*po.Users = make(map[string]*po.Users)
	ff(m)
	s, _ := json.Marshal(m)
	fmt.Println(string(s))
}

func ff(p interface{}) {
	rt := reflect.TypeOf(p)
	rv := reflect.ValueOf(p)
	switch rt.Kind() {
	case reflect.Map:
		//fmt.Println(rt.Key(),rt.Elem())
		elerv := reflect.New(rt.Elem().Elem())
		//elert := reflect.TypeOf(rt.Elem().Elem())
		elerv.Elem().FieldByName("Id").Set(reflect.ValueOf(456))
		rv.SetMapIndex(reflect.ValueOf("123"), elerv)
	}
}
