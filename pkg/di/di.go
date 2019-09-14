package di

import(
	"github.com/shadowlink0122/sakura_test/pkg/infra"
	"github.com/shadowlink0122/sakura_test/pkg/application"
)

var User application.IUserApp

func Init(){
	initUser()
}

func initUser(){
	user := infra.NewUserRepo()
	User = application.NewUserApp(user)
}
