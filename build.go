package main

import (
	"io"
	"os"
	"path"
	"text/template"
)

const (
	defaultTargetDirectory = "_site"
)

type Build struct {
	config *Config
	target string
}

func newBuild(target string, config *Config) *Build {
	return &Build{config, target}
}

func (b *Build) generate() error {
	var fileList []File

	b.clean()
	b.createDirectory()

	indexFile, err := os.Create(path.Join(b.target, "index.html"))
	if err != nil {
		return err
	}
	defer indexFile.Close()
	err = b.frontPage(fileList, indexFile)
	if err != nil {
		return err
	}
	return nil
}

func (b *Build) clean() error {
	return os.RemoveAll(b.target)
}

func (b *Build) createDirectory() error {
	return os.MkdirAll(b.target, 0766)
}

func (b *Build) frontPage(fileList []File, w io.Writer) error {
	layoutData, err := Asset("assets/templates/layout.html")
	if err != nil {
		return err
	}
	pageData, err := Asset("assets/templates/index.html")
	if err != nil {
		return err
	}
	t, err := template.New("layout").Parse(string(layoutData))
	t.Parse(string(pageData))
	if err != nil {
		return err
	}
	templateData := map[string]interface{}{
		"Title":       b.config.Title,
		"Description": b.config.Description,
	}
	return t.Execute(w, templateData)
}
