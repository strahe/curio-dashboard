package main

import (
	lcli "github.com/filecoin-project/lotus/cli"
	logging "github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"
)

var log = logging.Logger("cmd")

var harmonydbFlag = &cli.StringFlag{
	Name:     "harmonydb-url",
	EnvVars:  []string{"CURIO_HARMONYDB_URL"},
	Usage:    "URL to connect to harmonydb, e.g. postgres://username:password@localhost:5433/database_name?search_path=curio",
	Required: true,
}

var appdbFlag = &cli.StringFlag{
	Name:    "appdb-url",
	EnvVars: []string{"CURIO_APPDB_URL"},
	Value:   "sqlite3://app.sqlite",
	Usage:   "URL to connect to appdb, e.g. postgres://username:password@localhost:5433/database_name?search_path=dashboard",
}

func main() {
	app := &cli.App{
		Name:                 "curio-dashboard",
		Usage:                "A dashboard for the Curio",
		Version:              "0.0.1", // todo:
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "Enable debug logging",
				Value: false,
			},
		},
		Before: func(c *cli.Context) error {
			if c.Bool("debug") {
				return logging.SetLogLevel("*", "DEBUG")
			} else {
				return logging.SetLogLevel("*", "INFO")
			}
		},
		After: func(c *cli.Context) error {
			return nil
		},
		Commands: []*cli.Command{
			runCmd,
			fillCmd,
		},
	}
	app.Setup()
	lcli.RunApp(app)
}
