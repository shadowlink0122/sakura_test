package server

import(
	"net/http"
	"net/http/cgi"

	"sakura_test/pkg/di"
	"sakura_test/pkg/server/handler"
)

func Serve(){
	di.Init()

	// routing
	http.HandleFunc("/user", get(handler.UserGet()))

	// run
	cgi.Serve(nil)
}
