package main

import (
	_ "firstgo/controller"
	_ "firstgo/dao"
	_ "firstgo/frame"
	_ "firstgo/service"
	"net/http"
)

func main() {

	http.ListenAndServe(":8080", nil)

}
