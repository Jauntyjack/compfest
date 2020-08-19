package main

import (

	// dTwitterHttp "github.com/tokopedia/dev/compfest/delivery/twitter/http"
	dTwitterFlag "github.com/tokopedia/dev/compfest/delivery/twitter/flag"
	// rTwitterCli "github.com/tokopedia/dev/compfest/repository/twitter/cli"
	rTwitterFile "github.com/tokopedia/dev/compfest/repository/twitter/file"
	uTwitterModule "github.com/tokopedia/dev/compfest/usecase/twitter/module"
)

func main() {
	// twitterCli := rTwitterCli.New()
	twitterFile := rTwitterFile.New()
	twitterUsecase := uTwitterModule.New(twitterFile)
	// dTwitterHttp.New(twitterUsecase)
	dTwitterFlag.New(twitterUsecase)

	// http.ListenAndServe(":8090", nil)
}
