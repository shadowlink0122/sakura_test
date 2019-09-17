package application

type IHelloApp interface {
	GetHello() (string, error)
}

type HelloApp struct{}

func NewHelloApp() IHelloApp {
	return &HelloApp{}
}

func (ha *HelloApp) GetHello() (string, error) {
	return "Hello World!", nil
}
