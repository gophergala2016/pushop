package main

import (
	"os"
	"path"

	"github.com/codegangsta/cli"
)

const (
	appName    = "Pushop"
	appUsage   = "Command line image gallery generator"
	appVersion = "0.0.1"
)

func getWorkingPath() string {
	dirname, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return dirname
}

func getProjectPath(c *cli.Context) string {
	dirname := c.Args().First()
	if dirname == "" {
		dirname = getWorkingPath()
	}
	if _, err := os.Stat(dirname); os.IsNotExist(err) {
		panic(err)
	}
	return dirname
}

func main() {
	app := cli.NewApp()
	app.Name = appName
	app.Usage = appUsage
	app.Version = appVersion

	app.Commands = []cli.Command{
		{
			Name:  "init",
			Usage: "Initializes new project",
			Action: func(c *cli.Context) {
				projectPath := getProjectPath(c)
				err := NewConfig().Init(projectPath)
				if err != nil {
					panic(err)
				}
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
				config := NewConfig()
				projectPath := getProjectPath(c)
				if err := config.LoadFile(path.Join(projectPath, defaultConfigFileName)); err != nil {
					panic(err)
				}
				targetPath := path.Join(projectPath, defaultTargetPath)
				if err := NewBuild(projectPath, config).Generate(targetPath); err != nil {
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
		{
			Name:  "clean",
			Usage: "Cleans the target directory",
			Action: func(c *cli.Context) {
				println("clean")
			},
		},
	}
	app.Run(os.Args)
}
