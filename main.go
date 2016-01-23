package main

import (
	"bufio"
	"os"

	"github.com/codegangsta/cli"
)

const (
	appName    = "Pushop"
	appUsage   = "Command line image gallery generator"
	appVersion = "0.0.1"
)

func main() {
	app := cli.NewApp()
	app.Name = appName
	app.Usage = appUsage
	app.Version = appVersion

	config := newConfig()

	app.Commands = []cli.Command{
		{
			Name:  "init",
			Usage: "Initializes new project",
			Action: func(c *cli.Context) {
				var err error
				dirname := c.Args().First()
				if dirname == "" {
					if dirname, err = os.Getwd(); err != nil {
						panic(err)
					}
				}
				// TODO: Check if directory is valid
				if err = config.generate(dirname); err != nil {
					panic(err)
				}
				// TODO: Make the path OS agnostic
				file, err := os.Create(dirname + "/" + defaultConfigFileName)
				defer file.Close()
				if err != nil {
					panic(err)
				}
				w := bufio.NewWriter(file)
				err = config.save(w)
				if err != nil {
					panic(err)
				}
				w.Flush()
				file.Close()
			},
		},
		{
			Name:  "updates",
			Usage: "Updates existing project",
			Action: func(c *cli.Context) {
				println("update")
			},
		},
		{
			Name:  "build",
			Usage: "Builds the gallery",
			Action: func(c *cli.Context) {
				println("build")
			},
		},
		{
			Name:  "serve",
			Usage: "Builds and serves the gallery",
			Action: func(c *cli.Context) {
				println("serve")
			},
		},
	}
	app.Run(os.Args)
}
