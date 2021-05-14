package main

import (
	_ "firstgo/controller"
	_ "firstgo/dao"
	_ "firstgo/frame"
	_ "firstgo/service"
	"net/http"
)

func main() {

	//defer statsd.StartHttpHandlerStatsCollector("service/klook-p2pbackend/p2pbackend", "01",
	//	";;nats://ip-172-31-20-207.ap-southeast-1.compute.internal:6222").Stop()

	http.ListenAndServe(":8080", nil)

}
