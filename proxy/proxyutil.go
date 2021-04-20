package util

import (
	"fmt"
	"reflect"
)

// Computer 想办法实现java aop拦截
// 思路 ?

type ProxyIntercepter interface {
	Execute(target ProxyTarget, method ProxyTargetMethod,
		invoker reflect.Value, arg []reflect.Value) []reflect.Value
}

type ProxyTarget struct {
	ClassName string
	ClassType reflect.Type
	HttpPath  string
	Tags      []string
	Methods   []ProxyTargetMethod
}

type ProxyTargetMethod struct {
	HttpPath   string
	HttpMethod string
	HttpView   string
	MethodName string
	Intercepts []*ProxyIntercepter
	Tags       []string
}

type ProxyTargeter interface {
	GetProxyTarget() *ProxyTarget
}

// --------

type LogProxyIntercept struct {
}

func (l *LogProxyIntercept) Execute(target ProxyTarget, method ProxyTargetMethod,
	invoker reflect.Value, arg []reflect.Value) []reflect.Value {
	fmt.Println("log being")
	defer fmt.Println("log end")

}

type ComputerService struct {
	R1Impl      func(c *ComputerService)
	R2Impl      func(c *ComputerService)
	proxyTarget *ProxyTarget
}

func (c *ComputerService) R1() {
	c.R1Impl(c)
}
func (c *ComputerService) R2() {
	c.R2Impl(c)
}
func (c *ComputerService) GetProxyTarget() *ProxyTarget {
	return c.proxyTarget
}

var ComputerServiceImpl ComputerService = ComputerService{
	proxyTarget: &ProxyTarget{},
	R1Impl: func(c *ComputerService) {
		fmt.Println("invoke r1 impl")
	},
	R2Impl: func(c *ComputerService) {
		fmt.Println("invoke r2 impl")
	},
}

//func (c *ComputerServiceImpl) TxWrite(env *vo.UsersVo, data *po.Users ,retry int) (*po.Users,[]*po.Users,int,string){
//	return nil,nil,0,""
//}
//
//func (c *ComputerServiceImpl) TxRdRead(env *vo.UsersVo, data *po.Users ,retry int) (*po.Users,[]*po.Users,int,string) {
//	return nil,nil,0,""
//}
//
//func (c *ComputerServiceImpl) Read(env *vo.UsersVo, data *po.Users ,retry int) (*po.Users,[]*po.Users,int,string) {
//	return nil,nil,0,""
//}
//
//func (c *ComputerServiceImpl) xxxx(env *vo.UsersVo, data *po.Users ,retry int) (*po.Users,[]*po.Users,int,string) {
//	return nil,nil,0,""
//}

type proxyUtil struct {
}

func (p *proxyUtil) NewProxy(target interface{}) interface{} {

	if proxyService, ok := target.(ProxyTargeter); ok {
		classinfo := proxyService.GetProxyTarget()
		for _, method := range classinfo.Methods {
			if len(method.Intercepts) != 0 {

			}
		}
	}

	return nil
}

var ProxyUtil proxyUtil = proxyUtil{}

func init() {

	var ntptr interface{} = ProxyUtil.NewProxy(&ComputerServiceImpl)
	if ntptr != nil {
		nt := ntptr.(*ComputerService)
		ComputerServiceImpl = *nt
	}

}
