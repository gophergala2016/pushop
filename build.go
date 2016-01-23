package main

import (
	"io"
	"os"
	"text/template"
)

const (
	defaultTargetDirectory = "_site"
)

type Build struct {
	target string
}

func newBuild(target string) *Build {
	return &Build{target}
}

func (b *Build) clean() error {
	return os.RemoveAll(b.target)
}

func (b *Build) createDirectory() error {
	return os.MkdirAll(b.target, 0600)
}

func (b *Build) frontPage(config Config, fileList []File, w io.Writer) error {
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
	return t.Execute(w, nil)
}
