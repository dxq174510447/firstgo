package astree

import "fmt"

type Aaer interface {
	DoSay(param string) (int, error)
}

type AaImpl struct {
}

func (a *AaImpl) DoSay(param string) (int, error) {
	fmt.Println("AaImpl DoSay")
	return 0, nil
}
