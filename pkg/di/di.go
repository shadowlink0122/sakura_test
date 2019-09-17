package di

import (
	"log"
	"sakura_test/pkg/application"
	"sakura_test/pkg/infra"
)

var (
	User  application.IUserApp
	Hello application.IHelloApp
)

func Init() {
	initUser()
	initHello()
}

func initUser() {
	user := infra.NewUserRepo()
	User = application.NewUserApp(user)
}

func initHello() {
	log.Println("Init Hello")
	Hello = application.NewHelloApp()
}
