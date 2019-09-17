package handler

import (
	"log"
	"net/http"
	"sakura_test/pkg/di"
	"sakura_test/pkg/server/response"
)

type GetMessageResponse struct {
	Message string `json:"message"`
}

func GetMessage() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		msg, err := di.Hello.GetHello()
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		resp := messageToResponseBody(msg)

		response.Success(
			writer,
			resp,
		)
	}
}

func messageToResponseBody(msg string) GetMessageResponse {
	return GetMessageResponse{Message: msg}
}
