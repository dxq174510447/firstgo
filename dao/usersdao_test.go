package dao

import (
	"firstgo/frame/context"
	_ "firstgo/frame/db/filter"
	"firstgo/povo/po"
	"fmt"
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
	//m6, err6 := GetUsersDao().Find6(local, &po.Users{Id: 4, Name: "haha4", Status: 4}, &po.Users{Id: 5, Name: "haha", Status: 5})
	//fmt.Println(m6, err6)

}
