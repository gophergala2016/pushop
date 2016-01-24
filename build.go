package main

import (
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"
)

const (
	defaultTargetDirectory = "_site"
)

type Build struct {
	config *Config
	target string
}

func NewBuild(target string, config *Config) *Build {
	return &Build{config, target}
}

func (b *Build) Generate(source string) error {

	fileList, err := b.collectFiles(source)
	if err != nil {
		return err
	}

	// Clean the project
	b.Clean()
	b.createDirectory(b.target)

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

func (b *Build) Clean() error {
	return os.RemoveAll(b.target)
}

func (b *Build) createDirectory(path string) error {
	return os.MkdirAll(path, 0766)
}

func (b *Build) collectFiles(source string) (map[string]*File, error) {
	content := make(map[string]*File)
	files, err := ioutil.ReadDir(source)
	if err != nil {
		return nil, err
	}
	for _, fileInfo := range files {
		if strings.HasPrefix(fileInfo.Name(), skipPrefix) || strings.HasPrefix(fileInfo.Name(), hiddenFilePrefix) {
			continue
		}
		if fileInfo.IsDir() {
			// Skip directories
			continue
		}
		file, err := os.Open(fileInfo.Name())
		if err != nil {
			continue
		}
		defer file.Close()
		if !isSupported(getContentType(file)) {
			continue
		}
		if _, ok := b.config.Content[fileInfo.Name()]; ok {
			content[fileInfo.Name()] = b.config.Content[fileInfo.Name()]
			continue
		}
		content[fileInfo.Name()] = &File{
			Name:        fileInfo.Name(),
			Description: "",
			Exif:        true,
		}
	}
	return content, nil
}

func (b *Build) frontPage(fileList map[string]*File, w io.Writer) error {
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
		"Analytics":   b.config.Analytics,
		"Images":      fileList,
	}
	return t.Execute(w, templateData)
}
