package server

import (
	"log"
	"net/http"

	"sakura_test/pkg/di"
	"sakura_test/pkg/server/handler"
)

func Serve() {
	di.Init()

	// routing
	http.HandleFunc("/", get(handler.GetMessage()))
	http.HandleFunc("/user", get(handler.GetUser()))

	// run
	/* ===== サーバの起動 ===== */
	log.Println("Server running...")
	// LocalHost
	err := http.ListenAndServe(":8040", nil)
	if err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}

	// CGI
	// cgi.Serve(nil)
}
