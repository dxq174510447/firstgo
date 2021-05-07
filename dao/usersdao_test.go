package dao

import (
	"firstgo/frame/context"
	_ "firstgo/frame/db/filter"
	"firstgo/povo/po"
	"fmt"
	"reflect"
	"testing"
)

func TestName(t *testing.T) {
	local := context.NewLocalStack()
	//m1, err1 := GetUsersDao().Find1(local)
	//fmt.Println(m1, err1)
	//m2, err2 := GetUsersDao().Find2(local, 123)
	//fmt.Println(m2, err2)
	//m3, err3 := GetUsersDao().Find3(local, 123, "456")
	//fmt.Println(m3, err3)
	//m4, err4 := GetUsersDao().Find4(local, &po.Users{Id: 4, Name: "haha", Status: 4})
	//fmt.Println(m4, err4)
	m5, err5 := GetUsersDao().Find5(local, &po.Users{Id: 94, Name: "haha", Status: 4}, 123)
	fmt.Println(m5, err5)
	if err5 != nil {
		fmt.Println(err5)
		panic(err5)
	} else {
		if len(m5) > 0 {
			fmt.Println(len(m5))
			for _, v := range m5 {
				fmt.Println(*v)
			}
			fmt.Println(reflect.ValueOf(m5[0].CreateDate).Kind())
			fmt.Println(reflect.ValueOf(m5[0].CreateTime).Kind())
		} else {
			fmt.Println("result length 0")
		}
	}
	//m6, err6 := GetUsersDao().Find6(local, &po.Users{Id: 4, Name: "haha4", Status: 4}, &po.Users{Id: 5, Name: "haha", Status: 5})
	//fmt.Println(m6, err6)

}
