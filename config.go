package main

import (
	"io"
)

type Config struct {
	Title       string  `yaml:"title"`
	Description string  `yaml:"description"`
	Author      *string `yaml:"author"`
	Analytics   *string `yaml:"analytics"`
	Content     *[]File `yaml:"content"`
}

type File struct {
	Name        *string   `yaml:"name"`
	Description *string   `yaml:"description"`
	Tags        *[]string `yaml:"tags"`
	Exif        *bool     `yaml:"exif"`
	Permalink   *string   `yaml:"permaling"`
}

func newConfig() *Config {
	return &Config{}
}

func (c *Config) load(reader io.Reader) error {
	return nil
}

func (c *Config) save(writer io.WriteCloser) error {
	return nil
}
