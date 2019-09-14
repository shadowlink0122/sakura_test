package server

import(
	"log"
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
	/* ===== サーバの起動 ===== */
	log.Println("Server running...")
	// LocalHost
	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	log.Fatalf("Listen and serve failed. %+v", err)
	// }

	// CGI
	cgi.Serve(nil)
}
