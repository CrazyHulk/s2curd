package main

import (
	"os"

	"github.com/janfly79/s2curd/sqlorm"

	log "github.com/liudanking/goutil/logutil"

	"github.com/urfave/cli"
)

var (
	// the variables will be set at compile time from golang build ldflags
	service   = "service"
	version   = "version"
	buildDate = "build_date"
	commitID  = "commit_id"
)

func main() {
	app := cli.NewApp()
	app.Version = version
	app.Usage = "Codoon Backend Developer work flow Swiss Army Knife"

	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		sqlorm.SqlCommand(),
		sqlorm.CurdCommand(),
		sqlorm.CacheCommand(),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Error(err.Error())
		return
	}
}
