package main

import (
	"demo/conf"
	"demo/migrate"
	"demo/router"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "Account Tool"

	app.Usage = "User's command line"

	app.Version = "1.0.0"

	app.Commands = []cli.Command{
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "run http server",
			Action: func(c *cli.Context) error {
				conf := conf.ReadConfFile()
				app := router.RouteLink()

				s := &http.Server{
					Addr:           ":" + conf.GetString("service.host"),
					Handler:        app,
					ReadTimeout:    10 * time.Second,
					WriteTimeout:   10 * time.Second,
					MaxHeaderBytes: 1 << 20,
				}
				s.ListenAndServe()
				return nil
			},
		},
		{
			Name:    "migrate",
			Aliases: []string{"m"},
			Usage:   "migrate table",
			Action: func(c *cli.Context) error {
				migrate.Migrate()
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
