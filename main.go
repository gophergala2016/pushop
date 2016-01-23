package main

import (
	"bufio"
	"os"
	"path"

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
				if _, err := os.Stat(dirname); os.IsNotExist(err) {
					panic(err)
				}
				if err = config.generate(dirname); err != nil {
					panic(err)
				}
				file, err := os.Create(path.Join(dirname, defaultConfigFileName))
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
				var err error
				dirname := c.Args().First()
				var currentDir string
				if currentDir, err = os.Getwd(); err != nil {
					panic(err)
				}
				if dirname == "" {
					dirname = path.Join(currentDir, defaultTargetDirectory)
				}
				file, err := os.Open(path.Join(currentDir, defaultConfigFileName))
				defer file.Close()
				err = config.load(file)
				if err != nil {
					panic(err)
				}
				err = newBuild(dirname, config).generate()
				if err != nil {
					panic(err)
				}
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
