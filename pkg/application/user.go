package application

import(
	"github.com/shadowlink0122/sakura_test/pkg/infra"
)

type IUserApp interface{
	GetList()
}

type userApp{
	infra.IUserRepo
}

func NewUserApp(repo infra.IUserRepo) IUserApp{
	return &userApp{repo}
}

type (uapp *userApp) GetList() (,error){
	user, err := uapp.UserList()
	if err != nil{
		return nil, err
	}

	return user, nil
}
