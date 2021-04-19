package util

import (
	"firstgo/povo/po"
	"testing"
)

func TestConfigUtil_ClearHttpPath(t *testing.T) {
	r := ConfigUtil.ClearHttpPath("/api/user/name")
	t.Log(r)
}

func TestConfigUtil_Get(t *testing.T) {
	r := ConfigUtil.Get("DB_USER", "platform")
	t.Log(r)
}

func TestConfigUtil_RemovePrefix(t *testing.T) {
	r := ConfigUtil.RemovePrefix("/api/users/", "/api/users")
	r0 := ConfigUtil.RemovePrefix("/api/users", "/api/users")
	r1 := ConfigUtil.RemovePrefix("/api/users/change/status", "/api/users/")
	r2 := ConfigUtil.RemovePrefix("/api/users/change/status", "/api/users")
	t.Log(r, r0, r1, r2)

	e := &po.Users{Id: 123}
	e1 := interface{}(e)
	switch e1.(type) {
	case *po.Users:
		t.Log("ok")
	default:
		t.Log("error")
	}

}
