package infra

type IUserRepo interface{
	UserList ([]User, error)
}

type userRepo struct{
}

func NewUserRepo() IUserRepo {
	return &userRepo{}
}

type User struct{
	Name string `json:"name"`
}

func (urepo *userRepo) UserList ([]User, error) {
	lcap := 10
	userList := make([]UserList, lcap, lcap)

	for i := 0;i < lcap;i++{
		userName := fmt.Sprintf("Name%d", i)
		user := User{Name: userName}

		userList = append(userList, user)
	}

	return userList, nil
}
