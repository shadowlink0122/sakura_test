package server

import(
	"net/http"
	"net/http/cgi"
	
	"github.com/shadowlink0122/sakura_test/pkg/di"
	"github.com/shadowlink0122/sakura_test/pkg/handler"
)

func Serve(){
	di.Init()

	// routing
	http.HandleFunc("/user", get(handler.UserGet()))

	cgi.Serve(nil)
}
