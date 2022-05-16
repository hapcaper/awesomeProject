package service

type MyService interface {
	FindSomething() []string

	AddSomething(string, int)
}

type MyDao interface {
	Insert()
}

type SelfFuncInterface interface {
	SelfFunc(str string)
}
