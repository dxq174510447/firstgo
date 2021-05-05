package dao

import (
	"firstgo/frame/context"
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	local := context.NewLocalStack()
	m, err := GetUsersDao().Find1(local)
	fmt.Println(m, err)
}
