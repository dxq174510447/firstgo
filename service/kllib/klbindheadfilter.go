package kllib

import (
	"firstgo/frame/context"
	http2 "firstgo/frame/http"
	"firstgo/service/kllib/vo"
	"firstgo/util"
	"net/http"
	"strings"
)

// KlBindHeadFilter HttpFilter 把kl头信息放在线程变量里面
type KlBindHeadFilter struct {
}

func (b *KlBindHeadFilter) DoFilter(local *context.LocalStack,
	request *http.Request, response http.ResponseWriter, chain http2.FilterChain) {

	klreqId := strings.TrimSpace(request.Header.Get(KlRequestIdKey))

	if klreqId == "" {
		klreqId = util.WebUtil.GenKlRequestId()
	}

	klheader := &vo.KlHeaderVo{
		RequestID: klreqId,
	}
	SetCurrentKlHeader(local, klheader)

	chain.DoFilter(local, request, response)
}

func (b *KlBindHeadFilter) Order() int {
	return 5
}

var klBindHeadFilter KlBindHeadFilter = KlBindHeadFilter{}

func init() {
	http2.AddFilter(http2.Filter(&klBindHeadFilter))
}
