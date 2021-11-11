package util

import "github.com/jinzhu/copier"

type jsonUtil struct {
}


// Copy 结构体拷贝
func (c *jsonUtil) Copy(source interface{}, target interface{}) {
	copier.CopyWithOption(target, source, copier.Option{IgnoreEmpty: true, DeepCopy: true})
}

var JsonUtil jsonUtil = jsonUtil{}
