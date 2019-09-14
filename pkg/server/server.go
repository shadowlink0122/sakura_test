package server

import(
	"net/http"
	"sakura_test/pkg/di"
	"sakura_test/pkg/handler"
)

func Serve(){
	di.Init()

	// routing
	http.HandleFunc("/user", get(handler.UserGet()))

	cgi.Serve(nil)
}
