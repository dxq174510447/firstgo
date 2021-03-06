package listener

import (
	"github.com/dxq174510447/goframe/lib/frame/application"
	"github.com/dxq174510447/goframe/lib/frame/context"
	"github.com/dxq174510447/goframe/lib/frame/event"
	"github.com/dxq174510447/goframe/lib/frame/http"
	"github.com/dxq174510447/goframe/lib/frame/log/logclass"
	"github.com/dxq174510447/goframe/lib/frame/proxy/proxyclass"
	"github.com/dxq174510447/goframe/lib/frame/util"
	"reflect"
)

type WebStartedEchoListener struct {
	Logger logclass.AppLoger `FrameAutowired:""`
}

func (w *WebStartedEchoListener) OnEvent(local *context.LocalStack, event event.FrameEventer) error {
	f := util.ClassUtil.GetClassNameByType(reflect.TypeOf(event).Elem())
	w.Logger.Info(local, "启动 %s", f)
	return nil
}

func (w *WebStartedEchoListener) WatchEvent() event.FrameEventer {
	e := &http.WebServletStartedEvent{}
	return event.FrameEventer(e)
}

func (w *WebStartedEchoListener) Order() int {
	return 0
}

func (w *WebStartedEchoListener) ProxyTarget() *proxyclass.ProxyClass {
	return nil
}

var webStartedEchoListener WebStartedEchoListener = WebStartedEchoListener{}

func init() {
	application.AddProxyInstance("", proxyclass.ProxyTarger(&webStartedEchoListener))
}
