package service

import "fmt"

type MyServiceImpl struct {
}

func (myServiceImpl MyServiceImpl) FindSomething() []string {
	result := []string{"sds", "123213"}
	fmt.Println("findSomething invoked,,,,,")
	return result
}
func (myServiceImpl MyServiceImpl) AddSomething(var1 string, var2 int) {
	fmt.Println("addSomething invoked....", var1, var2)
}

func (myServiceImpl MyServiceImpl) Insert() {
	fmt.Println("insert sth")
}

func (myServiceImpl MyServiceImpl) SelfFunc(str string) {
	fmt.Println("SelfFunc exc", str)
}
