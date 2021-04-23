package main

import (
	_ "firstgo/controller"
	_ "firstgo/frame/context"
	_ "firstgo/frame/db"
	_ "firstgo/frame/http"
	_ "firstgo/frame/proxy"
	"net/http"
)

func main() {

	http.ListenAndServe(":8080", nil)

}
