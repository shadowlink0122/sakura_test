package handler

import(
	"log"

	"sakura_test/pkg/di"
	"sakura_test/pkg/server/response"
)

type UserGetResponse struct{
	Users []string `json:"user"`
}

func UserGet() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request){
		users, err := di.User.GetList()
		if err != nil{
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		response.Success(
			writer,
			UserGetResponse{
				Users: users
			}
		)
	}
}
