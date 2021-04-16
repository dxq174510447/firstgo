package main

import "net/http"
import _ "firstgo/controller"
import _ "firstgo/servlet"

func main() {

	http.ListenAndServe(":8080", nil)

}
