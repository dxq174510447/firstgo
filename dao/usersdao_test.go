package dao

import (
	"firstgo/frame/context"
	"firstgo/povo/po"
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	local := context.NewLocalStack()
	m, err := GetUsersDao().Save123(local, &po.Users{Name: "asdas"})
	fmt.Println(m, err)
}
