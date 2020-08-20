package main

import (

	// dTwitterHttp "github.com/jauntyjack/compfest/delivery/twitter/http"
	dTwitterFlag "github.com/jauntyjack/compfest/delivery/twitter/flag"
	// rTwitterCli "github.com/jauntyjack/compfest/repository/twitter/cli"
	rTwitterFile "github.com/jauntyjack/compfest/repository/twitter/file"
	uTwitterModule "github.com/jauntyjack/compfest/usecase/twitter/module"
)

func main() {
	// twitterCli := rTwitterCli.New()
	twitterFile := rTwitterFile.New()
	twitterUsecase := uTwitterModule.New(twitterFile)
	dTwitterFlag.New(twitterUsecase)

	// log.Print("Listen to port 8090")
	// http.ListenAndServe(":8090", nil)
}
