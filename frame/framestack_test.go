package frame

import (
	"firstgo/povo/po"
	"fmt"
	"testing"
)

func TestNewStack(t *testing.T) {

	s1 := NewFrameStack()
	defer func() {
		s1.Pop()
		s1 = nil
		fmt.Println("clear")

		if err := recover(); err != nil {
			panic(err)
		}

	}()
	s1.Set("a", 1)
	s1.Set("b", "2")
	s1.Set("c", &po.Users{Name: "parent"})

	s1.Push()
	s1.Set("a", 12)
	s1.Set("b", "22")
	s1.Set("c", &po.Users{Name: "son"})

	fmt.Println(s1.Get("a"))
	fmt.Println(s1.Get("b"))
	fmt.Println(s1.Get("c"))
	s1.Pop()
	fmt.Println(s1.Get("a"))
	fmt.Println(s1.Get("b"))
	fmt.Println(s1.Get("c"))
	panic(fmt.Errorf("%s", "self error"))
}
