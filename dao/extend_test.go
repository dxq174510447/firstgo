package dao

import (
	"fmt"
	"reflect"
	"testing"
)

type Person struct {
	Name string
	Sex  int
}

func (p *Person) Save(entity interface{}) (int, error) {
	return 0, nil
}

func (p *Person) Learn(clazz string) {
	fmt.Println(clazz, "---->", p.Name)
}

type Student struct {
	Person
	StudentNo string
}

func (p *Student) Learn(clazz string) {
	fmt.Println(clazz, p.Name)
}

func TestGetUsersDao(t *testing.T) {
	s := &Student{StudentNo: "no101"}
	s.Learn("shuxue")
	s.Save(nil)
	s.Learn("haha")
	fmt.Println(s.Name)
	m := reflect.TypeOf(s).Elem().NumField()
	for i := 0; i < m; i++ {
		f := reflect.TypeOf(s).Elem().Field(i)
		fmt.Println(f.Name, f.Type.Kind(), f.Type.String())
		if f.Type == reflect.TypeOf(Person{}) {
			reflect.ValueOf(s).Elem().FieldByName(f.Name).FieldByName("Name").Set(reflect.ValueOf("dynamic"))
		}
	}
	fmt.Println(s.Name)
	s.Learn("haha")
}
