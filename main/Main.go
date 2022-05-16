package main

import (
	"awesomeProject/service"
	"fmt"
)

func main() {

	fmt.Println("sdsdsdsds")

	var ms service.MyService
	ms = service.MyServiceImpl{}
	ms.AddSomething("sdsd", 1)
	ms.FindSomething()
}
