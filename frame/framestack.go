package frame

type FrameStack struct {
	element []map[string]interface{}
}

func (f *FrameStack) Push() map[string]interface{} {
	ele := make(map[string]interface{})
	if f.element == nil {
		f.element = []map[string]interface{}{ele}
	} else {
		f.element = append(f.element, ele)
	}
	return ele
}

func (f *FrameStack) Pop() map[string]interface{} {
	result := f.element[len(f.element)-1]
	f.element = f.element[0 : len(f.element)-1]
	return result
}

func (f *FrameStack) Peek() map[string]interface{} {
	return f.element[len(f.element)-1]
}

func (c *FrameStack) Set(key string, value interface{}) {
	top := c.Peek()
	top[key] = value
}

func (c *FrameStack) Get(key string) interface{} {

	for i := (len(c.element) - 1); i >= 0; i-- {
		ele := c.element[i]
		if res, ok := ele[key]; ok {
			return res
		}
	}
	return nil
}

func NewFrameStack() *FrameStack {
	result := &FrameStack{}
	result.Push()
	return result
}
