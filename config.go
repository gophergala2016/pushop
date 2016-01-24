package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"path"
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
	Permalink   string   `yaml:"permalink,omitempty"`
}

func NewConfig() *Config {
	return &Config{
		Title:       defaultTitle,
		Description: defaultDescription,
		Content:     make(map[string]*File),
	}
}

func (c *Config) LoadFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	return c.Load(file)
}

func (c *Config) Load(r io.Reader) error {
	var data []byte
	var err error
	if data, err = ioutil.ReadAll(r); err != nil {
		return err
	}
	return yaml.Unmarshal(data, c)
}

func (c *Config) Save(w io.Writer) error {
	d, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	_, err = w.Write(d)
	return err
}

func (c *Config) Init(projectPath string) error {
	configFile := path.Join(projectPath, defaultConfigFileName)
	if _, err := os.Stat(configFile); err == nil {
		c.LoadFile(configFile)
	}
	err := c.generate(projectPath)
	if err != nil {
		return err
	}
	file, err := os.Create(configFile)
	if err != nil {
		return err
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	err = c.Save(w)
	if err != nil {
		return err
	}
	w.Flush()
	return file.Close()
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
		if _, ok := c.Content[fileInfo.Name()]; ok {
			// Skip existing setup
			continue
		}
		c.Content[fileInfo.Name()] = &File{
			Name:        fileInfo.Name(),
			Description: "",
		}
	}
	return nil
}
