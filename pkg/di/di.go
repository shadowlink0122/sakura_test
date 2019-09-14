package di

import(
	"sakura_test/pkg/infra"
	"sakura_test/pkg/application"
)

var User application.IUserApp

func Init(){
	initUser()
}

func initUser(){
	user := infra.NewUserRepo()
	User = application.NewUserApp(user)
}
