package application

import(
	"github.com/shadowlink0122/sakura_test/pkg/infra"
)

type IUserApp interface{
	GetList()([]infra.User, error)
}

type userApp struct{
	infra.IUserRepo
}

func NewUserApp(repo infra.IUserRepo) IUserApp{
	return &userApp{repo}
}

func (uapp *userApp) GetList() ([]infra.User, error){
	user, err := uapp.UserList()
	if err != nil{
		return nil, err
	}

	return user, nil
}
