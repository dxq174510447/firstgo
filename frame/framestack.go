package frame

// 针对goroute设置的环境变量 嵌套
type FrameStack struct {
	element []map[string]interface{}
}

// Push 新增层次
func (f *FrameStack) Push() map[string]interface{} {
	ele := make(map[string]interface{})
	if f.element == nil {
		f.element = []map[string]interface{}{ele}
	} else {
		f.element = append(f.element, ele)
	}
	return ele
}

// Pop 出栈
func (f *FrameStack) Pop() map[string]interface{} {
	result := f.element[len(f.element)-1]
	f.element = f.element[0 : len(f.element)-1]
	return result
}

// Peek 查看栈顶的环境设置
func (f *FrameStack) Peek() map[string]interface{} {
	return f.element[len(f.element)-1]
}

// Set 在栈顶环境变量设置参数
func (c *FrameStack) Set(key string, value interface{}) {
	top := c.Peek()
	top[key] = value
}

// Get 从栈中依次取出环境变量key，从栈顶开始
func (c *FrameStack) Get(key string) interface{} {

	for i := (len(c.element) - 1); i >= 0; i-- {
		ele := c.element[i]
		if res, ok := ele[key]; ok {
			return res
		}
	}
	return nil
}

// NewFrameStack 创建新的变量栈
func NewFrameStack() *FrameStack {
	result := &FrameStack{}
	result.Push()
	return result
}
