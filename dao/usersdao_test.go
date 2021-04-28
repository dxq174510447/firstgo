package dao

import (
	"firstgo/frame/context"
	"firstgo/povo/po"
	"testing"
)

func TestName(t *testing.T) {
	local := context.NewLocalStack()
	GetUsersDao().Save123(local, &po.Users{Name: "asdas"})
}
