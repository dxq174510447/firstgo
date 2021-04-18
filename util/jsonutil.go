package util

import "firstgo/frame/vo"
import "github.com/jinzhu/copier"

type jsonUtil struct {
}

const WebSuccess int = 0
const WebFailure int = 500

func (c *jsonUtil) BuildJsonSuccess(r interface{}) *vo.JsonResult {
	var result vo.JsonResult = vo.JsonResult{}

	result.Code = WebSuccess
	result.Data = r

	return &result
}

func (c *jsonUtil) BuildJsonFailure(code int, message string, r interface{}) *vo.JsonResult {
	var result vo.JsonResult = vo.JsonResult{}

	result.Code = code
	result.Data = r
	result.Message = message

	return &result
}

func (c *jsonUtil) BuildJsonFailure1(message string, r interface{}) *vo.JsonResult {
	var result vo.JsonResult = vo.JsonResult{}

	result.Code = WebFailure
	result.Data = r
	result.Message = message

	return &result
}

func (c *jsonUtil) BuildJsonArraySuccess(r interface{}, total int) *vo.JsonResult {
	var result vo.JsonResult = vo.JsonResult{}

	result.Code = WebSuccess

	info := vo.JsonArrayResult{Count: total, Data: r}

	result.Data = info

	return &result
}

// ptr
func (c *jsonUtil) Copy(source interface{}, target interface{}) {
	copier.CopyWithOption(target, source, copier.Option{IgnoreEmpty: true, DeepCopy: true})
}

var JsonUtil jsonUtil = jsonUtil{}
