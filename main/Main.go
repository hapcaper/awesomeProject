package main

import (
	"awesomeProject/service"
)

func main() {

	var s = service.MyServiceImpl{}
	var mi service.SelfFuncInterface
	var ms service.MyService
	var dao service.MyDao
	mi = s
	dao = s
	ms = s
	ms.AddSomething("sdsd", 1)
	ms.FindSomething()
	dao.Insert()
	mi.SelfFunc("abcdefg")
}
