package frame

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestDispatchServlet_AddRequestMapping(t *testing.T) {
	var m interface{} = ""
	var n = reflect.ValueOf(m).Kind()
	a, _ := json.Marshal(&m)
	b := string(a)
	switch n {
	case reflect.Struct:
		fmt.Println("im struct")
		fmt.Println(b)
	case reflect.Ptr:
		fmt.Println("im ptr")
		fmt.Println(b)
	case reflect.String:
		fmt.Println("im string")
		fmt.Println(b)
	case reflect.Interface:
		fmt.Println("im Interface")
		fmt.Println(b)
	}
}
