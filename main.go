package main

import (
	"log"
	"net/http"
	"os"
	"path"

	"github.com/codegangsta/cli"
)

const (
	appName    = "Pushop"
	appUsage   = "Command line image gallery generator"
	appVersion = "0.0.1"

	serverPort = ":8088"
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

func buildProject(projectPath, targetPath string, c *cli.Context) error {
	config := NewConfig()
	configFile := path.Join(projectPath, defaultConfigFileName)
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return err
	}
	if err := config.LoadFile(configFile); err != nil {
		return err
	}
	if err := NewBuild(projectPath, config).Generate(targetPath); err != nil {
		return err
	}
	return nil
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
			Name:  "build",
			Usage: "Builds the gallery",
			Action: func(c *cli.Context) {
				projectPath := getProjectPath(c)
				targetPath := path.Join(projectPath, defaultTargetSegment)
				if err := buildProject(projectPath, targetPath, c); err != nil {
					log.Fatal(err)
					os.Exit(1)
				}
			},
		},
		{
			Name:  "serve",
			Usage: "Builds and serves the gallery",
			Action: func(c *cli.Context) {
				projectPath := getProjectPath(c)
				targetPath := path.Join(projectPath, defaultTargetSegment)
				if err := buildProject(projectPath, targetPath, c); err != nil {
					log.Fatal(err)
					os.Exit(1)
				}

				log.Printf("Listening on port %s", serverPort)
				log.Fatal(http.ListenAndServe(serverPort, http.FileServer(http.Dir(targetPath))))
			},
		},
		{
			Name:  "clean",
			Usage: "Cleans the target directory",
			Action: func(c *cli.Context) {
				projectPath := getProjectPath(c)
				targetPath := path.Join(projectPath, defaultTargetSegment)
				cleanProject(targetPath)
			},
		},
	}
	app.Run(os.Args)
}
