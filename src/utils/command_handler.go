package utils

import (
	. "GOAT/src/core"
	"github.com/urfave/cli"
	"log"
	"os"
)

func info(app *cli.App) {
	app.Name = "GOAT CLI"
	app.Usage = "GOAT CLI for code generation of golang backend"
	app.Author = "Samdy Vin"
	app.Version = "1.0.0"
}

func commands(app *cli.App) {
	goat := Goat{}

	app.Commands = []cli.Command{
		{
			Name:      "new",
			Aliases:   []string{"n"},
			Usage:     "Code Generation",
			ArgsUsage: "All available commands for code generation.",
			//Category:  "Generation",
			Subcommands: []cli.Command{
				{
					Name:      "service",
					Aliases:   []string{"svc"},
					Usage:     "Generate a new Service.",
					ArgsUsage: "Generate a new fresh Service code.",
					//Category:  "Generation",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:     "name",
							Usage:    "Name of your Service",
							Required: true,
						},
					},
					Action: func(c *cli.Context) {
						goat.GenerateService(c.String("name"))
					},
				},
				{
					Name:      "entity",
					Aliases:   []string{"e"},
					Usage:     "Generate a new Entity.",
					ArgsUsage: "Generate a new Entity code for Service.",
					//Category:  "Generation",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:     "name",
							Usage:    "Name of your Entity",
							Required: true,
						},
					},
					Action: func(c *cli.Context) {
						goat.GenerateEntity(c.String("name"))
					},
				},
				{
					Name:      "protocol",
					Aliases:   []string{"ptc"},
					Usage:     "Generate a new Protocol.",
					ArgsUsage: "Generate a new Protocol code for Service.",
					//Category:  "Generation",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:     "name",
							Usage:    "Name of your Protocol",
							Required: true,
						},
					},
					Action: func(c *cli.Context) {
						goat.GenerateProtocol(c.String("name"))
					},
				},
				{
					Name:      "repository",
					Aliases:   []string{"repo"},
					Usage:     "Generate a new Repository.",
					ArgsUsage: "Generate a new Repository code for Service.",
					//Category:  "Generation",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:     "name",
							Usage:    "Name of your Repository",
							Required: true,
						},
					},
					Action: func(c *cli.Context) {
						goat.GenerateRepo(c.String("name"))
					},
				},
			},
		},
	}
}

func CommandHandler() {
	var app = cli.NewApp()

	info(app)
	commands(app)

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
