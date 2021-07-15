package astree

import "fmt"

type Aaer interface {
	DoSay() (int,error)
}

type AaImpl struct {

}
func (a *AaImpl) DoSay() (int,error){
	fmt.Println("AaImpl DoSay")
	return 0,nil
}