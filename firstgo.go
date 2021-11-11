package main

import (
	_ "firstgo/src/main/golang"
	_ "firstgo/src/main/resources"
	_ "github.com/dxq174510447/goframe/lib/frame"
	"github.com/dxq174510447/goframe/lib/frame/application"
)


func main() {
	// http.ListenAndServe(":8080", nil)
	args := []string{"--appli=123"}
	application.NewApplication(nil).Run(args)
}
