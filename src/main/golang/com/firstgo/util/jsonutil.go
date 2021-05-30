package util

import "github.com/dxq174510447/goframe/lib/frame/vo"
import util2 "github.com/dxq174510447/goframe/lib/frame/util"
import "github.com/jinzhu/copier"

type jsonUtil struct {
}

// BuildJsonSuccess 构建成功返回
func (c *jsonUtil) BuildJsonSuccess(r interface{}) *vo.JsonResult {
	return util2.JsonUtil.BuildJsonSuccess(r)
}

// BuildJsonFailure 构建失败返回
func (c *jsonUtil) BuildJsonFailure(code int, message string, r interface{}) *vo.JsonResult {
	return util2.JsonUtil.BuildJsonFailure(code, message, r)
}

// BuildJsonFailure1 构建失败返回
func (c *jsonUtil) BuildJsonFailure1(message string, r interface{}) *vo.JsonResult {
	return util2.JsonUtil.BuildJsonFailure1(message, r)
}

// BuildJsonArraySuccess 构建返回数组 例如分页查询
func (c *jsonUtil) BuildJsonArraySuccess(r interface{}, total int) *vo.JsonResult {
	return util2.JsonUtil.BuildJsonArraySuccess(r, total)
}

// Copy 结构体拷贝
func (c *jsonUtil) Copy(source interface{}, target interface{}) {
	copier.CopyWithOption(target, source, copier.Option{IgnoreEmpty: true, DeepCopy: true})
}

var JsonUtil jsonUtil = jsonUtil{}
