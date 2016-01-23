package main

import (
	"io"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

const (
	defaultConfigFileName = "_config.yaml"

	defaultTitle       = "My Gallery"
	defaultDescription = "My Gallery Description"

	skipPrefix       = "_"
	hiddenFilePrefix = "."
)

type Config struct {
	Title       string           `yaml:"title"`
	Description string           `yaml:"description"`
	Author      string           `yaml:"author,omitempty"`
	Analytics   string           `yaml:"analytics"`
	Content     map[string]*File `yaml:"content"`
}

type File struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Tags        []string `yaml:"tags,omitempty"`
	Exif        bool     `yaml:"exif,omitempty"`
	Permalink   string   `yaml:"permalink,omitempty"`
}

func newConfig() *Config {
	return &Config{
		Title:       defaultTitle,
		Description: defaultDescription,
		Content:     make(map[string]*File),
	}
}

func (c *Config) load(r io.Reader) error {
	var data []byte
	var err error
	if data, err = ioutil.ReadAll(r); err != nil {
		return err
	}
	return yaml.Unmarshal(data, c)
}

func (c *Config) save(w io.Writer) error {
	d, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	_, err = w.Write(d)
	return err
}

func (c *Config) generate(dirname string) error {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		return err
	}
	for _, fileInfo := range files {
		if strings.HasPrefix(fileInfo.Name(), skipPrefix) || strings.HasPrefix(fileInfo.Name(), hiddenFilePrefix) {
			continue
		}
		if _, ok := c.Content[fileInfo.Name()]; ok {
			// Skip existing setup
			continue
		}
		c.Content[fileInfo.Name()] = &File{
			Name:        fileInfo.Name(),
			Description: "",
			Exif:        false,
		}
	}
	return nil
}
