package main

import "os"

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

func (b *Build) frontPage(config Config, fileList []File) error {

	return nil
}
