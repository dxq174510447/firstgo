package test

import (
	"fmt"
	"testing"
	"unsafe"
)

type UserHandler struct {
	Id   int
	Name string
	Age  int
}

func (u *UserHandler) GetName(defaultValue string) (string, error) {
	return u.Name, nil
}
func (u *UserHandler) GetAge(defaultValue int) (int, error) {
	return u.Age, nil
}

var GetName1 = func(defaultValue string) (string, error) {
	return "hehehe", nil
}

//pointer - > Pointer
//Pointer - > pointer
//uintptr - > Pointer
//Pointer - > uintptr

// remember uintptr cannot be stored in variable
// 1. 类型转换 float64 --> uint64  *(*uint64)(unsafe.Pointer(&f))
// 2. Pointer-->uintptr, 产生一个int类型的指针地址值 常用就是打印它
// 3. uintptr-->Pointer 常规下是无效的(枚举除外)，uintptr是整形不是引用，没有指针含义。当gc移除这个对象的时候，uintptr不会更新
// 4. Pointer<-->uintptr 算数级别的双向转换 p = unsafe.Pointer(uintptr(p)+offset) ,作用最多的方法struct的field和array里面的数组
//	// equivalent to f := unsafe.Pointer(&s.f)
//	f := unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Offsetof(s.f))  //struct field
//	// equivalent to e := unsafe.Pointer(&x[i])
//	e := unsafe.Pointer(uintptr(unsafe.Pointer(&x[0])) + i*unsafe.Sizeof(x[0])) //array
//  p := (*int)(unsafe.Pointer(reflect.ValueOf(new(int)).Pointer()))

func TestPtrName(t *testing.T) {

	u := &UserHandler{Id: 1234, Name: "xxxxxx"}
	var m *UserHandler = u
	fmt.Printf("%p %x \n", u, m)

	//TODO 通过f获取对应的值
	f := unsafe.Pointer(uintptr(unsafe.Pointer(u)) + unsafe.Offsetof(u.Id))
	fmt.Printf("%p %x %x %d \n", f, f, &u.Id, *(*int)(f))

	//rt := reflect.ValueOf(u).MethodByName("GetName")
	//fmt.Printf("%d %s \n",uintptr(rt.Pointer()),reflect.ValueOf(rt.Interface()).Kind())
	//rt.SetPointer(unsafe.Pointer(&GetName1))
	//rt.Set(reflect.ValueOf(&GetName1))

	var mm []string = nil

	for _, m1 := range mm {
		fmt.Printf(m1)
	}
}
