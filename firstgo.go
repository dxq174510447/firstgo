package main

import (
	_ "firstgo/src/main/golang"
	_ "firstgo/src/main/resources"
	_ "github.com/dxq174510447/goframe/lib/frame"
	"github.com/dxq174510447/goframe/lib/frame/application"
)

type FirstGo struct {
}

func (f *FirstGo) Run(args []string) {
	application.NewApplication(f).Run(args)
}

func main() {
	// http.ListenAndServe(":8080", nil)
	args := []string{"--appli=123"}

	var instance *FirstGo = &FirstGo{}
	instance.Run(args)
}
