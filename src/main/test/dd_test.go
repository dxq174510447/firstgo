package test

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"
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

	//u := &UserHandler{Id: 1234, Name: "xxxxxx"}
	//var m *UserHandler = u
	//fmt.Printf("%p %x \n", u, m)
	//
	////TODO 通过f获取对应的值
	//f := unsafe.Pointer(uintptr(unsafe.Pointer(u)) + unsafe.Offsetof(u.Id))
	//fmt.Printf("%p %x %x %d \n", f, f, &u.Id, *(*int)(f))

	//rt := reflect.ValueOf(u).MethodByName("GetName")
	//fmt.Printf("%d %s \n",uintptr(rt.Pointer()),reflect.ValueOf(rt.Interface()).Kind())
	//rt.SetPointer(unsafe.Pointer(&GetName1))
	//rt.Set(reflect.ValueOf(&GetName1))

	//var mm []string = nil
	//
	//for _, m1 := range mm {
	//	fmt.Printf(m1)
	//}
	//
	//var m2 []interface{}
	//m2 = append(m2, map[string]string{"a": "b"})
	//m2 = append(m2, map[string]int{"a": 1})
	//fmt.Println(m2)

	//var path string = "/Users/klook/log/access.log"
	//var tag string = "tt000"
	//var tag1 string = "tt111"
	//var tag2 string = "tt222"
	//go testFile(path,tag)
	//go testFile(path,tag1)
	//go testFile(path,tag2)
	//var m chan int = make(chan int)
	//<- m
	//f := "%date %date{2006-01-02} %date{2006-01-02T15:04:05Z07:00} %-5thread %-5level %logger %logger{5} %-30logger{5} %thread %-5line %file %msg %n "
	//reg := regexp.MustCompile(`(?m)(^\s+|\s+$)`)
	//reg := regexp.MustCompile("%(\\{[^\\}]+\\})?(date|thread|level|logger|line|file|msg|n)(\\{[^\\}]+\\})?")
	//last := reg.ReplaceAllStringFunc(f,func(row string) string{
	//	fmt.Println(row)
	//	return ""
	//})
	//fmt.Println("--->",last)
	// date|thread|level|line|file|msg|n|logger
	//var date1 = regexp.MustCompile("%date(\\{[^\\}]+\\})?")
	//f = date1.ReplaceAllStringFunc(f, func(row string) string {
	//	p := strings.Index(row, "{")
	//	dateFormat := "2006-01-02 15:04:05"
	//	if p >= 0 {
	//		p1 := strings.Index(row, "}")
	//		dateFormat = row[p+1 : p1]
	//	}
	//	return fmt.Sprintf(`{{logDate "%s"}}`, dateFormat)
	//})
	//var thread1 = regexp.MustCompile("%(\\-\\d+)?thread")
	//f = thread1.ReplaceAllStringFunc(f, func(row string) string {
	//	p := strings.Index(row, "-")
	//	maxSize := 0
	//	if p >= 0 {
	//		p1 := strings.Index(row, "thread")
	//		maxSizeStr := row[p+1 : p1]
	//		maxSize, _ = strconv.Atoi(maxSizeStr)
	//	}
	//	return fmt.Sprintf(`{{logThread %d}}`, maxSize)
	//})
	//var level1 = regexp.MustCompile("%(\\-\\d+)?level")
	//f = level1.ReplaceAllStringFunc(f, func(row string) string {
	//	p := strings.Index(row, "-")
	//	maxSize := 0
	//	if p >= 0 {
	//		p1 := strings.Index(row, "level")
	//		maxSizeStr := row[p+1 : p1]
	//		maxSize, _ = strconv.Atoi(maxSizeStr)
	//	}
	//	return fmt.Sprintf(`{{logLevel %d}}`, maxSize)
	//})
	//
	////line|file|msg|n|logger
	//var line1 = regexp.MustCompile("%(\\-\\d+)?line")
	//f = line1.ReplaceAllStringFunc(f, func(row string) string {
	//	p := strings.Index(row, "-")
	//	maxSize := 0
	//	if p >= 0 {
	//		p1 := strings.Index(row, "line")
	//		maxSizeStr := row[p+1 : p1]
	//		maxSize, _ = strconv.Atoi(maxSizeStr)
	//	}
	//	return fmt.Sprintf(`{{logLine %d}}`, maxSize)
	//})
	//
	//var file1 = regexp.MustCompile("%(\\-\\d+)?file")
	//f = file1.ReplaceAllStringFunc(f, func(row string) string {
	//	p := strings.Index(row, "-")
	//	maxSize := 0
	//	if p >= 0 {
	//		p1 := strings.Index(row, "file")
	//		maxSizeStr := row[p+1 : p1]
	//		maxSize, _ = strconv.Atoi(maxSizeStr)
	//	}
	//	return fmt.Sprintf(`{{logFile %d}}`, maxSize)
	//})
	////msg|n|logger
	//var msg1 = regexp.MustCompile("%(\\-\\d+)?msg")
	//f = msg1.ReplaceAllStringFunc(f, func(row string) string {
	//	return fmt.Sprintf(`{{logMsg %d}}`, 0)
	//})
	//var br = regexp.MustCompile("%(\\-\\d+)?n")
	//f = br.ReplaceAllStringFunc(f, func(row string) string {
	//	//return fmt.Sprintf(`{{logBr %d}}`,0)
	//	return `\n`
	//})
	////%logger{n}
	//var logger1 = regexp.MustCompile("%(\\-\\d+)?logger(\\{[^\\}]+\\})?")
	//f = logger1.ReplaceAllStringFunc(f, func(row string) string {
	//	//return fmt.Sprintf(`{{logBr %d}}`,0)
	//	p := strings.Index(row, "-")
	//	maxSize := 0
	//	if p >= 0 {
	//		p1 := strings.Index(row, "logger")
	//		maxSizeStr := row[p+1 : p1]
	//		maxSize, _ = strconv.Atoi(maxSizeStr)
	//	}
	//
	//	p1 := strings.Index(row, "{")
	//	clazzSize := -1
	//	if p1 >= 0 {
	//		p2 := strings.Index(row, "}")
	//		clazzSizeStr := row[p1+1 : p2]
	//		clazzSize, _ = strconv.Atoi(clazzSizeStr)
	//	}
	//	return fmt.Sprintf(`{{logLogger %d %d}}`, maxSize, clazzSize)
	//})
	//
	//fmt.Println("--->", f)
	//l := 10
	m := "Zbcd.Abg.go"
	n := m[0]
	fmt.Println(reflect.ValueOf(n).Type())
	fmt.Println(n)
	//A-Z
	if n >= 65 && n <= 90 {
		fmt.Println("A_Z")
	} else {
		fmt.Println("nonA_Z")
	}
	//n := []byte(m)
	//sp := make([]byte, 22, 22)
	//for i := 0; i < 22; i++ {
	//	sp[i] = 32
	//}
	//copy(sp, n)
	//n2 := string(sp)
	//fmt.Printf("%s-%s--\n", m, n2)
}

func testFile(path string, tag string) {
	f1, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer func() {
		f1.Close()
	}()
	b1 := bufio.NewWriter(f1)
	i := 0
	for {
		b1.WriteString(fmt.Sprintf("begin 测试生生世世生生世世 %s %d\n测试生生世世生生世世 %s %d end \n", tag, i, tag, i))
		i = i + 1
		if i/10 == 0 {
			b1.Flush()
			time.Sleep(2 * time.Second)
		}
	}
	f1.Sync()
}
