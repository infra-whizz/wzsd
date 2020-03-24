package main

import (
	"os"

	"github.com/infra-whizz/wzsd"
	"github.com/isbm/go-nanoconf"
	"github.com/urfave/cli/v2"
)

func run(ctx *cli.Context) error {
	stateDaemon := wzsd.NewWzStateDaemon()
	stateDaemon.Run().AppLoop()

	conf := nanoconf.NewConfig(ctx.String("config"))
	stateDaemon.GetTransport().AddNatsServerURL(
		conf.Find("transport").String("host", ""),
		conf.Find("transport").DefaultInt("port", "", 4222),
	)
	stateDaemon.Run().AppLoop()

	cli.ShowAppHelpAndExit(ctx, 1)
	return nil
}

func main() {
	appname := "wz-state"
	confpath := nanoconf.NewNanoconfFinder(appname).DefaultSetup(nil)

	app := &cli.App{
		Version: "0.1 Alpha",
		Name:    appname,
		Usage:   "Whizz State Worker Daemon",
		Action:  run,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "config",
				Aliases:  []string{"c"},
				Usage:    "Path to configuration file",
				Required: false,
				Value:    confpath.SetDefaultConfig(confpath.FindFirst()).FindDefault(),
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
