package main

import (
	"log"
	"net/http"

	dTwitterHttp "github.com/jauntyjack/compfest/delivery/twitter/http"
	rTwitterCli "github.com/jauntyjack/compfest/repository/twitter/cli"
	uTwitterModule "github.com/jauntyjack/compfest/usecase/twitter/module"
)

func main() {
	twitterCli := rTwitterCli.New()
	twitterUsecase := uTwitterModule.New(twitterCli)
	dTwitterHttp.New(twitterUsecase)

	log.Print("Listen to port 8090")

	http.ListenAndServe(":8090", nil)
}
